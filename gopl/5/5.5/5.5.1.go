package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string
	Year  int  `json:"released"`
	Color bool `json:"color"`
	// Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	// movie := Movie{Title: "movie", Year: 2016, Color: true, Actors: []string{"joe", "mic"}}
	movie := Movie{Title: "movie", Year: 2016, Actors: []string{"joe", "mic"}}
	// data, err := json.Marshal(movie)
	data, err := json.MarshalIndent(movie, "", "    ")
	if err != nil {
		fmt.Printf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	toMovie := Movie{}
	if err := json.Unmarshal(data, &toMovie); err != nil {
		fmt.Printf("JSON Unmarshal failed: %s", err)
	}

	fmt.Printf("%#v\n", toMovie)
}
