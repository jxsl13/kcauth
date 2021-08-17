package internal

import (
	"context"

	"github.com/Nerzal/gocloak/v8"
	"github.com/jxsl13/kcauth"
)

func UserInfo(accessToken string) (*kcauth.UserInfo, error) {
	url, realm, err := UrlRealmFromToken(accessToken)
	if err != nil {
		return nil, err
	}
	client := gocloak.NewClient(url)
	ctx := context.Background()
	info, err := client.GetUserInfo(ctx, accessToken, realm)
	if err != nil {
		return nil, err
	}
	return (*kcauth.UserInfo)(info), nil

}
