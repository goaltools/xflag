package slices

import (
	"flag"
	"testing"
	"time"
)

const errMsg = "Expected `%v`, got `%v`."

func TestSetString(t *testing.T) {
	for _, v := range []struct {
		v   flag.Value
		fn  func(v flag.Value)
		exp string
	}{
		{
			&Strings{Value: []string{"a", "b", "c"}},
			func(v flag.Value) {
				v.Set("x")
				v.Set("y")
				v.Set("z")
			},
			"[x; y; z]",
		},
		{
			&Ints{Value: []int{1, 2, 3}},
			func(v flag.Value) {
				v.Set("4")
				v.Set("5")
				v.Set("6")
			},
			"[4; 5; 6]",
		},
		{
			&Int64s{Value: []int64{1, 2, 3}},
			func(v flag.Value) {
				v.Set("4")
				v.Set("5")
				v.Set("6")
			},
			"[4; 5; 6]",
		},
		{
			&Uints{Value: []uint{1, 2, 3}},
			func(v flag.Value) {
				v.Set("4")
				v.Set("5")
				v.Set("6")
			},
			"[4; 5; 6]",
		},
		{
			&Uint64s{Value: []uint64{1, 2, 3}},
			func(v flag.Value) {
				v.Set("4")
				v.Set("5")
				v.Set("6")
			},
			"[4; 5; 6]",
		},
		{
			&Float64s{Value: []float64{1, 2, 3}},
			func(v flag.Value) {
				v.Set("4.0")
				v.Set("5.0")
				v.Set("6.0")
			},
			"[4; 5; 6]",
		},
		{
			&Bools{Value: []bool{true, false, true}},
			func(v flag.Value) {
				v.Set("0")
				v.Set("false")
				v.Set("1")
			},
			"[false; false; true]",
		},
		{
			&Durations{Value: []time.Duration{}},
			func(v flag.Value) {
				v.Set("1ms")
				v.Set("1m")
				v.Set("8h")
			},
			"[1ms; 1m0s; 8h0m0s]",
		},
	} {
		v.fn(v.v)
		if res := v.v.String(); res != v.exp {
			t.Errorf(errMsg, v.exp, res)
		}
	}
}

func TestAdd_IncorrectInput(t *testing.T) {
	for inp, obj := range map[string]slice{
		"incorrect_int":      &Ints{},
		"incorrect_int64":    &Int64s{},
		"incorrect_uint":     &Uints{},
		"incorrect_uint64":   &Uint64s{},
		"incorrect_float64":  &Float64s{},
		"incorrect_bool":     &Bools{},
		"incorrect_duration": &Durations{},
	} {
		if err := obj.add(inp); err == nil {
			t.Errorf(`"%s": Error expected, got nil.`, inp)
		}
	}
}
