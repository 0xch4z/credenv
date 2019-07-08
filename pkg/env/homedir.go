package env

import (
	"os"
	"os/user"
	"runtime"
)

var homeDir string

// guessHomeDir guesses home directory from environment variables
func guessHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDIR") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
	}
	return os.Getenv("HOME")
}

// GetHomeDir gets the home directory
func GetHomeDir() string {
	if homeDir == "" {
		usr, err := user.Current()
		if err != nil {
			homeDir = guessHomeDir()
		} else {
			homeDir = usr.HomeDir
		}
	}
	return homeDir
}
