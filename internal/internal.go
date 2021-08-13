package internal

import (
	"fmt"
	"regexp"
	"time"

	"github.com/Nerzal/gocloak/v8"
	"github.com/jxsl13/kcauth"
)

var (
	extractRegex = regexp.MustCompile(`^(https?://[a-z0-9-\.:]{5,})/auth/realms/([^/]+)/?.*$`)
)

func ExtractRealm(url string) (domain, realm string, err error) {
	if matches := extractRegex.FindStringSubmatch(url); matches != nil {
		return matches[1], matches[2], nil
	}
	return "", "", fmt.Errorf("invalid keycloak realm url: %s", url)
}

func NewTokenFromGoCloak(token *gocloak.JWT) *kcauth.Token {
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
