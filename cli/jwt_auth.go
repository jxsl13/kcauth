package cli

import (
	"fmt"
	"path"

	configo "github.com/jxsl13/simple-configo"
	"github.com/mitchellh/go-homedir"
)

var (

	// DefaultClientID is usually a public client that doe snot require any credentials, thus the secret is empty.
	DefaultClientID = "public"

	// DefaultClientSecret is usually the public client that does not require any further configuration nor credentials.
	DefaultClientSecret = ""

	// DefaultCacheDirectory is the directory that is used to store cached tokens.
	DefaultCacheDirectory = "$HOME/.oidc_keys"
)

func init() {
	// initialize default home directory with a valid path
	home, err := homedir.Dir()
	if err != nil {
		fmt.Printf("Failed to find home directory: %v\n", err)
		return
	}
	DefaultCacheDirectory = path.Join(home, ".oidc_keys")
}

// Login logs you in via the cli workflow, prompts for username and password and fetches an oflfine token from
// the rarget keycloak and caches the access token as well as th erefresh token locally.
func Login(accessToken, issuerURL *string) configo.ParserFunc {
	return func(value string) error {
		token, err := jwtLogin(*issuerURL, DefaultClientID, DefaultClientSecret, DefaultCacheDirectory)
		if err != nil {
			return err
		}
		*accessToken = token.AccessToken
		return nil
	}
}
