package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Movie stores the attributes of a movie
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func getEncodedJSON(movies *[]Movie) string {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}
	return fmt.Sprintf("%s", data)
}

func getEncodedIndentedJSON(movies *[]Movie) string {
	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}
	return fmt.Sprintf("%s", data)
}

func getDecodedDataFromJSON(encodedMoviesData string) []Movie {
	var decodedMovies []Movie
	if err := json.Unmarshal([]byte(encodedMoviesData), &decodedMovies); err != nil {
		log.Fatalf("JSON unmarshalling failed: %s", err)
	}
	return decodedMovies
}

func main() {
	movies := []Movie{
		Movie{
			Title: "Casablanca",
			Year:  1942,
			Color: false,
			Actors: []string{
				"Humphrey Bogart",
				"Ingrid Bergman",
			},
		},
		Movie{
			Title: "Cool Hand Luke",
			Year:  1967,
			Color: true,
			Actors: []string{
				"Paul Newman",
			},
		},
		Movie{
			Title: "Bullitt",
			Year:  1968,
			Color: true,
			Actors: []string{
				"Steve McQueen",
				"Jacqueline Bisset",
			},
		},
	}

	encodedMoviesJSON := getEncodedJSON(&movies)
	fmt.Printf("Encoded JSON:\n%v\n", encodedMoviesJSON)

	encodedMoviesIndentedJSON := getEncodedIndentedJSON(&movies)
	fmt.Printf("Encoded indented JSON:\n%v\n", encodedMoviesIndentedJSON)

	decodedMoviesData := getDecodedDataFromJSON(encodedMoviesJSON)
	fmt.Printf("%[1]T\n%[1]v\n%#[1]v\n", decodedMoviesData)
}
