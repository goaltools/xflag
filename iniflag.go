// Package iniflag is an abstraction around standard go's flag,
// environment variables, and ini configuration reader.
package iniflag

import (
	"flag"
	"os"
)

// Example:
//
//	package main
//
//	import (
//		"flag"
//		"log"
//
//		"github.com/colegion/contrib/configs/iniflag"
//	)
//
//	var sampleFlag = flag.String("test:sample", "default value", "comment here...")
//
//	func main() {
//		err := iniflag.Parse("path/to/file1.ini", "path/to/file2.ini")
//		if err != nil {
//			log.Fatalf(err)
//		}
//	}

// Context represents a single instance of iniflag.
// It contains available arguments and parsed configuration files.
type Context struct {
	args []string
	conf *config
}

// New allocates and returns a new Context.
// Arguments should not include the command name.
func New(args []string) *Context {
	return &Context{
		args: args,
		conf: newConfig(),
	}
}

// Files gets a number of INI configuration files and
// parses them. An error is returned if some file does not exist
// or it is of invalid format.
// Every subsequent file overrides conflicting values of the previous one.
func (c *Context) Files(files ...string) error {
	for i := range files {
		if err := c.conf.parse(files[i]); err != nil {
			return err
		}
	}
	return nil
}

// ParseSet parses flag definitions using the following sources:
// 1. INI configuration files (that supports Environment variables);
// 2. argument list.
// The latter has higher priority.
func (c *Context) ParseSet(fset *flag.FlagSet) error {
	// Iterate over all available flags.
	fset.VisitAll(func(f *flag.Flag) {
		// Check whether we have such flag in the configuration files.
		if v, ok := c.conf.get(parseArgName(f.Name)); ok {
			// If so, use it.
			f.Value.Set(v)
		}
	})

	// Override the flags that are listed in the arguments.
	return fset.Parse(c.args)
}

// Parse is an equivalent of ParseSet with flag.CommandLine
// as a flag set input parameter.
func (c *Context) Parse() error {
	return c.ParseSet(flag.CommandLine)
}

// Parse is a shorthand for the following:
//	c := iniflag.New(os.Args[1:])
//	err := c.Files(files...)
//	AssertNil(err)
//	err = c.Parse()
//	AssertNil(err)
func Parse(files ...string) error {
	// Allocate a new context using os.Args as input.
	c := New(os.Args[1:])

	// Parse requested configuration files.
	err := c.Files(files...)
	if err != nil {
		return err
	}

	// Parse the default flag set, i.e. flag.CommandLine.
	return c.Parse()
}
