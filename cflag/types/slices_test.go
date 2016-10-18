package types

import (
	"reflect"
	"testing"
)

func TestBaseInit(t *testing.T) {
	msg := `Method initialized must return "%v", got "%v".`

	e := false
	b := &base{}
	ok := b.initialized()
	if ok {
		t.Errorf(msg, e, ok)
	}

	e = false
	b.requireInit(e)
	ok = b.initialized()
	if !ok {
		t.Errorf(msg, !e, !ok)
	}

	e = true
	b.requireInit(e)
	ok = b.initialized()
	if ok {
		t.Errorf(msg, e, ok)
	}
}

func TestStr(t *testing.T) {
	for exp, inp := range map[string]*test{
		"[]":        &test{},
		"[a]":       &test{d: []string{"a"}},
		"[a; b; c]": &test{d: []string{"a", "b", "c"}},
	} {
		if res := str(inp); res != exp {
			t.Errorf(`"%v": Expected "%s", got "%s".`, inp, exp, res)
		}
	}
}

func TestSet(t *testing.T) {
	for _, v := range []struct {
		fn  func(*testing.T, *test)
		exp []string
	}{
		{
			func(t *testing.T, st *test) {
			},
			nil,
		},
		{
			func(t *testing.T, st *test) {
				set(st, "a")
			},
			[]string{"a"},
		},
		{
			func(t *testing.T, st *test) {
				set(st, "a")
				set(st, "b")
				set(st, "c")
			},
			[]string{"a", "b", "c"},
		},
		{
			func(t *testing.T, st *test) {
				set(st, "a")
				set(st, "b")
				set(st, "c")

				set(st, EOI)

				set(st, "x")
				set(st, "y")
				set(st, "z")
			},
			[]string{"x", "y", "z"},
		},
	} {
		st := &test{}
		v.fn(t, st)
		if !reflect.DeepEqual(st.d, v.exp) {
			t.Errorf("Incorrect slice values. Expected:\n`%#v`.\nGot:\n`%#v`.", v.exp, st.d)
		}
	}
}

//
// Test object that implements a slice interface is below.
//

type test struct {
	base
	d []string
}

func (t *test) lenght() int {
	return len(t.d)
}

func (t *test) get(i int) string {
	return t.d[i]
}

func (t *test) alloc() {
	t.d = []string{}
}

func (t *test) add(v string) error {
	t.d = append(t.d, v)
	return nil
}
