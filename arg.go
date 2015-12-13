package iniflag

import (
	"strings"
)

const argDelim = ":"

// parseArgName gets an argument name, parses it, and returns
// section and key of INI file.
// The format of argument name is assumed to be the following:
// section:key.
func parseArgName(arg string) (section string, key string) {
	// If no delimiter is found means section is empty
	// and the whole argument name is a key.
	i := strings.Index(arg, argDelim)
	if i == -1 {
		return "", arg
	}

	// Otherwise, return the part before delimiter as section
	// and after it as a key.
	return arg[:i], arg[i+1:] // Do not include the dilimiter itself.
}
