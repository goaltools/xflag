// Package ini implements a support of INI configuration files
// by the xflag package.
package ini

import (
	"flag"
	"strings"

	"github.com/conveyer/xflag/cflag/slices"
	"github.com/conveyer/xflag/config"

	"github.com/conveyer/ini"
)

// Config represents an INI configuration file.
type Config struct {
	body map[string]map[string]interface{}
}

// New allocates and returns a new config.
func (c *Config) New() config.Interface {
	return &Config{
		body: map[string]map[string]interface{}{},
	}
}

// AddFile opens and parses the requested configuration file
// and then merges it with the current config.
func (c *Config) AddFile(file string) error {
	m, err := ini.OpenFile(file)
	c.join(m)
	return err
}

// join gets a new configuration object, and joins it with c.body.
// The input map has a priority of the current configuration.
// I.e. it overrides values of c.body.
func (c *Config) join(m map[string]map[string]interface{}) {
	// Iterating over all available sections of the input config.
	for section := range m {
		// Make sure such section exists in the current config's map.
		if _, ok := c.body[section]; !ok {
			c.body[section] = map[string]interface{}{}
		}

		// Iterate over all available keys of the section and join them.
		for key := range m[section] {
			c.body[section][key] = m[section][key]
		}
	}
}

// Prepare gets a flag and initializes it with a value from the configuration
// file, if possible. If no value is associated with the flag, it is ignored.
func (c *Config) Prepare(f *flag.Flag) {
	// Split the flag name into section and key.
	section, key := parseArgName(f.Name)

	// Make sure such section exists.
	if _, ok := c.body[section]; !ok {
		return
	}

	// Make sure such key exists.
	v, ok := c.body[section][key]
	if !ok {
		return
	}

	// Process the flag depending on the value's type.
	switch v.(type) {
	case []string:
		// In case of a slice, set every value separately.
		sl := v.([]string)
		for i := range sl {
			f.Value.Set(sl[i])
		}

		// Indicate the end of input by using
		// a special EOI value.
		f.Value.Set(slices.EOI)
	default:
		// By-default a string is expected, so just set it.
		f.Value.Set(v.(string))
	}
}

// parseArgName gets a flag name, parses it, and returns a
// section name and a key of INI file.
// The format of argument name is assumed to be the following:
// section + section_separator + key.
func parseArgName(arg string) (section string, key string) {
	// If no delimiter is found means section is empty
	// and the whole argument name is a key.
	i := strings.Index(arg, *config.FlagNameSectSep)
	if i == -1 {
		return "", arg
	}

	// Otherwise, return the part before separator as section
	// and after it as a key.
	return arg[:i], arg[i+1:] // Do not include the dilimiter itself.
}
