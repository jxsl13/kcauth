package cache

import (
	"errors"

	configo "github.com/jxsl13/simple-configo"
	"github.com/zalando/go-keyring"
)

// DeleteTokenFromKeyring returns a simple-configo ActionFunc
// that deletes the saved token from the provided location appName -> username -> token
func DeleteTokenFromKeyring(appName, username *string) configo.ActionFunc {
	return func() error {
		err := keyring.Delete(*appName, *username)

		// we do not care if the token was found or not
		if errors.Is(err, keyring.ErrNotFound) {
			return nil
		}
		return err
	}
}
