package browser

var (
	// DefaultRedirectURL must be one pointing to 127.0.0.1 but you may change the port.
	// This redirect URL must be an allowed url for your passed clientID, so e.g. the public client must allow for users to
	// redirect to http://127.0.0.1:16666. Do not use http://localhost:16666, as Keycloak redirects you to 127.0.0.1.
	DefaultRedirectURL = "http://127.0.0.1:16666/"
)
