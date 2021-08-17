package internal

import (
	"github.com/lestrrat-go/jwx/jwt"
)

func UrlRealmFromToken(token string) (url, realm string, err error) {
	rToken, err := jwt.Parse([]byte(token))
	if err != nil {
		return "", "", err
	}

	return ExtractRealm(rToken.Issuer())
}
