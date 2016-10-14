package xflag

import (
	"reflect"
	"testing"
)

func TestINIConfigParse_NonExistentFile(t *testing.T) {
	if err := (&INIConfig{}).New().Parse("file_that_does_not_exist"); err == nil {
		t.Errorf("File does not exist, error expected.")
	}
}

func TestINIConfigParse_InvalidConfig(t *testing.T) {
	if err := (&INIConfig{}).New().Parse("./testdata/invalid.ini"); err == nil {
		t.Errorf("INI file is invalid, error expected.")
	}
}

func TestINIConfigParse(t *testing.T) {
	c := &INIConfig{
		body: map[string]map[string]string{},
	}
	err := c.Parse("./testdata/config1.ini")
	assertNil(t, err)
	err = c.Parse("./testdata/config2.ini")
	assertNil(t, err)
	err = c.Parse("./testdata/config3.ini")
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

func TestINIConfigGet(t *testing.T) {
	c := &INIConfig{
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
		"paths:xxx":   {Value: "${GOPATH} - ${GOPATH}", Found: true},
		"section:key": {Found: false},
		"":            {Found: false},
		"paths:zzz":   {Found: false},
	} {
		if v, f := c.Get(parseArgName(a)); v != vs.Value || f != vs.Found {
			t.Errorf(
				`Requested "%s". Expected "%s", "%v"; got "%s", "%v".`,
				a, vs.Value, vs.Found, v, f,
			)
		}
	}
}

func TestINIConfigJoin_IncorrectArgument(t *testing.T) {
	if (&INIConfig{}).Join(157) == nil {
		t.Fail()
	}
}

func assertNil(t *testing.T, err error) {
	if err != nil {
		t.Errorf(`No error expected, got "%v".`, err)
	}
}
