package cache

import (
	"encoding/json"

	"github.com/jxsl13/kcauth"
	configo "github.com/jxsl13/simple-configo"
	"github.com/zalando/go-keyring"
)

// SaveToken is an action function that takes the inToken and saves it to
// the provided file destination at tokenFilePath.
func SaveTokenInKeyring(inToken *kcauth.Token, appName, username *string) configo.ActionFunc {
	if inToken == nil {
		panic("inToken is nil")
	}

	if appName == nil {
		panic("appName is nil")
	}

	if username == nil {
		panic("username is nil")
	}

	return func() error {

		err := saveTokenInKeyring(*appName, *username, inToken)
		if err != nil {
			return err
		}
		return nil
	}
}

func saveTokenInKeyring(appName, keyringUser string, token *kcauth.Token) error {
	b, err := json.Marshal(token)
	if err != nil {
		return err
	}
	return keyring.Set(appName, keyringUser, string(b))
}
