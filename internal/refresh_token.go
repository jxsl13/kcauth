package internal

import (
	"context"

	"github.com/Nerzal/gocloak/v8"
	"github.com/jxsl13/kcauth"
)

func RefreshToken(clientID, clientSecret, refreshToken string) (*kcauth.Token, error) {
	url, realm, err := UrlRealmFromToken(refreshToken)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	client := gocloak.NewClient(url)
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
	return NewTokenFromGoCloak(token), nil
}
