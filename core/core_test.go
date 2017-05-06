package core

import (
	"reflect"
	"testing"
)

func TestSortedKeys(t *testing.T) {
	input := map[string]bool{
		"a": true,
		"b": true,
		"1": true,
		"A": true,
	}
	expected := []string{"1", "A", "a", "b"}
	output := sortedKeys(input)
	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("sortedKeys returns incorrect result. expected=%+v, got %+v",
			expected, output)
	}
}
