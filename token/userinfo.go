package token

import (
	"github.com/jxsl13/kcauth"
	"github.com/jxsl13/kcauth/internal"
	configo "github.com/jxsl13/simple-configo"
)

// UserInfo fetches the User Info for the specified token.
// token must contain a valid accessToken
func UserInfo(outInfo *kcauth.UserInfo, token *kcauth.Token) configo.ActionFunc {
	return func() error {
		info, err := internal.UserInfo(token.AccessToken)
		if err != nil {
			return err
		}
		*outInfo = *info
		return err
	}
}
