// Package xflag is an abstraction around the Go's standard "flag"
// package, INI or other configuration files, and environment variables.
package xflag

import (
	"flag"
	"os"

	"github.com/conveyer/xflag/config"
	"github.com/conveyer/xflag/ini"
)

// Example:
//
//	package main
//
//	import (
//		"flag"
//		"log"
//
//		"github.com/conveyer/xflag"
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

// DefaultConfig defines a format of configuration file
// that will be used by the package.
var DefaultConfig config.Interface = &ini.Config{}

// Context represents a single instance of xflag.
// It contains available arguments and parsed configuration files.
type Context struct {
	args []string
	conf config.Interface
}

// New allocates and returns a new Context.
// Input arguments should not include the command name.
func New(args []string) *Context {
	return &Context{
		args: args,
		conf: DefaultConfig.New(),
	}
}

// Files method gets a number of INI configuration files and
// parses them. An error is returned if some of the files do not exist
// or their format is not valid.
// Every subsequent file overrides conflicting values of the previous one.
func (c *Context) Files(files ...string) error {
	for i := range files {
		if err := c.conf.AddFile(files[i]); err != nil {
			return err
		}
	}
	return nil
}

// ParseSet parses flag definitions using the following sources:
// 1. Configuration files (that may contain Environment variables);
// 2. Command argument list.
// The latter has higher priority.
func (c *Context) ParseSet(fset *flag.FlagSet) error {
	// Iterate over all available flags.
	fset.VisitAll(func(f *flag.Flag) {
		// And try to set them.
		c.conf.Prepare(f)
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
//	c := xflag.New(os.Args[1:])
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
	c := New(os.Args[1:])

	// Parse requested configuration files.
	err := c.Files(files...)
	if err != nil {
		return err
	}

	// Parse the default flag set, i.e. flag.CommandLine.
	return c.Parse()
}
