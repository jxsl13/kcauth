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
	appName := fileNameWithoutExtension(os.Args[0])
	cacheDir := path.Join(home, ".config", appName)

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
