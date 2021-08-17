package browser

import (
	"github.com/jxsl13/kcauth"
	configo "github.com/jxsl13/simple-configo"
)

// LoginToken uses the provided issuerURL to fetch an access token as well as a refresh token.
// usually the expected URL is something along the line sof https://keycloak.com/auth/realms/MYREALM
// accessToken returns the user's access token that can be used to call api endpoints.
// That token is usually passed in the Authorization http header like this:
// Authorization: Bearer <access_token>
func Login(outToken *kcauth.Token, oidcURL *string) configo.ActionFunc {

	return func() error {
		token, err := oidcLogin(kcauth.DefaultClientID, kcauth.DefaultClientSecret, *oidcURL, DefaultRedirectURL)
		if err != nil {
			return err
		}
		*outToken = *token
		return nil
	}
}
