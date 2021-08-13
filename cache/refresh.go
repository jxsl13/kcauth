package cache

import (
	"context"

	"github.com/Nerzal/gocloak/v8"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/jxsl13/kcauth"
	"github.com/jxsl13/kcauth/internal"
)

func refreshToken(clientID, clientSecret, refreshToken string) (*kcauth.Token, error) {
	rToken, err := jwt.Parse([]byte(refreshToken))
	if err != nil {
		return nil, err
	}

	keycloakURL, realm, err := internal.ExtractRealm(rToken.Issuer())
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
	return internal.NewTokenFromGoCloak(token), nil
}
