package cache

import (
	"encoding/json"
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"path"

	"github.com/jxsl13/kcauth"
	configo "github.com/jxsl13/simple-configo"
)

// SaveToken is an action function that takes the inToken and saves it to
// the provided file destination at tokenFilePath.
func SaveToken(inToken *kcauth.Token, tokenFilePath *string) configo.ActionFunc {
	return func() error {
		if inToken == nil {
			return errors.New("inToken is nil")
		}
		err := saveToken(*tokenFilePath, inToken)
		if err != nil {
			return err
		}
		// sdo not add this to any result map, as
		// this is a pseudo option that doe snot serialize anything into a
		// configuration map
		return configo.ErrSkipUnparse
	}
}

func saveToken(tokenPath string, token *kcauth.Token) error {
	// in order to access and list its folder content, execution permissions are required, thus 0700 instead of 0600
	var perm fs.FileMode = 0700

	data, err := json.MarshalIndent(token, "", " ")
	if err != nil {
		return err
	}
	dir := path.Dir(tokenPath)
	if !exists(dir) {
		err = os.MkdirAll(dir, perm)
		if err != nil {
			return err
		}
	}
	err = ioutil.WriteFile(tokenPath, data, perm)
	if err != nil {
		return err
	}
	return nil
}

// Exists reports whether the named file or directory exists.
func exists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
