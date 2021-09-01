package utils

import (
	"os"
	"strings"
)

// ReplaceEnv replaces HOME and ~ in a string with contents of the $HOME environment variable.
func ReplaceEnvs(txt string) string {
	homeDir := os.Getenv("HOME")
	txt = strings.Replace(txt, "$HOME", homeDir, 1)
	txt = strings.Replace(txt, "~/", homeDir+"/", 1)
	return txt
}
