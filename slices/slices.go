// Package slices defines slice flag types. So, not only
// simple types (e.g. string, int, float64) can be used as flags
// but their slices, too. E.g. []string, []int, etc.
// Set methods of slice types work pretty much like Add.
// To make them behave as Set, call Set(EOI) before setting
// an actual value.
package slices

const (
	// EOI (end of input) is a special value of a string that can be passed
	// to the Set method of any type to mark it as uninitialized.
	// This is necessary as Set in current implementation works as Add. I.e.:
	//	Set("a")
	//	Set("b")
	//	Set("c")
	// The code above will produce []string{"a", "b", "c"},
	// not just []string{"c"}.
	// But sometimes original behaviour of Set is required. E.g.
	// we inserted "a", "b", and "c" first and now
	// want to redefine the values to "x", "y", "z".
	// To do it, the following call is required after the code above:
	//	Set(EOI)
	// And then we can readd all the required values:
	//	Set("x")
	//	Set("y")
	//	Set("z")
	EOI = "\t\n\"\000\"\n\t"
)

// slice is an interface that defines methods that
// every slice type must implement.
type slice interface {
	lenght() int
	get(i int) string

	alloc()
	add(val string) error

	initialized() bool
	requireInit(bool)
}

// base is a type that is wrapped by every real slice
// type of the package. It provides basic fields
// and methods.
type base struct {
	inited bool
}

// initialized is a getter of the "inited" field.
func (b *base) initialized() bool {
	return b.inited
}

// requireInit sets an "initialized" mark so the
// slice can be reinitialized.
func (b *base) requireInit(yes bool) {
	b.inited = !yes
}

// str gets a slice and returns it in a human
// readable format.
func str(s slice) string {
	// If there are no elements in the slice,
	// return nothing.
	l := s.lenght()
	if l == 0 {
		return "[]"
	}

	// Otherwise, prepare a list and return it.
	res := s.get(0)
	for i := 1; i < l; i++ {
		res += "; " + s.get(i)
	}
	return "[" + res + "]"
}

func set(s slice, v string) error {
	// Check whether the end of the input
	// is stated.
	if v == EOI {
		s.requireInit(true)
		return nil
	}

	// If the slice is marked as uninitialized,
	// reallocate it, and set as initialized.
	if !s.initialized() {
		s.requireInit(false)
		s.alloc()
	}

	// Add a new value.
	return s.add(v)
}
