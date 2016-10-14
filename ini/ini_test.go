package ini

import (
	"reflect"
	"testing"
)

func TestConfigParse_NonExistentFile(t *testing.T) {
	if err := (&Config{}).New().Parse("file_that_does_not_exist"); err == nil {
		t.Errorf("File does not exist, error expected.")
	}
}

func TestConfigParse_InvalidConfig(t *testing.T) {
	if err := (&Config{}).New().Parse("./testdata/invalid.ini"); err == nil {
		t.Errorf("INI file is invalid, error expected.")
	}
}

func TestConfigParse(t *testing.T) {
	c := &Config{
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

func TestConfigGet(t *testing.T) {
	c := &Config{
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
		if v, f := c.Get(a); v != vs.Value || f != vs.Found {
			t.Errorf(
				`Requested "%s". Expected "%s", "%v"; got "%s", "%v".`,
				a, vs.Value, vs.Found, v, f,
			)
		}
	}
}

func TestConfigJoin_IncorrectArgument(t *testing.T) {
	if (&Config{}).Join(157) == nil {
		t.Fail()
	}
}

func TestParseArgName(t *testing.T) {
	for k, vs := range map[string][]string{
		"key":               {"", "key"},
		"section:key":       {"section", "key"},
		"section:":          {"section", ""},
		":key":              {"", "key"},
		":":                 {"", ""},
		"":                  {"", ""},
		"::":                {"", ":"},
		"section:some:key:": {"section", "some:key:"},
	} {
		if sec, key := parseArgName(k); sec != vs[0] || key != vs[1] {
			t.Errorf(
				`Input "%s": Expected "%s", "%s"; got "%s", "%s".`,
				k, vs[0], vs[1], sec, key,
			)
		}
	}
}

func assertNil(t *testing.T, err error) {
	if err != nil {
		t.Errorf(`No error expected, got "%v".`, err)
	}
}
