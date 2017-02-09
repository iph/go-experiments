package properties

import (
	"strings"
	"testing"
)

const TestPeriodsInKey = `
a.b=1`

func TestPeriodsInKeyParse(t *testing.T) {
	result, _ := Parse(strings.NewReader(TestPeriodsInKey))

	expected := []Pair{
		{Key: "a.b", Value: "1"},
	}

	if !equalPairSlice(result, expected) {
		t.Errorf("Expected Properties %s, got %s", expected, result)
		t.Fail()
	}
}

const CommentsAndProperty = `
a=1
!this is a comment
#this is another comment
b=2`

func TestCommentParse(t *testing.T) {
	result, _ := Parse(strings.NewReader(CommentsAndProperty))

	expected := []Pair{
		{Key: "a", Value: "1"},
		{Key: "b", Value: "2"},
	}

	if !equalPairSlice(result, expected) {
		t.Errorf("Expected Properties %s, got %s", expected, result)
		t.Fail()
	}
}

const MultipleSeparatorProperties = `
a=1
b:2
c = 3`

func TestMultipleSeparators(t *testing.T) {
	result, _ := Parse(strings.NewReader(MultipleSeparatorProperties))

	expected := []Pair{
		{Key: "a", Value: "1"},
		{Key: "b", Value: "2"},
		{Key: "c", Value: "3"},
	}

	if !equalPairSlice(result, expected) {
		t.Errorf("Expected Properties %s, got %s", expected, result)
		t.Fail()
	}
}

func TestIncorrectValues(t *testing.T) {
	_, err := Parse(strings.NewReader("a =1"))
	if err == nil {
		t.Error("Expected Properties errors with incorrect separator")
		t.Fail()
	}

	_, err = Parse(strings.NewReader("a"))
	if err == nil {
		t.Error("Expected error since there is no corresponding value with `a`")
		t.Fail()
	}
}

// Compares two slices because I am lazy.
func equalPairSlice(a, b []Pair) bool {
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
