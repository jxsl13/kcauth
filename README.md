# kcauth is a Keycloak authentification library for CLIs - simple-configo

*Info* This is a simple-configo extension library for CLI authentification.

You may want to talk to different APIs from your command line interface, but lack the means of authentification.
This library allows you to use different workflows in order to fetch an offline token from any Keycloak identity provider.

Multiple tokens are cached locally, an access token as well as a refresh token in a single file.
The refresh token is an offline token that has no expiration time, thus you may refresh your access token indefinitly.

Some expiration metadata is also cached locally in order to know when the cached token expires and needs to be refreshed.

You have two main workflows, one that opens a browser and redirects you after logging in to a locally short lived webserver that the CLI starts. In this process the application never has access to your credentials, as you do authenticate through your identity provider.

The second workflow is where you pass your credentials to the application and the application fetches multiple tokens from your keycloak and caches them locally. After the initial setup, the application doe snot prompt for any user credentials anymore.


The most convenient way to use this library is to use the **auth** subpackage which provides a single function that allows you to do all of the login steps necessary in order to fetch an offline token either via the browser in case you have a display of via command line prompts that take your user's credentials, fetch th ekeylcoak tokens and wipe those credentials from memory for good.

This one step authentification is shown in the example below.

# Example usage in combination with the  simple-configo library

```go
package main

import (
	"github.com/jxsl13/kcauth"
	"github.com/jxsl13/kcauth/auth"
	configo "github.com/jxsl13/simple-configo"
	"github.com/jxsl13/simple-configo/parsers"
	"github.com/jxsl13/simple-configo/unparsers"
	"github.com/manifoldco/promptui"
)

// all of the options shown in the init function are optional configuration parameters
// that can be left untouched as this library provides sane default values.
func init() {
	kcauth.DefaultTokenFilePath = "$HOME/.config/kcauth/token.json"
	kcauth.DefaultClientID = "public"
	kcauth.DefaultClientSecret = ""

	// function that determines whether we are currently in a headless environment
	// where you cannot use a web browser due to not having a display attached
	auth.HeadlessFunction = auth.HeadlessWindowsNoRestYes

	// prompt behavior of Password prompts
	auth.DefaultPasswordPrompt = promptui.Prompt{
		Label:       "Password",
		Mask:        '*',
		HideEntered: true,
	}

	// prompt behavior of Username prompts
	auth.DefaultUsernamePrompt = promptui.Prompt{
		Label:       "Username",
		HideEntered: true,
	}
}

type Config struct {
	issuerURL string
	Token     kcauth.Token
}

func (c *Config) Name() string {
	return "my cli app"
}

func (c *Config) Options() configo.Options {

	return configo.Options{
		{
			Key:             "KEYCLOAK_URL",
			Mandatory:       true,
			Description:     "Authentication Keycloak that provides the authorization token.",
			DefaultValue:    "https://some-keycloak.com/auth/realms/my_realm",
			ParseFunction:   parsers.String(&c.issuerURL),
			UnparseFunction: unparsers.String(&c.issuerURL),
		},
		{
			Key:             "User Login",
			IsPseudoOption:  true,
			ParseFunction:   auth.Login(&c.Token, &c.issuerURL),
			UnparseFunction: auth.SaveToken(&c.Token),
		},
	}
}
```