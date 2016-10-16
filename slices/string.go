package slices

// Strings represents a slice of string values.
type Strings struct {
	base
	Value []string
}

// Len returns a number of string elements in the slice.
func (s *Strings) Len() int {
	return len(s.Value)
}

// Get returns a string with the specified index.
func (s *Strings) Get(i int) string {
	return s.Value[i]
}

// Alloc allocates a new slice of strings.
func (s *Strings) Alloc() {
	s.Value = []string{}
}

// Add adds a new value to the slice.
func (s *Strings) Add(v string) {
	s.Value = append(s.Value, v)
}

// String returns the type in a human readable format.
func (s *Strings) String() string {
	return str(s)
}

// Set gets a string value and adds it to the slice.
func (s *Strings) Set(v string) error {
	set(s, v)
	return nil
}
