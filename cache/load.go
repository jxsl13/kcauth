package cache

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"

	"github.com/jxsl13/kcauth"
	"github.com/jxsl13/kcauth/internal"
	configo "github.com/jxsl13/simple-configo"
)

// This function returns a Simple-Configo ParserFunc that may be executed in a simple-configo context.
// The returned function tries to read the json formated token file at the given location tokenFilePath
// in case the token is expired but the refresh token still active, the whole token is refreshed before
// the outToken is set to the new token. The refreshed token is never
func LoadToken(outToken *kcauth.Token, tokenFilePath *string) configo.ActionFunc {
	return func() error {
		if outToken == nil {
			return errors.New("outToken is nil")
		}

		cachedToken, err := loadToken(*tokenFilePath)
		if err != nil {
			// failed to load
			return err
		}

		// loaded token successfully
		if !cachedToken.IsAccessTokenExpiredIn(5 * time.Second) {
			*outToken = *cachedToken
			return nil
		}

		if cachedToken.IsRefreshTokenExpired() {
			// refresh token is also expired or failed to refresh
			return errors.New("refresh_token expired")
		}

		// only access token expired
		refreshedToken, err := internal.RefreshToken(kcauth.DefaultClientID, kcauth.DefaultClientSecret, cachedToken.RefreshToken)
		if err != nil {
			return err
		}
		// save refreshed token back to file
		err = saveToken(*tokenFilePath, refreshedToken)
		if err != nil {
			return err
		}

		// no errors -> set internally
		*outToken = *refreshedToken
		return nil
	}
}

func loadToken(tokenPath string) (*kcauth.Token, error) {
	b, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		return nil, err
	}
	token := &kcauth.Token{}
	err = json.Unmarshal(b, token)
	if err != nil {
		return nil, err
	}
	return token, nil
}
