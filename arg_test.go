package iniflag

import (
	"testing"
)

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
