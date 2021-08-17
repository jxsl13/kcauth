package internal

import (
	"context"

	"github.com/Nerzal/gocloak/v8"
)

func RevokeToken(token string, hint ...tokenTypeHint) error {
	url, realm, err := UrlRealmFromToken(token)
	if err != nil {
		return err
	}

	form := make(map[string]string, 2)
	form["token"] = token

	if len(hint) > 0 {
		form["token_type_hint"] = string(hint[0])
	}

	client := gocloak.NewClient(url)
	ctx := context.Background()
	req := client.RestyClient().
		SetHostURL(url).
		R().
		SetContext(ctx).
		SetFormData(form).
		SetPathParams(map[string]string{
			"realm": realm,
		})
	resp, err := req.Post("/auth/realms/{realm}/protocol/openid-connect/revoke")
	if err != nil {
		return err
	}

	_, err = checkResponse(resp)
	if err != nil {
		return err
	}
	return nil

}
