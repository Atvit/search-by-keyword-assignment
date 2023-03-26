package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch_Get(t *testing.T) {
	s := NewSearch()

	s.Add(1, "lorem ipsum dolor sit amet")
	s.Add(2, "consectetur adipiscing elit sed do eiusmod tempor")
	s.Add(3, "ipsum eiusmod lorem tempor")

	tests := []struct {
		name     string
		keyword  string
		expected []int
	}{
		{"matches multiple documents", "lorem", []int{1, 3}},
		{"matches a single document", "consectetur", []int{2}},
		{"does not match any documents", "aliqua", []int{}},
		{"query multiple words", "lorem ipsum", []int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := s.Get(test.keyword)
			assert.ElementsMatch(t, test.expected, result)
		})
	}
}

func TestSearch_Remove(t *testing.T) {
	s := NewSearch()

	s.Add(1, "lorem ipsum dolor sit amet")
	s.Add(2, "consectetur adipiscing elit sed do eiusmod tempor")
	s.Add(3, "ipsum eiusmod lorem tempor")

	t.Run("remove index that exit in the data", func(t *testing.T) {
		expected := []int{1}
		s.Remove(3)
		result := s.Get("lorem")
		assert.ElementsMatch(t, expected, result)
	})

	t.Run("remove all indices that contain the term lorem.", func(t *testing.T) {
		expected := []int{}
		s.Remove(1)
		result := s.Get("lorem")
		assert.ElementsMatch(t, expected, result)
	})
}
