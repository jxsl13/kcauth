package internal

import (
	"context"

	"github.com/Nerzal/gocloak/v8"
)

func RevokeToken(token string, hint ...tokenTypeHint) (*string, error) {
	url, realm, err := UrlRealmFromToken(token)
	if err != nil {
		return nil, err
	}

	form := make(map[string]string, 2)
	form["token"] = token

	if len(hint) > 0 {
		form["token_type_hint"] = string(hint[0])
	}

	client := gocloak.NewClient(url)
	ctx := context.Background()
	req := client.RestyClient().R().
		SetContext(ctx).
		SetFormData(form).
		SetPathParams(map[string]string{
			"url":   url,
			"realm": realm,
		})
	resp, err := req.Post("{url}/auth/realms/{realm}/protocol/openid-connect/revoke")
	if err != nil {
		return nil, err
	}
	body := string(resp.Body())
	return &body, nil

}
