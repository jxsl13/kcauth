package cache

import (
	"errors"
	"time"

	"github.com/jxsl13/kcauth"
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
		refreshedToken, err := refreshToken(kcauth.DefaultClientID, kcauth.DefaultClientSecret, cachedToken.RefreshToken)
		if err != nil {
			return err
		}
		// save refreshed token back to file
		cachedToken, err = saveToken(*tokenFilePath, refreshedToken)
		if err != nil {
			return err
		}

		// no errors -> set internally
		*outToken = *cachedToken
		return nil
	}
}

// SaveToken is an action function that takes the inToken and saves it to
// the provided file destination at tokenFilePath.
func SaveToken(inToken *kcauth.Token, tokenFilePath *string) configo.ActionFunc {
	return func() error {
		if inToken == nil {
			return errors.New("inToken is nil")
		}
		_, err := saveToken(*tokenFilePath, inToken)
		if err != nil {
			return err
		}
		// sdo not add this to any result map, as
		// this is a pseudo option that doe snot serialize anything into a
		// configuration map
		return configo.ErrSkipUnparse
	}
}
