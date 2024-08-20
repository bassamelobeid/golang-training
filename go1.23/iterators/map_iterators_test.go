package iterators

import (
	"fmt"
	"maps"
	"testing"
)

func TestAllMapIterator(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	var keyValuePairs []string

	for k, v := range maps.All(m) {
		keyValuePairs = append(keyValuePairs, fmt.Sprintf("%s:%d", k, v))
	}

	expected := []string{"a:1", "b:2", "c:3"}
	if !compareSlices(keyValuePairs, expected) {
		t.Errorf("Expected %v, but got %v", expected, keyValuePairs)
	}
}

func TestKeysMapIterator(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	var keys []string

	for k := range maps.Keys(m) {
		keys = append(keys, k)
	}

	expected := []string{"a", "b", "c"}
	if !compareSlices(keys, expected) {
		t.Errorf("Expected %v, but got %v", expected, keys)
	}
}

func TestValuesMapIterator(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	var values []int

	for v := range maps.Values(m) {
		values = append(values, v)
	}

	expected := []int{1, 2, 3}
	if !compareIntSlices(values, expected) {
		t.Errorf("Expected %v, but got %v", expected, values)
	}
}

func TestInsertMapIterator(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 12, "c": 3, "d": 4}
	maps.Insert(m1, maps.All(m2))

	expected := map[string]int{"a": 1, "b": 12, "c": 3, "d": 4}
	if !compareMaps(m1, expected) {
		t.Errorf("Expected %v, but got %v", expected, m1)
	}
}

func TestCollectMapIterator(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	collected := maps.Collect(maps.All(m1))

	expected := map[string]int{"a": 1, "b": 2, "c": 3}
	if !compareMaps(collected, expected) {
		t.Errorf("Expected %v, but got %v", expected, collected)
	}
}
