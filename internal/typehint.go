package internal

var (
	// TypeHintAccessToken is the token type hint when revoking an access token
	// this helps increase performance when revoking a troken
	TypeHintAccessToken = tokenTypeHint("access_token")

	// TypeHintRefreshToken is the token type hint when revoking a refresh token
	// this helps increase performance when revoking a troken
	TypeHintRefreshToken = tokenTypeHint("refresh_token")
)

type tokenTypeHint string
