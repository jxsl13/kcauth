module github.com/jxsl13/kcauth

go 1.16

require (
	github.com/Nerzal/gocloak/v8 v8.5.0
	github.com/go-resty/resty/v2 v2.6.0
	github.com/jxsl13/oidc v0.6.3
	github.com/jxsl13/simple-configo v1.24.1
	github.com/lestrrat-go/jwx v1.2.6
	github.com/manifoldco/promptui v0.8.0
	github.com/zalando/go-keyring v0.1.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace (
	github.com/jxsl13/kcauth => ./
	github.com/jxsl13/kcauth/browser => ./browser
	github.com/jxsl13/kcauth/cli => ./cli
)
