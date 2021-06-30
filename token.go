package kcauth

import (
	"time"
)

var (
	EpochZero = time.Unix(0, 0)
)

// Token is a combination of a JWT access token as well as a refresh token which can be used to fetch a new access token.
// This token struct is sued to cache tokens with a specified expiration  time in order to know when th eaccess token needs to be refreshed.
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`

	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at,omitempty"`
}

// IsAccessTokenExpired checks whether the access token has expired.
func (ct *Token) IsAccessTokenExpired() bool {
	if ct.AccessTokenExpiresAt.Equal(EpochZero) {
		return false
	}
	return time.Now().After(ct.AccessTokenExpiresAt)
}

//IsRefreshTokenExpired checks whether the refresh token has already expired.
// Is that's the case,the user needs to usually enter his credentials again.
func (ct *Token) IsRefreshTokenExpired() bool {
	if ct.RefreshTokenExpiresAt.Equal(EpochZero) {
		return false
	}
	return time.Now().After(ct.RefreshTokenExpiresAt)
}
