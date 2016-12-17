// Package ini provides a type that implements Interface of the
// "github.com/conveyer/config" for the INI configuration format.
package ini

import (
	"strings"

	"github.com/conveyer/config"

	"github.com/conveyer/ini"
)

// INI is an implementation of config.Interface for ini
// configuration files.
type INI struct {
	data    map[string]map[string]interface{}
	section *string

	// Separator is a string that separates elements of sectionPath
	// and keyPath of At and Value methods.
	// If INI type is allocated using the New constructor, "." is used
	// as a separator by default.
	Separator string

	// DefaultSection is a name of the section where Value method will
	// retrieve values from if no other sections are specified explicitly
	// by the At method.
	// If INI type is allocated using the New constructor, "" is used
	// as a default section by default.
	DefaultSection string
}

// New allocates and returns a new INI type.
func New(data map[string]map[string]interface{}) *INI {
	return &INI{data: data, Separator: ".", DefaultSection: ""}
}

// New allocates a new configuration by parsing the
// requested file and returns it.
func (c *INI) New(file string) (config.Interface, error) {
	m, err := ini.OpenFile(file)
	if err != nil {
		return nil, err
	}
	return New(m), nil
}

// Join merges a requested file with the current configuration file.
// Values of a new file must have priority over the values of the
// current configuration. E.g. if the config we have looks as follows:
//	obj:
//		key1 = value1
//		key2 = value2
// and the new input configuration is:
//	obj:
//		key2 = another_value
//		key3 = value3
// The original config must be turned into:
//	obj:
//		key1 = value1
//		key2 = another_value
//		key3 = value3
func (c *INI) Join(file string) error {
	// Open the requested configuration file and parse it.
	m, err := ini.OpenFile(file)
	if err != nil {
		return err
	}

	// If current configuration data hasn't been
	// allocated yet, do it now.
	if c.data == nil {
		c.data = map[string]map[string]interface{}{}
	}

	// Iterate over all available sections of the input config.
	for section := range m {
		// Make sure such section exists in the current config's map.
		if _, ok := c.data[section]; !ok {
			c.data[section] = map[string]interface{}{}
		}

		// Iterate over all available keys of the section and join them.
		for key := range m[section] {
			c.data[section][key] = m[section][key]
		}
	}
	return nil
}

// At defines a section where Value method will retrieve values from.
// If no input arguments are specified or no At method is called, default section
// will be used instead that is "". Multiple inputs will be joined
// using "." as separator. For illustration, there is an INI configuration:
//	key1 = value1
//	key2 = value2
//
//	[mySection]
//	key3 = value3
//
//	[some.section.name]
//	some.key.name = value4
// The code below extracts values from the described configuration:
//	// No section is specified, so the default one is used.
//	c.Value("key1") // value1
//
//	// No input arguments are received, default section is used.
//	c.At().Value("key2") // value2
//
// // Section name is specified explicitly.
//	c.At("mySection").Value("key3") // value3
//
// // Section and keys are specified as a number of arguments.
//	c.At("some", "section", "name").Value("some", "key", "name") // value4
func (c *INI) At(sectionPath ...string) config.Interface {
	config := New(c.data)
	s := strings.Join(sectionPath, c.Separator)
	config.section = &s
	return config
}

// Value retrieves a value by its key. The key is a result of Join
// method on keyPath with "." as separators. As an example, there is
// an INI configuration:
//	key1 = value1
//	some.other.key2 = value2
// To retrieve the values above the following code is used:
//	c.Value("key1") // value1
//	c.Value("some", "other", "key2") // value2
func (c *INI) Value(keyPath ...string) config.ValueInterface {
	// If section hasn't been specified explicitly, use
	// the default one.
	if c.section == nil {
		c.section = &c.DefaultSection
	}

	// Check whether the previously specified section does exist.
	if _, ok := c.data[*c.section]; !ok {
		return config.NewValue(nil)
	}

	// Prepare a key and make sure it is presented
	// in the previously specified section.
	k := strings.Join(keyPath, c.Separator)
	if v, ok := c.data[*c.section][k]; ok {
		return config.NewValue(v)
	}
	return config.NewValue(nil)
}

// Names returns a list of sections if no arguments are specified,
// or a list of keys in the specified section that is a result of
// strings.Join(sectionPath, ".").
func (c *INI) Names(sectionPath ...string) (lst []string) {
	// If no arguments are specified, return a list of sections.
	if len(sectionPath) == 0 {
		return c.sections()
	}

	// Otherwise, return a list of keys in the specified section.
	s := strings.Join(sectionPath, c.Separator)
	return c.keys(s)
}

// sections returns a list of all sections presented in the config.
func (c *INI) sections() []string {
	var i int
	lst := make([]string, len(c.data))
	for k := range c.data {
		lst[i] = k
		i++
	}
	return lst
}

// keys returns a list of all keys presented in the specified section.
func (c *INI) keys(sect string) []string {
	var i int
	lst := make([]string, len(c.data[sect]))
	for k := range c.data[sect] {
		lst[i] = k
		i++
	}
	return lst
}
