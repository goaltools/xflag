package xflag

// DefaultConfig defines a format of configuration file
// that will be used by the package.
var DefaultConfig Config = &INIConfig{}

// Config interface describes the methods that must be implemented
// in order to add support of an arbitrary configuration file format.
type Config interface {
	// New should allocate and return a new instance of Config.
	New() Config

	// Parse should open and parse the requested configuration file.
	// It should be possible to call it a few times, so the files
	// will be joined.
	Parse(string) error

	// Join gets a body of a new configuration file that must be merged.
	// The input value must have a priority over current Config values,
	// i.e. values of the input argument must override the ones of
	// the Config.
	Join(interface{}) error

	// Get receives 2 arguments as keys and returns a string value associated
	// with them as a result. If no values are found, an empty string and false
	// as a second result are returned.
	Get(string, string) (string, bool)
}
