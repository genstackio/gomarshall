package gomarshall

import "testing"

type TestCase struct {
	name     string
	value    interface{}
	opts     Options
	expected string
}

type yy struct {
	E map[string]int
	F *string
}
type ss struct {
	A string `json:"blaBla"`
	B bool   `json:"other_bla,omitempty"`
	c int
	D yy
	G *yy
}

func TestToJsonBytes(t *testing.T) {
	s := "my string 1"
	ss1 := ss{"Hello", true, 42, yy{}, nil}
	ss2 := ss{"world", false, 21, yy{map[string]int{"z": 34}, nil}, nil}
	for _, test := range []TestCase{
		{"integer", 12, Options{}, "12"},
		{"string", "Hello world", Options{}, "\"Hello world\""},
		{"bool", true, Options{}, "true"},
		{"bool", false, Options{}, "false"},
		{"float", 12.34, Options{}, "12.34"},
		{"map[string]string", map[string]string{"a": "b", "cc": "dd"}, Options{}, "{\"a\":\"b\",\"cc\":\"dd\"}"},
		{"nil", nil, Options{}, "null"},
		{"*string", &s, Options{}, "\"my string 1\""},
		{"struct", ss1, Options{}, "{\"D\":{\"E\":{},\"F\":null},\"G\":null,\"blaBla\":\"Hello\",\"other_bla\":true}"},
		{"*struct", &ss1, Options{}, "{\"D\":{\"E\":{},\"F\":null},\"G\":null,\"blaBla\":\"Hello\",\"other_bla\":true}"},
		{"*struct#2", &ss2, Options{}, "{\"D\":{\"E\":{\"z\":34},\"F\":null},\"G\":null,\"blaBla\":\"world\",\"other_bla\":false}"},
	} {
		t.Run(test.name, func(t *testing.T) {
			r, _ := ToJsonBytes(test.value, test.opts)
			if string(r) != test.expected {
				t.Fatalf("ToJsonBytes(%s: %s) != %s, => %s", test.name, test.value, test.expected, string(r))
			}
		})
	}
}
