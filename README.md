# xflag
Package `xflag` is a hybrid configuration library that combines Go's standard
[`flag`](https://golang.org/pkg/flag/) package, INI or other configuration files,
and environment variables. Plus, provides support of slice flags.

[![GoDoc](https://godoc.org/github.com/goaltools/xflag?status.svg)](https://godoc.org/github.com/goaltools/xflag)
[![Build Status](https://travis-ci.org/goaltools/xflag.svg?branch=master)](https://travis-ci.org/goaltools/xflag)
[![Windows Build Status](https://ci.appveyor.com/api/projects/status/ee1b1c8tx7d5k2tc?svg=true)](https://ci.appveyor.com/project/alconaut/xflag)
[![Coverage](https://codecov.io/github/goaltools/xflag/coverage.svg?branch=master)](https://codecov.io/github/goaltools/xflag?branch=master)
[![Go Report Card](http://goreportcard.com/badge/goaltools/xflag?t=3)](http:/goreportcard.com/report/goaltools/xflag)

### Installation
*Use `-u` ("update") flag to make sure the latest version of package is installed.*
```bash
go get -u github.com/goaltools/xflag
```

### Basic Principles
1. Every flag has its own default value.
2. That default value can be overridden by INI or some other configuration file.
*The configuration file may contain Environment Variables, e.g. `${OPENSHIFT_PORT}`.*
3. It is possible to override values of the configuration file when running your app using flags.

### Usage
By default, [`INI` configuration](https://github.com/goaltools/ini#format) files are expected.
```go
package main

import (
	"flag"

	"github.com/goaltools/xflag"
)

var (
	someFlag = flag.String("someFlag", "default value", "Description of the flag.")

	name = flag.String("name", "John Doe", "Your full name.")
	age  = flag.Int("age", 18, "Your age.")
)

func main() {
	err := xflag.Parse("/path/to/file1.ini")
	if err != nil {
		panic(err)
	}
}
```
The program above has default `age = 18`. We can override it by adding to `/path/to/file1.ini`:
```ini
age = 55
```
So, now the value is `age = 55`.
But if we run the program above as `$ main --age 99` the value will be `age = 99`
no matter what inside the configuration file is.

#### Multiple Files
Function `xflag.Parse(...)` may get any number of paths to INI files. E.g.:
```go
xflag.Parse("file1.ini", "file2.ini", "file3.ini")
```
Every subsequent file will override values conflicting with the previous one. I.e. `file3.ini` has higher priority than
`file2.ini`. And if both contain `name = ...`, the value from `file3.ini` will be used.

#### INI Sections
INI file may contain sections, e.g.:
```ini
[user]
name = James Bond

[database]
port = 28015

[app]
addr = https://${OPENSHIFT_HOST}:${OPENSHIFT_PORT}
```
Code for use of such kind of configuration file will look as follows (note the flag names):
```go
name = flag.String("user:name", "...", "...")
port = flag.Int("database:port", 0, "...")
```
And the values can be overriden by running your app as `$ main --user:name "Jane Roe" --database:port 8888`.

#### Slice Flags
Standard `flag` package supports simple types such as `string`, `int`, etc. Package `xflag/cflag`
brings support of *complex* types including slices:
```go
var (
	names = cflag.Strings("names", []string{"John", "Joe"}, "A list of names.")
	ages  = cflag.Ints("ages", []int{16, 18}, "A list of ages.")
)
```
*A list of related functions includes:*
*`Bools`, `Durations`, `Float64s`, `Ints`, `Int64s`, `Strings`, `Uints`, `Uint64s`.*

Your configuration file may look as follows:
```ini
names[] = Name1
names[] = Name2

ages[] = 35
ages[] = 46
```

And in order to pass an array of parameters using console line, repeat the same flag name multiple times:
```sh
$ ./main --names[] James --names[] Bob
```

#### Custom Configuration Format
To add support of a custom configuration format, implement the
[`config.Interface`](https://godoc.org/github.com/conveyer/config#Interface).
Then use it as follows:
```go
package main

import (
	"os"

	"github.com/goaltools/xflag"
)

func init() {
	// Allocate a new xflag context.
	c := xflag.New(MyCustomConfig, os.Args[1:])

	// Parse the configuration files you need.
	if err := c.Files("/path/to/file1.xxx", "...", "/path/to/fileN.xxx"); err != nil {
		panic(err)
	}

	// Parse the flag set.
	if err := c.Parse(); err != nil {
		panic(err)
	}
}
```

### License
Distributed under the BSD 2-clause "Simplified" License unless otherwise noted.
