package cli

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"time"

	gocloak "github.com/Nerzal/gocloak/v8"
	"github.com/jxsl13/kcauth"
)

func newTokenFromGoCloak(token *gocloak.JWT) *kcauth.Token {
	resultToken := &kcauth.Token{
		AccessToken:           token.AccessToken,
		RefreshToken:          token.RefreshToken,
		AccessTokenExpiresAt:  kcauth.EpochZero,
		RefreshTokenExpiresAt: kcauth.EpochZero,
	}

	if token.ExpiresIn > 0 {
		resultToken.AccessTokenExpiresAt = time.Now().Add(time.Second * time.Duration(token.ExpiresIn))
	}
	if token.RefreshExpiresIn > 0 {
		resultToken.RefreshTokenExpiresAt = time.Now().Add(time.Second * time.Duration(token.RefreshExpiresIn))
	}

	return resultToken
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

func saveToken(tokenPath string, token *kcauth.Token) (*kcauth.Token, error) {
	// in order to access and list its folder content, execution permissions are required, thus 0700 instead of 0600
	var perm fs.FileMode = 0700
	cacheToken := token

	data, err := json.Marshal(cacheToken)
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
