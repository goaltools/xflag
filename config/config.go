// Package config defines global variables and interfaces
// that are used by configuration file parsers compatible
// with the xflag package.
package config

import (
	"flag"
)

// Separator is what is used for separating sections in flag names.
// E.g. there is a flag named "person:firstname". If the separator is ":"
// and INI configuration file used, the value associated with key "firstname"
// will be extracted from [person] section of the configuration file.
var Separator = flag.String("xflag:section.separator", ":", "A separator of sections in flag names.")

// Interface describes the methods that must be implemented
// in order to add support of an arbitrary configuration file format.
type Interface interface {
	// New should allocate and return a new instance of Config.
	New() Interface

	// AddFile should open and parse one more configuration file
	// that is expected to be merged with the current one.
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
	AddFile(string) error

	// Prepare gets a flag that should be set if a value with appropriate name and
	// type is presented in the configuration file.
	// Otherwise, the flag must be ignored.
	Prepare(*flag.Flag)
}
