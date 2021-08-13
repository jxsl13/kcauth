package cache

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
	"path"

	"github.com/jxsl13/kcauth"
)

func saveToken(tokenPath string, token *kcauth.Token) (*kcauth.Token, error) {
	// in order to access and list its folder content, execution permissions are required, thus 0700 instead of 0600
	var perm fs.FileMode = 0700
	cacheToken := token

	data, err := json.MarshalIndent(cacheToken, "", " ")
	if err != nil {
		return nil, err
	}
	dir := path.Dir(tokenPath)
	if !exists(dir) {
		err = os.MkdirAll(dir, perm)
		if err != nil {
			return nil, err
		}
	}
	err = ioutil.WriteFile(tokenPath, data, perm)
	if err != nil {
		return nil, err
	}
	return cacheToken, nil
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
