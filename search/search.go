package search

import (
	"strings"
)

type Search interface {
	Add(index int, sentence string)
	Get(keyword string) []int
	Remove(index int)
}

type search struct {
	data map[int]string
	dic  map[string]map[int]bool
}

func NewSearch() Search {
	// store original key-value pairs of data to optimize the remove function.
	data := make(map[int]string)
	// dictionary of keywords
	dic := make(map[string]map[int]bool)
	return &search{
		dic:  dic,
		data: data,
	}
}

func (s *search) Add(index int, sentence string) {
	s.data[index] = sentence

	words := strings.Fields(sentence)
	for _, word := range words {
		if s.dic[word] == nil {
			s.dic[word] = make(map[int]bool)
		}

		s.dic[word][index] = true
	}
}

func (s *search) Get(keyword string) []int {
	var result []int
	// use a dictionary of keywords to find matching data.
	if word, ok := s.dic[keyword]; ok {
		for index := range word {
			result = append(result, index)
		}
	}

	return result
}

func (s *search) Remove(index int) {
	// loop through items to delete only when we are certain of the index.
	if sentence, ok := s.data[index]; ok {
		delete(s.data, index)
		words := strings.Fields(sentence)
		for _, word := range words {
			delete(s.dic[word], index)
		}
	}
}
