package flags_test

import (
	"flag"
	"reflect"
	"testing"
	"time"

	"github.com/conveyer/xflag/flags"
)

var (
	nilStrs = flags.String("nil[]", nil, "Nil default strings.")

	defStrs = []string{"a", "b", "c"}
	strs    = flags.String("name[]", defStrs, "A list of strings.")

	defInts = []int{1, 2, 3}
	ints    = flags.Int("int[]", defInts, "A list of ints.")

	defInt64s = []int64{1, 2, 3}
	int64s    = flags.Int64("int64[]", defInt64s, "A list of int64s.")

	defUints = []uint{1, 2, 3}
	uints    = flags.Uint("uint[]", defUints, "A list of uints.")

	defUint64s = []uint64{1, 2, 3}
	uint64s    = flags.Uint64("uint64[]", defUint64s, "A list of uint64s.")

	defFloat64s = []float64{1, 2, 3}
	float64s    = flags.Float64("float64[]", defFloat64s, "A list of float64s.")

	defBools = []bool{true, false, true}
	bools    = flags.Bool("bool[]", defBools, "A list of bools.")

	dur, _       = time.ParseDuration("1h5m0s")
	defDurations = []time.Duration{dur}
	durations    = flags.Duration("duration[]", defDurations, "A list of durations.")
)

func TestString_NilDefault(t *testing.T) {
	if strs == nil {
		t.Fail()
	}
}

func TestFuncs(t *testing.T) {
	for _, v := range [][2]interface{}{
		{defStrs, *strs},
		{defInts, *ints},
		{defInt64s, *int64s},
		{defUints, *uints},
		{defUint64s, *uint64s},
		{defFloat64s, *float64s},
		{defBools, *bools},
		{defDurations, *durations},
	} {
		if !reflect.DeepEqual(v[0], v[1]) {
			t.Errorf("Expected `%v`, got `%v`.", v[0], v[1])
		}
	}
}

func init() {
	flag.Parse()
}
