package cli

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"time"

	gocloak "github.com/Nerzal/gocloak/v8"
)

var (
	epochZero = time.Unix(0, 0)
)

func newCachedToken(token *gocloak.JWT) *cachedToken {
	cacheToken := &cachedToken{token, epochZero, epochZero}
	cacheToken.init()
	return cacheToken
}

type cachedToken struct {
	*gocloak.JWT
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at,omitempty"`
}

func (ct *cachedToken) IsAccessTokenExpired() bool {
	if ct.AccessTokenExpiresAt.Equal(epochZero) {
		return false
	}
	return time.Now().After(ct.AccessTokenExpiresAt)
}

func (ct *cachedToken) IsRefreshTokenExpired() bool {
	if ct.RefreshTokenExpiresAt.Equal(epochZero) {
		return false
	}
	return time.Now().After(ct.RefreshTokenExpiresAt)
}

func (ct *cachedToken) init() {
	if ct.ExpiresIn > 0 {
		ct.AccessTokenExpiresAt = time.Now().Add(time.Second * time.Duration(ct.ExpiresIn))
	} else {
		ct.AccessTokenExpiresAt = epochZero
	}
	if ct.RefreshExpiresIn > 0 {
		ct.RefreshTokenExpiresAt = time.Now().Add(time.Second * time.Duration(ct.RefreshExpiresIn))
	} else {
		ct.RefreshTokenExpiresAt = epochZero
	}
}

func loadCachedToken(tokenPath string) (*cachedToken, error) {
	b, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		return nil, err
	}
	token := &cachedToken{}
	err = json.Unmarshal(b, token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func saveToken(tokenPath string, token *gocloak.JWT) (*cachedToken, error) {
	// in order to access and list its folder content, execution permissions are required, thus 0700 instead of 0600
	var perm fs.FileMode = 0700
	cacheToken := newCachedToken(token)

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
