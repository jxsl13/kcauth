package cache

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/jxsl13/kcauth"
	"github.com/jxsl13/kcauth/internal"
	configo "github.com/jxsl13/simple-configo"
	"github.com/zalando/go-keyring"
)

// This function returns a Simple-Configo ParserFunc that may be executed in a simple-configo context.
// The returned function tries to read the json formated token file at the given location tokenFilePath
// in case the token is expired but the refresh token still active, the whole token is refreshed before
// the outToken is set to the new token. The refreshed token is never
func LoadTokenFromKeyring(outToken *kcauth.Token, appName, username *string) configo.ActionFunc {
	return func() error {
		if outToken == nil {
			return errors.New("outToken is nil")
		}

		cachedToken, err := loadTokenFromKeyring(*appName, *username)
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
		err = saveTokenInKeyring(*appName, *username, refreshedToken)
		if err != nil {
			return err
		}

		// no errors -> set internally
		*outToken = *refreshedToken
		return nil
	}
}

func loadTokenFromKeyring(appName, keyringUser string) (*kcauth.Token, error) {
	str, err := keyring.Get(appName, keyringUser)
	if err != nil {
		return nil, err
	}
	token := &kcauth.Token{}
	err = json.Unmarshal([]byte(str), token)
	if err != nil {
		return nil, err
	}
	return token, nil
}
