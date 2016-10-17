package config

import (
	"flag"
)

var (
	// FlagNameSectSep is what is used for separating sections in flag names.
	// E.g. there is a flag name "person:firstname". If the separator is ":"
	// and INI configuration file used, the value associated with key "firstname"
	// will be extracted from [person] section of the configuration file.
	FlagNameSectSep = flag.String("flagNameSectSep", ":", "A separator of sections in flag names.")
)

// Interface describes the methods that must be implemented
// in order to add support of an arbitrary configuration file format.
type Interface interface {
	// New should allocate and return a new instance of Config.
	New() Interface

	// Parse should open and parse the requested configuration file.
	Parse(string) error

	// Join should get a new configuration file and merge it with the current one.
	// Values of the input object have priority over the values of the current configuration.
	// E.g. if the object we have looks as follows:
	//	obj:
	//		key1 = value1
	//		key2 = value2
	// and the input configuration is:
	//	obj:
	//		key2 = another_value
	//		key3 = value3
	// Join must turn the original object into:
	//	obj:
	//		key1 = value1
	//		key2 = another_value
	//		key3 = value3
	Join(interface{}) error

	// Prepare gets a flag that should be set if a value with appropriate name and
	// type is presented in the configuration file.
	// Otherwise, the flag must be ignored.
	Prepare(*flag.Flag)
}
