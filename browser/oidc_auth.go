package browser

import (
	"context"
	"log"
	"os"

	"github.com/jxsl13/kcauth"
	"github.com/jxsl13/oidc"
	"github.com/jxsl13/oidc/login"
	disk "github.com/jxsl13/oidc/login/diskcache"
)

func newTokenFromOidc(token *oidc.Token) *kcauth.Token {
	resultToken := &kcauth.Token{
		AccessToken:           token.AccessToken,
		RefreshToken:          token.RefreshToken,
		AccessTokenExpiresAt:  token.AccessTokenExpiry,
		RefreshTokenExpiresAt: kcauth.EpochZero,
	}

	return resultToken
}

func refreshToken(token *oidc.Token, oidcConfig login.OIDCConfig, issuerURL, redirectURL string) (*oidc.Token, error) {
	// token is expired
	client, err := oidc.NewClient(context.Background(), issuerURL)
	if err != nil {
		return nil, err
	}
	refresher := oidc.NewTokenRefresher(client, oidc.Config{
		ClientID:     oidcConfig.ClientID,
		ClientSecret: oidcConfig.ClientSecret,
		RedirectURL:  redirectURL,
		Scopes:       oidcConfig.Scopes,
	}, token.RefreshToken)
	return refresher.OIDCToken(context.Background())
}

func oidcLogin(clientID, clientSecret, issuerURL, redirectURL, cacheDirectory string) (*kcauth.Token, error) {

	// config
	oidcConfig := login.OIDCConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Provider:     issuerURL,
		// Make sure you ask for offline_access if you want to use refresh tokens!
		Scopes: []string{"openid", "offline_access", "email", "profile"},
	}

	// protect against replay attacks
	sourceConfig := login.Config{
		NonceCheck: true,
	}
	// see also other caches e.g k8s.NewCache.
	cache := disk.NewCache(cacheDirectory, oidcConfig)

	// try getting token from file
	// if successfully fetched token, simply return it
	token, err := cache.Token()
	if err == nil && token != nil {
		if !token.IsAccessTokenExpired() {
			return newTokenFromOidc(token), nil
		}
		token, err = refreshToken(token, oidcConfig, issuerURL, redirectURL)
		if err != nil {
			return nil, err
		}
		err = cache.SaveToken(token)
		if err != nil {
			return nil, err
		}
		return newTokenFromOidc(token), err
	}

	cb, closeSrv, err := login.NewServer(redirectURL)
	if err != nil {
		return nil, err
	}
	defer closeSrv()

	// in case we want to revoke tokens, we might the _ parameter
	source, _, err := login.NewOIDCTokenSource(
		context.Background(),
		log.New(os.Stdout, "", 0),
		sourceConfig,
		cache,
		cb,
	)
	if err != nil {
		return nil, err
	}

	token, err = source.OIDCToken(context.Background())
	if err != nil {
		return nil, err
	}

	return newTokenFromOidc(token), nil
}
