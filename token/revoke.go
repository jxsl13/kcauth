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
		fmt.Printf("revoked token: %s\n", *resp)
		return err

	}
}

func RevokeRefreshToken(token *kcauth.Token) configo.ActionFunc {
	return func() error {
		resp, err := internal.RevokeToken(token.RefreshToken, internal.TypeHintRefreshToken)
		fmt.Printf("revoked token: %s\n", *resp)
		return err
	}
}
