package kcauth

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
)

var (

	// DefaultCacheDirectory is the directory that is used to store cached tokens.
	DefaultCacheDirectory = "$HOME/.oidc_keys"

	// DefaultClientID is usually a public client that doe snot require any credentials, thus the secret is empty.
	DefaultClientID = "public"

	// DefaultClientSecret is usually the public client that does not require any further configuration nor credentials.
	DefaultClientSecret = ""
)

func init() {
	// initialize default home directory with a valid path
	home, err := homedir.Dir()
	if err != nil {
		fmt.Printf("Failed to find home directory: %v\n", err)
		return
	}
	appName := fileNameWithoutExtension(os.Args[0])

	// init
	DefaultCacheDirectory = path.Join(home, ".oidc_keys", appName)

}

// OS agnostic application name
func fileNameWithoutExtension(filePath string) string {
	fileName := filepath.Base(filePath)

	if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
		return fileName[:pos]
	}
	return fileName
}
