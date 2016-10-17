// Package flags is an equivalent of the standard flag package
// but for work with slice types.
package flags

import (
	"flag"
	"time"

	"github.com/conveyer/xflag/flags/slices"
)

// String is an equivalent of flag.String but for []string value.
// It defines a slice flag with the specified name, default value,
// and usage string. The returned value is the address of a string
// slice variable that stores the value of the flag.
func String(name string, value []string, usage string) *[]string {
	p := &slices.Strings{Value: value}
	flag.Var(p, name, usage)
	return &p.Value
}

// Int is an equivalent of flag.Int but for []int value.
// It defines a slice flag with the specified name, default value,
// and usage string. The returned value is the address of an int
// slice variable that stores the value of the flag.
func Int(name string, value []int, usage string) *[]int {
	p := &slices.Ints{Value: value}
	flag.Var(p, name, usage)
	return &p.Value
}

// Int64 is an equivalent of flag.Int64 but for []int64 value.
// It defines a slice flag with the specified name, default value,
// and usage string. The returned value is the address of an int64
// slice variable that stores the value of the flag.
func Int64(name string, value []int64, usage string) *[]int64 {
	p := &slices.Int64s{Value: value}
	flag.Var(p, name, usage)
	return &p.Value
}

// Uint is an equivalent of flag.Uint but for []uint value.
// It defines a slice flag with the specified name, default value,
// and usage string. The returned value is the address of a uint
// slice variable that stores the value of the flag.
func Uint(name string, value []uint, usage string) *[]uint {
	p := &slices.Uints{Value: value}
	flag.Var(p, name, usage)
	return &p.Value
}

// Uint64 is an equivalent of flag.Uint64 but for []uint64 value.
// It defines a slice flag with the specified name, default value,
// and usage string. The returned value is the address of a uint64
// slice variable that stores the value of the flag.
func Uint64(name string, value []uint64, usage string) *[]uint64 {
	p := &slices.Uint64s{Value: value}
	flag.Var(p, name, usage)
	return &p.Value
}

// Float64 is an equivalent of flag.Float64 but for []float64 value.
// It defines a slice flag with the specified name, default value,
// and usage string. The returned value is the address of a float64
// slice variable that stores the value of the flag.
func Float64(name string, value []float64, usage string) *[]float64 {
	p := &slices.Float64s{Value: value}
	flag.Var(p, name, usage)
	return &p.Value
}

// Bool is an equivalent of flag.Bool but for []bool value.
// It defines a slice flag with the specified name, default value,
// and usage string. The returned value is the address of a bool
// slice variable that stores the value of the flag.
func Bool(name string, value []bool, usage string) *[]bool {
	p := &slices.Bools{Value: value}
	flag.Var(p, name, usage)
	return &p.Value
}

// Duration is an equivalent of flag.Duration but for []time.Duration value.
// It defines a slice flag with the specified name, default value,
// and usage string. The returned value is the address of a time.Duration
// slice variable that stores the value of the flag.
func Duration(name string, value []time.Duration, usage string) *[]time.Duration {
	p := &slices.Durations{Value: value}
	flag.Var(p, name, usage)
	return &p.Value
}
