package config

// Interface describes the methods that must be implemented
// in order to add support of an arbitrary configuration file format.
type Interface interface {
	// New should allocate and return a new instance of Config.
	New() Interface

	// Parse should open and parse the requested configuration file.
	// It should be possible to call it a few times, so the files
	// will be joined.
	Parse(string) error

	// Join gets a body of a new configuration file that must be merged.
	// The input value must have a priority over current Config values,
	// i.e. values of the input argument must override the ones of
	// the Config.
	Join(interface{}) error

	// Get receives an argument name and returns a string value associated
	// with it as a result. If no values are found, an empty string and false
	// as a second result are returned.
	Get(string) (string, bool)
}
