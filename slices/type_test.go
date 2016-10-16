package slices

import (
	"flag"
	"testing"
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
	} {
		v.fn(v.v)
		if res := v.v.String(); res != v.exp {
			t.Errorf(errMsg, v.exp, res)
		}
	}
}
