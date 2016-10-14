package xflag

import (
	"errors"
	"os"

	"github.com/thomasdezeeuw/ini"
)

// INIConfig represents an INI configuration file.
type INIConfig struct {
	body map[string]map[string]string
}

// New allocates and returns a new config.
func (c *INIConfig) New() Config {
	return &INIConfig{
		body: map[string]map[string]string{},
	}
}

// Parse opens and parses the requested configuration file.
// It may be called multiple times, the files will be joined.
func (c *INIConfig) Parse(file string) error {
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
func (c *INIConfig) Join(newC interface{}) error {
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

// Get receives a key name as input and returns associated
// value in the configuration file, if any.
// Otherwise, false is returned as the second argument.
func (c *INIConfig) Get(section, key string) (string, bool) {
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
