package browser

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/jxsl13/oidc"
	"github.com/jxsl13/oidc/login"
	disk "github.com/jxsl13/oidc/login/diskcache"
	configo "github.com/jxsl13/simple-configo"
	"github.com/mitchellh/go-homedir"
)

var (
	// DefaultRedirectURL must be one pointing to 127.0.0.1 but you may change the port.
	// This redirect URL must be an allowed url for your passed clientID, so e.g. the public client must allow for users to
	// redirect to http://127.0.0.1:16666. Do not use http://localhost:16666, you will see why.
	DefaultRedirectURL = "http://127.0.0.1:16666/"

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

func refreshToken(token *oidc.Token, oidcConfig login.OIDCConfig, issuerURL, redirectURL string) (*oidc.Token, error) {
	// token is expired
	client, err := oidc.NewClient(context.Background(), issuerURL)
	if err != nil {
		return nil, err
	}
	refresher := oidc.NewTokenRefresher(client, oidc.Config{
		ClientID:     oidcConfig.ClientID,
		ClientSecret: oidcConfig.ClientSecret,
		RedirectURL:  redirectURL,
		Scopes:       oidcConfig.Scopes,
	}, token.RefreshToken)
	return refresher.OIDCToken(context.Background())
}

func oidcLogin(clientID, clientSecret, issuerURL, redirectURL, cacheDirectory string) (*oidc.Token, error) {

	// config
	oidcConfig := login.OIDCConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Provider:     issuerURL,
		// Make sure you ask for offline_access if you want to use refresh tokens!
		Scopes: []string{"openid", "offline_access", "email", "profile"},
	}

	// protect against replay attacks
	sourceConfig := login.Config{
		NonceCheck: true,
	}
	// see also other caches e.g k8s.NewCache.
	cache := disk.NewCache(cacheDirectory, oidcConfig)

	// try getting token from file
	// if successfully fetched token, simply return it
	token, err := cache.Token()
	if err == nil && token != nil {
		if !token.IsAccessTokenExpired() {
			return token, nil
		}
		token, err = refreshToken(token, oidcConfig, issuerURL, redirectURL)
		if err != nil {
			return nil, err
		}
		err = cache.SaveToken(token)
		if err != nil {
			return nil, err
		}
		return token, err
	}

	cb, closeSrv, err := login.NewServer(redirectURL)
	if err != nil {
		return nil, err
	}
	defer closeSrv()

	// in case we want to revoke tokens, we might the _ parameter
	source, _, err := login.NewOIDCTokenSource(
		context.Background(),
		log.New(os.Stdout, "", 0),
		sourceConfig,
		cache,
		cb,
	)
	if err != nil {
		return nil, err
	}

	token, err = source.OIDCToken(context.Background())
	if err != nil {
		return nil, err
	}

	return token, nil
}

// Login uses the provided issuerURL to fetch an access token as well as a refresh token.
// usually the expected URL is something along the line sof https://keycloak.com/auth/realms/MYREALM
// accessToken returns the user's access token that can be used to call api endpoints.
// That token is usually passed in the Authorization http header like this:
// Authorization: Bearer <access_token>
func Login(accessToken, oidcURL *string) configo.ParserFunc {

	return func(value string) error {
		token, err := oidcLogin(DefaultClientID, DefaultClientSecret, *oidcURL, DefaultRedirectURL, DefaultCacheDirectory)
		if err != nil {
			return err
		}
		*accessToken = token.AccessToken
		return nil
	}
}
