package kcauth

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (

	// DefaultTokenFilePath is the full file path to the cached token.
	DefaultTokenFilePath = "$HOME/.config/kcauth/token.json"

	// DefaultAppName is used in order to have an individual key for the current application.
	DefaultAppName = "kcauth"

	// DefaultKeyringUsername is not necessarily a username, as in this case we
	// use a file name as key and the offline token will be the value
	DefaultKeyringUsername = "token.json"

	// DefaultClientID is usually a public client that doe snot require any credentials, thus the secret is empty.
	DefaultClientID = "public"

	// DefaultClientSecret is usually the public client that does not require any further configuration nor credentials.
	DefaultClientSecret = ""
)

func init() {
	// initialize default home directory with a valid path
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Failed to find home directory: %v\n", err)
		return
	}
	DefaultAppName = fileNameWithoutExtension(os.Args[0])
	cacheDir := path.Join(home, ".config", DefaultAppName)

	DefaultTokenFilePath = path.Join(cacheDir, "token.json")
}

// OS agnostic application name
func fileNameWithoutExtension(filePath string) string {
	fileName := filepath.Base(filePath)

	if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
		return fileName[:pos]
	}
	return fileName
}
