package cli

import (
	"github.com/jxsl13/kcauth"
	configo "github.com/jxsl13/simple-configo"
)

// LoginToken logs you in via the cli workflow, prompts for username and password and fetches an oflfine token from
// the rarget keycloak and caches the access token as well as the refresh token locally.
func LoginToken(tokenOut *kcauth.Token, issuerURL *string) configo.ParserFunc {
	return func(value string) error {
		token, err := jwtLogin(*issuerURL, "", "", kcauth.DefaultClientID, kcauth.DefaultClientSecret)
		if err != nil {
			return err
		}
		*tokenOut = *token
		return nil
	}
}

//Login tries to login via the provided user credentials.
// if oyu want to modify the clientID as well as the clientSecret, you need to do so in the base kcauth package/module.
// This function returns a simple-configo ParserFunc that uses the provided credentials in order to fetch a keycloak
// access_token as well as a refresh_token.
func Login(outToken *kcauth.Token, issuerURL, username, password *string) configo.ParserFunc {
	return func(value string) error {
		token, err := jwtLogin(*issuerURL, *username, *password, kcauth.DefaultClientID, kcauth.DefaultClientSecret)
		if err != nil {
			return err
		}
		*outToken = *token
		return nil
	}
}
