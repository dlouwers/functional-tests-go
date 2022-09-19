package functional_tests_go

import "testing"

func ThreeParameters(one string) func(two string) func(three string) string {
	return func(two string) func(three string) string {
		return func(three string) string {
			return one + " " + two + " " + three
		}
	}
}

func TestCurrying(t *testing.T) {
	res := ThreeParameters("one")("two")("three")
	if res != "one two three" {
		t.Errorf("Expected 'one two three' got '%v'", res)
	}
}
