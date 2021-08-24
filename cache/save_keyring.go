package cache

import (
	"encoding/json"
	"errors"

	"github.com/jxsl13/kcauth"
	configo "github.com/jxsl13/simple-configo"
	"github.com/zalando/go-keyring"
)

// SaveToken is an action function that takes the inToken and saves it to
// the provided file destination at tokenFilePath.
func SaveTokenInKeyring(inToken *kcauth.Token, appName, username *string) configo.ActionFunc {
	return func() error {
		if inToken == nil {
			return errors.New("inToken is nil")
		}
		err := saveTokenInKeyring(*appName, *username, inToken)
		if err != nil {
			return err
		}

		// sdo not add this to any result map, as
		// this is a pseudo option that doe snot serialize anything into a
		// configuration map
		return configo.ErrSkipUnparse
	}
}

func saveTokenInKeyring(appName, keyringUser string, token *kcauth.Token) error {
	b, err := json.Marshal(token)
	if err != nil {
		return err
	}
	return keyring.Set(appName, keyringUser, string(b))
}
