package kcauth

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	// ApplicationName is used to construct OS agnostic cached token names.
	ApplicationName = fileNameWithoutExtension(os.Args[0])
)

// ApplicationName sets the internally used application name which
// caches tokens under the specified application name
func SetApplicatioName(name string) {
	ApplicationName = name
}

func fileNameWithoutExtension(filePath string) string {
	fileName := filepath.Base(filePath)

	if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
		return fileName[:pos]
	}
	return fileName
}
