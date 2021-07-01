package cli

import (
	"context"
	"fmt"
	"path"
	"regexp"

	gocloak "github.com/Nerzal/gocloak/v8"
	"github.com/jxsl13/kcauth"
	"github.com/jxsl13/simple-configo/parsers"
)

var (
	extractRegex = regexp.MustCompile(`^(https?://[a-z0-9-\.:]{5,})/auth/realms/([^/]+)/?.*$`)
)

func extractRealm(url string) (domain, realm string, err error) {
	if matches := extractRegex.FindStringSubmatch(url); matches != nil {
		return matches[1], matches[2], nil
	}
	return "", "", fmt.Errorf("invalid keycloak realm url: %s", url)
}

func getOfflineToken(realmURL, clientID, clientSecret, username, password string) (*kcauth.Token, error) {
	keycloakURL, realm, err := extractRealm(realmURL)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	client := gocloak.NewClient(keycloakURL)

	grantType := "password"
	if username == "" && password == "" {
		grantType = "client_credentials"
	}

	scopes := &[]string{"openid"}

	// offline access only for non automated workflows.
	if grantType == "password" {
		*scopes = append(*scopes, "offline_access")
	}

	token, err := client.GetToken(
		ctx,
		realm,
		gocloak.TokenOptions{
			ClientID:      gocloak.StringP(clientID),
			ClientSecret:  gocloak.StringP(clientSecret),
			GrantType:     gocloak.StringP(grantType),
			Scopes:        scopes,
			ResponseTypes: &[]string{"token"},
			Username:      gocloak.StringP(username),
			Password:      gocloak.StringP(password),
		},
	)
	if err != nil {
		return nil, err
	}

	return newTokenFromGoCloak(token), nil
}

func refreshToken(issuerUrl, clientID, clientSecret, refreshToken string) (*kcauth.Token, error) {
	keycloakURL, realm, err := extractRealm(issuerUrl)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	client := gocloak.NewClient(keycloakURL)
	token, err := client.RefreshToken(
		ctx,
		refreshToken,
		clientID,
		clientSecret,
		realm,
	)
	if err != nil {
		return nil, err
	}
	return newTokenFromGoCloak(token), nil
}

func jwtLogin(issuerURL, clientID, clientSecret, cacheDirectory string) (*kcauth.Token, error) {
	tokenFile := fmt.Sprintf("token_jwt_%s", clientID)
	cachedFile := path.Join(cacheDirectory, tokenFile)
	cachedToken, err := loadToken(cachedFile)
	if err == nil {
		// loaded token successfully
		if !cachedToken.IsAccessTokenExpired() {
			return cachedToken, nil
		}
		// access token expired
		if !cachedToken.IsRefreshTokenExpired() {
			token, err := refreshToken(issuerURL, clientID, clientSecret, cachedToken.RefreshToken)
			if err == nil {
				// save refreshed token back to file
				cachedToken, err = saveToken(cachedFile, token)
				if err != nil {
					return nil, err
				}
				return cachedToken, nil
			}
		}
		// refresh token is also expired or failed to refresh
	}

	// we need a new token
	username := ""
	password := ""

	// this is only empty when we use the public grant type
	if clientSecret == "" {
		err = parsers.PromptText(&username, "Enter your username>")("")
		if err != nil {
			return nil, err
		}
		err = parsers.PromptPassword(&password, "Enter your password>")("")
		if err != nil {
			return nil, err
		}
	}
	// passing an empty username and password triggers the client_credentials grant type
	token, err := getOfflineToken(issuerURL, clientID, clientSecret, username, password)
	if err != nil {
		return nil, err
	}
	cachedToken, err = saveToken(cachedFile, token)
	if err != nil {
		return nil, err
	}
	return cachedToken, nil
}
