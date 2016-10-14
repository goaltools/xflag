// Package ini implements a support of INI configuration files
// by the xflag package.
package ini

import (
	"errors"
	"os"
	"strings"

	"github.com/conveyer/xflag/config"

	"github.com/thomasdezeeuw/ini"
)

const argDelim = ":"

// Config represents an INI configuration file.
type Config struct {
	body map[string]map[string]string
}

// New allocates and returns a new config.
func (c *Config) New() config.Interface {
	return &Config{
		body: map[string]map[string]string{},
	}
}

// Parse opens and parses the requested configuration file.
// It may be called multiple times, the files will be joined.
func (c *Config) Parse(file string) error {
	// Trying to open the configuration file.
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	// Parsing the openned file.
	m, err := ini.Parse(f)
	if err != nil {
		return err
	}

	// Join the newly parsed file with the previous one.
	return c.Join(m)
}

// Join gets an ini.Config (maps of maps) interface, and joins it with c.body.
// The input map has a priority (it overrides values of c.body).
func (c *Config) Join(newC interface{}) error {
	m, ok := newC.(ini.Config)
	if !ok {
		return errors.New("input argument of ini.Config type expected")
	}

	// Iterating over all available sections in the config.
	for section := range m {
		// Make sure such section exists in the source map.
		if _, ok := c.body[section]; !ok {
			c.body[section] = map[string]string{}
		}

		// Iterating over all available keys of the section.
		for key := range m[section] {
			c.body[section][key] = m[section][key]
		}
	}
	return nil
}

// Get receives an argument name as input and returns associated
// value in the configuration file, if it does exist.
// Otherwise, false is returned as the second argument.
func (c *Config) Get(argName string) (string, bool) {
	// Split the argument into section and key.
	section, key := parseArgName(argName)

	// Make sure such section exists.
	if _, ok := c.body[section]; !ok {
		return "", false
	}

	// Make sure such key exist.
	v, ok := c.body[section][key]
	if !ok {
		return "", false
	}

	// If it is, return the associated value.
	return v, true
}

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
