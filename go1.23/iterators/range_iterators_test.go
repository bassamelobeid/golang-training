package iterators

import (
	"fmt"
	"iter"
	"sync"
	"testing"
)

type CanWatchDeadpoolPerson struct {
	Name string
	Age  int
}

func TestRangeIteration(t *testing.T) {
	var m sync.Map

	m.Store("alice", 18)
	m.Store("bob", 16)
	m.Store("cindy", 20)

	var canWatchDeadpool []CanWatchDeadpoolPerson

	// 1.22 implementation
	// m.Range(func(key, value any) bool {
	// 	if value.(int) >= 17 {
	// 		canWatchDeadpool = append(canWatchDeadpool, CanWatchDeadpoolPerson{Name: key.(string), Age: value.(int)})
	// 	}
	// 	return true
	// })

	// 1.23 implementation
	for key, val := range m.Range {
		if val.(int) >= 17 {
			canWatchDeadpool = append(canWatchDeadpool, CanWatchDeadpoolPerson{Name: key.(string), Age: val.(int)})
		}
	}

	//Expectation: People under 17 should not be able to watch deadpool
	for person := range canWatchDeadpool {
		if canWatchDeadpool[person].Age < 17 {
			t.Errorf("%s should not be able to watch deadpool", canWatchDeadpool[person].Name)
		}
	}

}

// Struct to represent a movie and its rating
type EpicMovie struct {
	Title  string
	Rating int
}

func (e EpicMovie) ToSeq2() iter.Seq2[string, int] {
	return func(yield func(string, int) bool) {
		yield(e.Title, e.Rating)
	}
}

func (e *EpicMovie) FromSeq2(sequence iter.Seq2[string, int]) {
	for key, value := range sequence {
		e.Title = key
		e.Rating = value
	}
}

func IterateMap[V, K any](m *sync.Map) iter.Seq2[V, K] {
	return func(yield func(V, K) bool) {
		m.Range(func(key, value any) bool {
			return yield(key.(V), value.(K))
		})
	}
}

// PrintAll prints all elements in a sequence.
func PrintAll[V, K any](s iter.Seq2[V, K]) {
	for v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func TestMovieRatingsIteration(t *testing.T) {
	var m sync.Map

	m.Store("Star Wars: Revenge of the Sith", 999)
	m.Store("Lord of the Rings", 10)
	m.Store("Game of Thrones", 8)
	m.Store("Harry Potter", 7)

	var epicMovies []EpicMovie

	// Using Seq2 to iterate over key-value pairs
	for title, rating := range IterateMap[string, int](&m) {
		epicMovies = append(epicMovies, EpicMovie{Title: title, Rating: rating})
	}

	// Expectation: All movies should have valid ratings, with Star Wars: Revenge of the Sith being an exception
	for _, movie := range epicMovies {
		if movie.Title == "Star Wars: Revenge of the Sith" && movie.Rating != 999 {
			t.Errorf("Expected Star Wars: Revenge of the Sith to have a rating of 999, but got %d", movie.Rating)
		} else if movie.Title != "Star Wars: Revenge of the Sith" && (movie.Rating > 10 || movie.Rating < 0) {
			t.Errorf("Expected a rating between 0 and 10 for %s, but got %d", movie.Title, movie.Rating)
		}
	}
}

func Reversed[V any](s []V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(s[i]) {
				return
			}
		}
	}
}

// Push Iterator => Pull Iterator
func TestReversedMoviesPullIterator(t *testing.T) {
	movies := []string{"Star Wars", "Lord of the Rings", "Game of Thrones", "Harry Potter"}

	next, stop := iter.Pull(Reversed(movies))
	defer stop()

	var epicMovies []string

	for {
		movie, ok := next()
		if !ok {
			break
		}
		epicMovies = append(epicMovies, movie)
	}

	// Expected reversed order
	expected := []string{"Harry Potter", "Game of Thrones", "Lord of the Rings", "Star Wars"}

	// Verify the results
	for i, movie := range epicMovies {
		if movie != expected[i] {
			t.Errorf("Expected %v, but got %v", expected[i], movie)
		}
	}

	// For illustrative purposes, print the results
	fmt.Println("Reversed Movies (using pull iterator):")
	for _, movie := range epicMovies {
		fmt.Println(movie)
	}
}
