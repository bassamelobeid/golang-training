package iterators

import (
	"cmp"
	"fmt"
	"slices"
	"testing"
)

func TestAllIterator(t *testing.T) {
	movies := []string{"Star Wars", "Lord of the Rings", "Game of Thrones", "Harry Potter"}
	var allMovies []string

	for i, v := range slices.All(movies) {
		allMovies = append(allMovies, fmt.Sprintf("%d:%s", i, v))
	}

	expected := []string{"0:Star Wars", "1:Lord of the Rings", "2:Game of Thrones", "3:Harry Potter"}
	if !compareSlices(allMovies, expected) {
		t.Errorf("Expected %v, but got %v", expected, allMovies)
	}
}

func TestValuesIterator(t *testing.T) {
	movies := []string{"Star Wars", "Lord of the Rings", "Game of Thrones", "Harry Potter"}
	var valuesMovies []string

	for v := range slices.Values(movies) {
		valuesMovies = append(valuesMovies, v)
	}

	if !compareSlices(valuesMovies, movies) {
		t.Errorf("Expected %v, but got %v", movies, valuesMovies)
	}
}

func TestBackwardIterator(t *testing.T) {
	movies := []string{"Star Wars", "Lord of the Rings", "Game of Thrones", "Harry Potter"}
	var backwardMovies []string

	for _, v := range slices.Backward(movies) {
		backwardMovies = append(backwardMovies, v)
	}

	expected := []string{"Harry Potter", "Game of Thrones", "Lord of the Rings", "Star Wars"}
	if !compareSlices(backwardMovies, expected) {
		t.Errorf("Expected %v, but got %v", expected, backwardMovies)
	}
}

func TestCollect(t *testing.T) {
	nums := []int{11, 12, 13}
	collected := slices.Collect(slices.Values(nums))

	expected := []int{11, 12, 13}
	if !compareIntSlices(collected, expected) {
		t.Errorf("Expected %v, but got %v", expected, collected)
	}
}

func TestAppendSeq(t *testing.T) {
	nums1 := []int{11, 12}
	nums2 := []int{13, 14}
	result := slices.AppendSeq(nums1, slices.Values(nums2))

	expected := []int{11, 12, 13, 14}
	if !compareIntSlices(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestSorted(t *testing.T) {
	unsorted := []int{13, 11, 12}
	sorted := slices.Sorted(slices.Values(unsorted))

	expected := []int{11, 12, 13}
	if !compareIntSlices(sorted, expected) {
		t.Errorf("Expected %v, but got %v", expected, sorted)
	}
}

func TestSortedFunc(t *testing.T) {
	people := []person{{"cindy", 20}, {"alice", 25}, {"bob", 30}}
	compare := func(p1, p2 person) int {
		return cmp.Compare(p1.name, p2.name)
	}
	sortedPeople := slices.SortedFunc(slices.Values(people), compare)

	expected := []person{{"alice", 25}, {"bob", 30}, {"cindy", 20}}
	if !comparePeopleSlices(sortedPeople, expected) {
		t.Errorf("Expected %v, but got %v", expected, sortedPeople)
	}
}

func TestChunk(t *testing.T) {
	numsToChunk := []int{1, 2, 3, 4, 5}
	chunks := slices.Chunk(numsToChunk, 2)

	expected := [][]int{{1, 2}, {3, 4}, {5}}
	var chunkedResults [][]int
	for chunk := range chunks {
		chunkedResults = append(chunkedResults, chunk)
	}

	if !compare2DSlices(chunkedResults, expected) {
		t.Errorf("Expected %v, but got %v", expected, chunkedResults)
	}
}
