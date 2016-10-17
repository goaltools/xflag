package flags_test

import (
	"flag"
	"reflect"
	"testing"

	"github.com/conveyer/xflag/flags"
)

var (
	nilStrs = flags.String("nil[]", nil, "Nil default strings.")

	defaultNames = []string{"a", "b", "c"}
	names        = flags.String("name[]", defaultNames, "A list of names")
)

func TestString_NilDefault(t *testing.T) {
	if names == nil {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	if !reflect.DeepEqual(*names, defaultNames) {
		t.Fail()
	}
}

func init() {
	flag.Parse()
}
