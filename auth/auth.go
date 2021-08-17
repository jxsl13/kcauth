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

// Login is provides a full fledged login flow that tries to fetch your cached tokens
// or tries to authenticate you by providing your credentials via the cli or via your web browser that
// allows you to login at your provided issuer URL.
// issuerUrl: e.g. https://auth.example.com/auth/realms/myRealm
func Login(outToken *kcauth.Token, issuerUrl *string) configo.ActionFunc {
	// variable sthat are
	var (
		username string
		password string
	)
	return actions.Or(
		cache.LoadToken(outToken, &kcauth.DefaultTokenFilePath), // in case loading of the token fails, we want to trigger a login flow
		actions.If(HeadlessFunction(), // in case we are headless, trigger cli login flow, else oidc web browser login flow
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
			browser.Login(outToken, issuerUrl),
		),
	)
}

// SaveToken provides an action to save a token to a sane default location.
// The default token location depends on the application name and can be
// found in the variable kcauth.DefaultTokenFilePath.
func SaveToken(inToken *kcauth.Token) configo.ActionFunc {
	return cache.SaveToken(inToken, &kcauth.DefaultTokenFilePath)
}
