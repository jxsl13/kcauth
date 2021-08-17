package token

import (
	"github.com/jxsl13/kcauth"
	"github.com/jxsl13/kcauth/internal"
	configo "github.com/jxsl13/simple-configo"
)

func RefreshAccessToken(token *kcauth.Token) configo.ActionFunc {
	return func() error {
		newToken, err := internal.RefreshToken(kcauth.DefaultClientID, kcauth.DefaultClientSecret, token.AccessToken)
		if err != nil {
			return err
		}
		*token = *newToken
		return nil
	}
}
