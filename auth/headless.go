package auth

import (
	"os"
	"runtime"
)

// HeadlessWindowsNoRestYes: windows systems always have a display
func HeadlessWindowsNoRestYes() bool {
	if runtime.GOOS == "windows" {
		// windows always has a display
		return false
	} else {
		// linux does not have a display
		return true
	}
}
// HeadlessWindowsNoRestDisplayEnv means that on windows we always us ethe 
func HeadlessWindowsNoRestDisplayEnv() bool {
	// windows has display
	if runtime.GOOS == "windows" {
		return false
	} else {
		// linux only if env variable is set
		return len(os.Getenv("DISPLAY")) == 0
	}
}

func HeadlessYes() bool {
	return true
}

func HeadlessNo() bool {
	return false
}
