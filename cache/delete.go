package cache

import (
	"os"

	configo "github.com/jxsl13/simple-configo"
)

// This function returns an action that deletes the token at the specified location
// in case that the token does not exist, nothing is deleted and no error returned.
func DeleteToken(tokenFilePath *string) configo.ActionFunc {
	return func() error {
		if !exists(*tokenFilePath) {
			return nil
		}
		return os.Remove(*tokenFilePath)
	}
}
