// Package config defines a common interface for configuration parsers.
// Demonstration of the package's API when used with INI file format:
//	// Allocate a new configuration.
//	c, err := MyConfig.New("/path/to/config.ini")
//
//	// Merge a new config to the current one.
//	err = c.Join("/path/to/another/config.ini")
//
//	// Extract values from the "mySection" of INI configuration.
//	s, ok := c.At("mySection").Value("myKey").String()
//	s1 := c.At("mySection").Value("myKey1").StringDefault("default value")
package config

// Interface describes the methods that must be implemented by every
// config parser in order to be compatible with the package.
type Interface interface {
	// New should try to open, parse, and process the requested file, allocate
	// a new config, and return it if that's possible. A non-nil error is expected
	// as a second argument otherwise.
	// Argument pathPrefix specifies section / object where method Value
	// will look for values to return.
	New(file string) (Interface, error)

	// Join should merge the current configuration with the content of the
	// requested configuration file. Values of the input file should have a
	// higher priority than the values of the current config and thus
	// override them in case similar keys are used.
	Join(file string) error

	// At should specify a path to the object where Value method will retrieve
	// values from. For illustration, possible use case is below:
	//	c, _ := MyYAMLConfig.New("/path/to/config.yaml")
	//	admins := c.At("users", "admins")
	//	rootEmail := admins.Value("root", "email")
	// The code above will extract values from the following YAML:
	//	users:
	//		admins:
	//			root:
	//				email: abc@xyz.xx
	At(objectPath ...string) Interface

	// Value should return a value associated with the specified element path.
	Value(elementPath ...string) ValueInterface

	// Names should return object names. For illustration, there is a YAML:
	//	users:
	//		admins:
	//			a:
	//				...
	//			b:
	//				...
	//			c:
	//				...
	// The code below is expected to return ["a", "b", "c"]:
	//	c, _ := MyYAMLConfig.New("/path/to/config.yaml")
	//	names := c.Names("users", "admins")
	Names(objectPath ...string) []string
}

// ValueInterface describes a type of value that is expected
// to be returned from configuration.
//
// NOTE: In most cases you DO NOT have to implement this interface yourself, consider
// the use of config.Value type and particularly config.NewValue function instead.
type ValueInterface interface {
	// Interface should return an inner value as an interface{}.
	Interface() interface{}

	// String should return the value as a string or false as a second
	// argument is expected if that is not possible.
	String() (string, bool)

	// StringDefault is an equiavalent of String that should return a specified
	// default value if no conversion to string is possible.
	StringDefault(string) string

	// Strings is an equivalent of String but for []string data.
	Strings() ([]string, bool)

	// StringsDefault is an equivalent of StringDefault but for []string data.
	StringsDefault([]string) []string
}
