# kcauth is a Keycloak authentification library for CLIs

You may want to talk to different APIs from your command line interface, but lack the means of authentification.
THis library allows you to use different workflows in order to fetch an offline token from any Keycloak identity provider.

Multiple tokens are cache dlocally, an access token as well as a refresh token.
The refresh token is an offline token that has no expiration time, thus you may refresh your access token indefinitly.

Some expiration metadata is also cached locally in order to know when the cached token expires and needs to be refreshed.

You have two main workflows, one that opens a browser and redirects you after logging in to a locally short lived webserver that the CLI starts. In this process the application never has access to your credentials, as you do authenticate through your identity provider.

The second workflow is where you pass your credentials to the application and the application fetches multiple tokens from your keycloak and caches them locally. After the initial setup, the application doe snot prompt for any user credentials anymore.


# Example usage in combination with the  simple-configo library

```go
type Config struct {
	KeycloakURL     string
    OIDCToken       string
	JWTToken        string
}

func (c *Config) Name() string {
	return "my cli app"
}

func (c *Config) Options() configo.Options {
	appName := c.Name()
	cli.SetApplicatioName(appName)
	tokenFile := fmt.Sprintf("token_%s_%s", appName, DefaultClientID)

	return configo.Options{
		{
			Key:           "KEYCLOAK_URL",
			Mandatory:     true,
			Description:   "Authentication Keycloak that provides the authorization token.",
			DefaultValue:  "https://some-keycloak.com/auth/realms/my_realm",
			ParseFunction: parsers.String(&c.KeycloakURL),
		},
		{
            Key: "Browser login",
			IsPseudoOption: true,
			ParseFunction: browser.Login(
				&c.OIDCToken,
				&c.KeycloakURL,
            ),
		},
		{
			Key:            "CLI Login prompt",
			IsPseudoOption: true,
			ParseFunction: cli.Login(
				&c.JWTToken,
				&c.KeycloakURL,
                ),
		},
	}
}

```