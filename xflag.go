// Package xflag is an abstraction around the Go's standard "flag"
// package, INI or other configuration files, and environment variables.
package xflag

import (
	"flag"
	"os"
	"strings"

	"github.com/goaltools/xflag/cflag/types"

	"github.com/conveyer/config"
	"github.com/conveyer/config/ini"
)

// Example:
//
//	package main
//
//	import (
//		"flag"
//		"log"
//
//		"github.com/goaltools/xflag"
//	)
//
//	var sampleFlag = flag.String("test:sample", "default value", "comment here...")
//
//	func main() {
//		err := xflag.Parse("path/to/file1.ini", "path/to/file2.ini")
//		if err != nil {
//			log.Fatalf(err)
//		}
//	}

// Context represents a single instance of xflag.
// It contains available arguments and parsed configuration files.
type Context struct {
	args []string
	conf config.Interface

	// Separator is a string that separates different objects or
	// section from key in flag names.
	// By default ":" is used as a separator if Context is allocated
	// using the New constructor.
	// Flag name with the separator may look as "mySection:myKey".
	Separator string

	// ArrLiteral is a string that is if included at the end of a flag
	// name means that the flag must be treated as an array rather than
	// as a scalar type.
	// By default "[]" is used as an array literal if Context is allocated
	// using the New constructor.
	// Flag name with the array literal may look as "mySection:myKey[]".
	ArrLiteral string
}

// New allocates and returns a new Context.
// A slice of input arguments should not include
// the command name.
func New(conf config.Interface, args []string) *Context {
	return &Context{
		args: args,
		conf: conf,

		Separator:  ":",
		ArrLiteral: "[]",
	}
}

// Files method gets a number of INI configuration files and
// parses them. An error is returned if some of the files do not exist
// or their format is not valid.
// Every subsequent file overrides conflicting values of the previous one.
func (c *Context) Files(files ...string) error {
	for i := range files {
		if err := c.conf.Join(files[i]); err != nil {
			return err
		}
	}
	return nil
}

// ParseSet parses flag definitions using the following sources:
// 1. Configuration files (that may contain Environment variables);
// 2. Command line arguments list.
// The latter has higher priority.
func (c *Context) ParseSet(fset *flag.FlagSet) error {
	// Iterate over all available flags.
	fset.VisitAll(func(f *flag.Flag) {
		// And try to initialize them using values of configuration files.
		c.process(f)
	})

	// Override the flags that are listed in the arguments.
	return fset.Parse(c.args)
}

// Parse is an equivalent of ParseSet with flag.CommandLine
// as a flag set input parameter.
func (c *Context) Parse() error {
	return c.ParseSet(flag.CommandLine)
}

// Parse is a shorthand for the following code:
//	c := xflag.New(INIConfigParser, os.Args[1:])
//	err := c.Files(files...)
//	if err != nil {
//		...
//	}
//	err = c.Parse()
//	if err != nil {
//		...
//	}
func Parse(files ...string) error {
	// Allocate a new context using os.Args as input.
	c := New(ini.New(nil), os.Args[1:])

	// Parse requested configuration files.
	err := c.Files(files...)
	if err != nil {
		return err
	}

	// Parse the default flag set, i.e. flag.CommandLine.
	return c.Parse()
}

// process receives a flag as an input argument and processes it.
func (c *Context) process(f *flag.Flag) {
	// Split the flag name into parts.
	path, arr := c.parseFlagName(f.Name)

	// Receive a value associated with the path.
	var v config.ValueInterface
	switch len(path) > 1 {
	case true:
		// If there are many elements in the path, use the first
		// one as an object path (in terms of config.Interface).
		v = c.conf.At(path[0]).Value(path[1:]...)
	default:
		// Otherwise, use all of them, if any,
		// as an element path.
		v = c.conf.Value(path...)
	}

	// Process the flag depending on the expected type.
	switch arr {
	case true:
		// Make sure a slice can be retrieved from the configuration.
		ss, ok := v.Strings()
		if !ok {
			return
		}

		// Emulate Add behaviour calling Set multiple times.
		// NOTE: This is supported by xflag/cflag package only
		// (standard flag package doesn't allow slice flags).
		for i := range ss {
			f.Value.Set(ss[i])
		}

		// Indicate the end of input by using
		// a special EOI value.
		f.Value.Set(types.EOI)
	default:
		// By default a string value is expected, so just set it.
		if s, ok := v.String(); ok {
			f.Value.Set(s)
		}
	}
}

// parseFlagName splits a flag name into a set of fragments using the
// earlier specified separator.
// The second arr argument is true if the flag name ends with an
// array literal that was expected to be specified earlier as well.
func (c *Context) parseFlagName(n string) (path []string, arr bool) {
	// Trim the array literal.
	s := strings.TrimRight(n, c.ArrLiteral)

	// Split the name using the specified separator.
	path = strings.Split(s, c.Separator)

	// Return the result.
	return path, s != n
}
