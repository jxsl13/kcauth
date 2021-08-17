package token

import (
	"github.com/jxsl13/kcauth"
	"github.com/jxsl13/kcauth/internal"
	configo "github.com/jxsl13/simple-configo"
)

func RevokeAccessToken(token *kcauth.Token) configo.ActionFunc {
	return func() error {
		return internal.RevokeToken(token.AccessToken, internal.TypeHintAccessToken)

	}
}

func RevokeRefreshToken(token *kcauth.Token) configo.ActionFunc {
	return func() error {
		return internal.RevokeToken(token.RefreshToken, internal.TypeHintRefreshToken)
	}
}
