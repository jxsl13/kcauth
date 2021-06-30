package cli

import (
	"github.com/jxsl13/kcauth"
	configo "github.com/jxsl13/simple-configo"
)

// LoginToken logs you in via the cli workflow, prompts for username and password and fetches an oflfine token from
// the rarget keycloak and caches the access token as well as the refresh token locally.
func LoginToken(tokenOut *kcauth.Token, issuerURL *string) configo.ParserFunc {
	return func(value string) error {
		token, err := jwtLogin(*issuerURL, kcauth.DefaultClientID, kcauth.DefaultClientSecret, kcauth.DefaultCacheDirectory)
		if err != nil {
			return err
		}
		*tokenOut = *token
		return nil
	}
}

// Login logs you in via the cli workflow, prompts for username and password and fetches an oflfine token from
// the rarget keycloak and caches the access token as well as the refresh token locally.
func Login(accessToken *string, issuerURL *string) configo.ParserFunc {
	return func(value string) error {
		token, err := jwtLogin(*issuerURL, kcauth.DefaultClientID, kcauth.DefaultClientSecret, kcauth.DefaultCacheDirectory)
		if err != nil {
			return err
		}
		*accessToken = token.AccessToken
		return nil
	}
}
