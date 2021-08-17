package token

import (
	"fmt"

	"github.com/jxsl13/kcauth"
	"github.com/jxsl13/kcauth/internal"
	configo "github.com/jxsl13/simple-configo"
)

func RevokeAccessToken(token *kcauth.Token) configo.ActionFunc {
	return func() error {
		resp, err := internal.RevokeToken(token.AccessToken, internal.TypeHintAccessToken)
		if err != nil {
			return err
		}
		fmt.Printf("revoked token: %s\n", *resp)
		return nil

	}
}

func RevokeRefreshToken(token *kcauth.Token) configo.ActionFunc {
	return func() error {
		resp, err := internal.RevokeToken(token.RefreshToken, internal.TypeHintRefreshToken)
		if err != nil {
			return err
		}
		fmt.Printf("revoked token: %s\n", *resp)
		return nil
	}
}
