module github.com/jxsl13/kcauth

go 1.16

require (
	github.com/Nerzal/gocloak/v8 v8.5.0
	github.com/jxsl13/oidc v0.6.3
	github.com/jxsl13/simple-configo v1.7.0
	golang.org/x/term v0.0.0-20210406210042-72f3dc4e9b72
)

replace (
	github.com/jxsl13/kcauth => ./
	github.com/jxsl13/kcauth/browser => ./browser
	github.com/jxsl13/kcauth/cli => ./cli
)
