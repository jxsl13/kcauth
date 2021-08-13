package browser

import (
	"errors"

	"github.com/jxsl13/oidc"
	"github.com/jxsl13/oidc/login"
)

// dummy cache to satisfy the external library
type cache struct {
	token  *oidc.Token
	config login.OIDCConfig
}

func newInMemoryCache(config login.OIDCConfig) *cache {
	return &cache{
		token:  nil,
		config: config,
	}
}

func (c *cache) SaveToken(token *oidc.Token) error {
	c.token = token
	return nil
}
func (c *cache) Token() (*oidc.Token, error) {
	if c.token != nil {
		return c.token, nil
	}
	return nil, errors.New("no token in memory")

}
func (c *cache) Config() login.OIDCConfig {
	return c.config
}
