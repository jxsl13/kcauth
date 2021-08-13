package cache

import (
	"encoding/json"
	"io/ioutil"

	"github.com/jxsl13/kcauth"
)

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
