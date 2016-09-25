package sting

import (
	"testing"
)

func TestAppend(t *testing.T) {
	s1 := "f"
	s2 := "o"
	s3 := "o"
	result := Append(s1, s2, s3)
	if result != "foo" {
		t.Error("Expected 'foo' got", result)
	}
}

func TestAppendArray(t *testing.T) {
	s1 := "f"
	_strings := []string{"o", "o", "b", "a", "r"}
	result := AppendArray(s1, _strings)
	if result != "foobar" {
		t.Error("Expected 'foobar' got", result)
	}
}

func TestAt(t *testing.T) {
	s1 := "foobar"
	i := 1
	result := At(s1, i)
	if result != "o" {
		t.Error("Expected 'o' got", result)
	}
}

func TestChars(t *testing.T) {
	s := "foobar"
	result := Chars(s)
	if !eqSlices(result, []string{"f", "o", "o", "b", "a", "r"}) {
		t.Error(`Expected {"f", "o", "o", "b", "a", "r"} got`, result)
	}
}
func eqSlices(a, b []string) bool {
	if a == nil && b == nil {
		return true;
	}

	if a == nil || b == nil {
		return false;
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestCollapseWhiteSpace(t *testing.T) {
	s := "foo      bar baz     quzz"
	result := CollapseWhiteSpace(s)
	if result != "foo bar baz quzz" {
		t.Error(`Expected "foo bar baz quzz" got`, result)
	}
}

func TestContainsAll(t *testing.T) {
	n := "foo     bar baz"
	h := []string{"foo", "baz", "bar"}
	result := ContainsAll(n, h)
	if !result {
		t.Error(`Expected "true" got`, result)
	}
}

func TestEndsWith(t *testing.T) {
	s := "foo bar"
	suffix := "BAR"
	result := EndsWith(s, suffix)
	if !result {
		t.Error(`Expected "true" got`, result)
	}
}

func TestBase64Encode(t *testing.T) {
	s := "sting"
	result := Base64Encode(s)
	if result != "c3Rpbmc=" {
		t.Error("Expected 'c3Rpbmc=' got ", result)
	}
}

func TestBase64Decode(t *testing.T) {
	s := "c3Rpbmc="
	result := Base64Decode(s)
	if result != "sting" {
		t.Error("Expected 'sting' got ", result)
	}
}

func TestFirst(t *testing.T) {
	s := "foobar"
	result := First(s, 3)
	if result != "foo" {
		t.Error("Expected 'foo' got ", result)
	}
}

func TestHead(t *testing.T) {
	s := "foobar"
	result := Head(s)
	if result != "f" {
		t.Error("Expected 'f' got ", result)
	}
}

func TestInsert(t *testing.T) {
	target := "foobarqooz"
	s := "baz"
	result := Insert(target, s, 5)
	if result != "foobarbazqooz" {
		t.Error("Expected 'foobarbazqooz' got ", result)
	}
}

func TestLeftPad(t *testing.T) {
	target := "1"
	symbol := "0"
	result := LeftPad(target, symbol, 6)
	if result != "000001" {
		t.Error("Expected '000001' got ", result)
	}
}

func TestRightPad(t *testing.T) {
	target := "1"
	symbol := "0"
	result := RightPad(target, symbol, 6)
	if result != "100000" {
		t.Error("Expected '100000' got ", result)
	}
}

func TestCamelCase(t *testing.T) {
	s := "CamelCase"
	result := CamelCase(s)
	if result != "camelCase" {
		t.Error("Expected 'camelCase' got ", result)
	}
}

func TestKebabCase(t *testing.T) {
	s := "kebab Case"
	result := KebabCase(s)
	if result != "kebab-case" {
		t.Error("Expected 'kebab-case' got", result)
	}
}

func TestSnakeCase(t *testing.T) {
	s := "snake Case"
	result := SnakeCase(s)
	if result != "snake_case" {
		t.Error("Expected 'snake_case' got", result)
	}
}