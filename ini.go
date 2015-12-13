package iniflag

import (
	"os"
	"regexp"

	"github.com/Thomasdezeeuw/ini"
)

// Variables in a ${NAME} form inside configuration file are
// expected to be treated as ENV vars.
var envVar = regexp.MustCompile(`\${([A-Za-z0-9._\-]+)}`)

// config represents an INI configuration file.
type config struct {
	body map[string]map[string]string
}

// newConfig allocates and returns a new config.
func newConfig() *config {
	return &config{
		body: map[string]map[string]string{},
	}
}

// parse opens and parses the requested configuration file.
// It may be called multiple times, the files will be joined.
func (c *config) parse(file string) error {
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
	c.join(m)
	return nil
}

// join gets a maps of maps, and joins it with c.body.
// The input map has a priority (it overrides values of c.body).
func (c *config) join(m map[string]map[string]string) {
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
}

// get receives a key name as input and returns associated
// value in the configuration file, if any.
// Otherwise, false is returned as the second argument.
func (c *config) get(section, key string) (string, bool) {
	// Make sure such section exists.
	if _, ok := c.body[section]; !ok {
		return "", false
	}

	// Make sure such key exist.
	v, ok := c.body[section][key]
	if !ok {
		return "", false
	}

	// Replace environment variables in the received value
	// and return it.
	v = envVar.ReplaceAllStringFunc(v, func(k string) string {
		return os.Getenv(envVar.ReplaceAllString(k, "$1"))
	})
	return v, true
}
