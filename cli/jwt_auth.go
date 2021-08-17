package cli

import (
	"context"

	gocloak "github.com/Nerzal/gocloak/v8"
	"github.com/jxsl13/kcauth"
	"github.com/jxsl13/kcauth/internal"
)

func getOfflineToken(realmURL, clientID, clientSecret, username, password string) (*kcauth.Token, error) {
	keycloakURL, realm, err := internal.ExtractRealm(realmURL)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	client := gocloak.NewClient(keycloakURL)

	grantType := "password"
	if username == "" && password == "" {
		grantType = "client_credentials"
	}

	token, err := client.GetToken(
		ctx,
		realm,
		gocloak.TokenOptions{
			ClientID:      gocloak.StringP(clientID),
			ClientSecret:  gocloak.StringP(clientSecret),
			GrantType:     gocloak.StringP(grantType),
			Scopes:        &[]string{"openid", "offline_access"},
			ResponseTypes: &[]string{"token"},
			Username:      gocloak.StringP(username),
			Password:      gocloak.StringP(password),
		},
	)
	if err != nil {
		return nil, err
	}

	return internal.NewTokenFromGoCloak(token), nil
}

func jwtLogin(issuerURL, username, password, clientID, clientSecret string) (*kcauth.Token, error) {
	return getOfflineToken(issuerURL, clientID, clientSecret, username, password)
}
