package main

import (
	"fmt"
	search "search-by-keyword/search"
)

func main() {
	s := search.NewSearch()

	// Add
	s.Add(1, "lorem ipsum dolor sit amet")
	s.Add(2, "consectetur adipiscing elit sed do eiusmod tempor")
	s.Add(3, "ipsum eiusmod lorem tempor")

	// Get
	keywords := []string{
		"lorem",
		"consectetur",
		"aliqua",
		"lorem ipsum",
	}

	for _, keyword := range keywords {
		result := s.Get(keyword)
		fmt.Println(result)
	}

	// Remove
	s.Remove(3)
	result := s.Get("lorem")
	fmt.Println(result)
}
