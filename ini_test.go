package iniflag

import (
	"fmt"
	"go/build"
	"reflect"
	"testing"
)

func TestConfigParse_NonExistentFile(t *testing.T) {
	if err := newConfig().parse("file_that_does_not_exist"); err == nil {
		t.Errorf("File does not exist, error expected.")
	}
}

func TestConfigParse_InvalidConfig(t *testing.T) {
	if err := newConfig().parse("./testdata/invalid.ini"); err == nil {
		t.Errorf("INI file is invalid, error expected.")
	}
}

func TestConfigParse(t *testing.T) {
	c := newConfig()
	err := c.parse("./testdata/config1.ini")
	assertNil(t, err)
	err = c.parse("./testdata/config2.ini")
	assertNil(t, err)
	err = c.parse("./testdata/config3.ini")
	assertNil(t, err)

	exp := map[string]map[string]string{
		"": {
			"key1": "value1",
			"key2": "value3",
		},
		"some_section1": {
			"key1": "value1",
			"key2": "value2",
		},
		"some_section2": {
			"key1": "value3",
		},
		"some_section3": {
			"key1": "value2",
		},
	}
	if !reflect.DeepEqual(exp, c.body) {
		t.Errorf("Incorrectly parsed file. Expected:\n%v.\nGot:\n%v.", exp, c.body)
	}
}

func TestConfigGet(t *testing.T) {
	c := &config{
		body: map[string]map[string]string{
			"": {
				"key1": "value1",
			},
			"paths": {
				"xxx": "${GOPATH} - ${GOPATH}",
			},
		},
	}
	for a, vs := range map[string]struct {
		Value string
		Found bool
	}{
		"key1":        {Value: "value1", Found: true},
		"paths:xxx":   {Value: fmt.Sprintf("%[1]s - %[1]s", build.Default.GOPATH), Found: true},
		"section:key": {Found: false},
		"":            {Found: false},
		"paths:zzz":   {Found: false},
	} {
		if v, f := c.get(parseArgName(a)); v != vs.Value || f != vs.Found {
			t.Errorf(
				`Requested "%s". Expected "%s", "%v"; got "%s", "%v".`,
				a, vs.Value, vs.Found, v, f,
			)
		}
	}
}

func assertNil(t *testing.T, err error) {
	if err != nil {
		t.Errorf(`No error expected, got "%v".`, err)
	}
}
