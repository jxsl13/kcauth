// auth is a package that combines all of the provided login flows into a single parser function
// that caches the token or tries to login via various login flows, either via headless cli prompts
// or via a oidc login flow in your default web browser.
package auth

import (
	"github.com/jxsl13/kcauth"
	"github.com/jxsl13/kcauth/browser"
	"github.com/jxsl13/kcauth/cache"
	"github.com/jxsl13/kcauth/cli"
	configo "github.com/jxsl13/simple-configo"
	"github.com/jxsl13/simple-configo/actions"
)

var (
	// HeadlessFunction is the function that determines whether we want to use the cli login flow or the browser login flow.
	HeadlessFunction = HeadlessWindowsNoRestYes
)

// Login is provides a full fledged login flow that tries to fetch your cached tokens (either from the OS specific keyring or from file)
// or tries to authenticate you by providing your credentials via the cli or via your web browser that
// allows you to login at your provided issuer URL.
// issuerUrl: e.g. https://auth.example.com/auth/realms/myRealm
func Login(outToken *kcauth.Token, issuerUrl *string) configo.ActionFunc {
	// variable sthat are
	var (
		username string
		password string
		// isHeadlessEnv baked into returned function
		// this allows to change the headless function in between two calls of the Login function
		// the behavior of the returned ActionFunc will then depend on the set HeadlessFunction
		isHeadlessEnv bool = HeadlessFunction()
	)
	return actions.Or(
		// if it's not possible to fetch from keyring nor from file, expect user to login
		// first successful action wins
		actions.Or(
			cache.LoadTokenFromKeyring(outToken, &kcauth.DefaultAppName, &kcauth.DefaultKeyringUsername), // try to fetch from keying
			cache.LoadToken(outToken, &kcauth.DefaultTokenFilePath),                                      // in case loading of the token fails, we want to trigger a login flow
		),
		actions.If(&isHeadlessEnv, // in case we are headless, trigger cli login flow, else oidc web browser login flow
			actions.And(
				PromptText(&username),
				PromptPassword(&password),
				cli.Login(outToken, issuerUrl, &username, &password),
				func() error {
					// wipe memory after login
					username = ""
					password = ""
					return nil
				},
			),
			browser.Login(outToken, issuerUrl), // else case when we have a display
		),
	)
}

// SaveToken provides an action to save a token to a sane default location.
// It tries to save the token in the host specific keyring location or in a
// sane default file location.
// The default token location depends on the application name and can be
// found in the variable kcauth.DefaultTokenFilePath.
func SaveToken(inToken *kcauth.Token) configo.ActionFunc {
	// first successful action wins
	return actions.Or(
		cache.SaveTokenInKeyring(inToken, &kcauth.DefaultAppName, &kcauth.DefaultKeyringUsername),
		cache.SaveToken(inToken, &kcauth.DefaultTokenFilePath),
	)
}

// DeleteToken deletes the token from whatever location that it has been saved at
// either from the keyring or from a local directory path.
func DeleteToken() configo.ActionFunc {
	// we delete both, in case both tokens do not exist, nothing happens.
	return actions.And(
		cache.DeleteTokenFromKeyring(&kcauth.DefaultAppName, &kcauth.DefaultKeyringUsername),
		cache.DeleteToken(&kcauth.DefaultTokenFilePath), // does not return an error in case no such file was found, basically never errors
	)
}
