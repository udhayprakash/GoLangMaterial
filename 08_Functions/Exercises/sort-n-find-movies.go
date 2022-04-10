package main

import (
	"errors"
	"fmt"
	"sort"
)

type Rating int

const (
	G    Rating = 1
	PG          = 2
	PG13        = 3
	R           = 4
)

type Movie struct {
	Title       string
	Rating      Rating
	ReleaseYear int
}

var Movies = []Movie{
	{Title: "The Terminator", Rating: 4, ReleaseYear: 1984},
	{Title: "Casablanca", Rating: 2, ReleaseYear: 1942},
	{Title: "Star Trek", Rating: 3, ReleaseYear: 2009},
	{Title: "2001: A Space Odessy", Rating: 1, ReleaseYear: 1968},
	{Title: "Fast Times at Ridgemont High", Rating: 4, ReleaseYear: 1982},
	{Title: "Sixteen Candles", Rating: 2, ReleaseYear: 1984},
	{Title: "Bambi", Rating: 1, ReleaseYear: 1942},
	{Title: "Avengers: Endgame", Rating: 3, ReleaseYear: 2019},
	{Title: "Alien", Rating: 4, ReleaseYear: 1979},
}

func sortMovies(movies []Movie) []Movie {
	sort.Slice(movies, func(i, j int) bool {
		return movies[i].ReleaseYear < movies[j].ReleaseYear
	})
	return movies

}

func findMoviesByYear(movies []Movie, year int) ([]Movie, error) {
	var result []Movie
	for _, movie := range movies {
		if movie.ReleaseYear == year {
			result = append(result, movie)
		}
	}
	if len(result) > 0 {
		return result, nil
	} else {
		return nil, errors.New("no movie found on that release year")
	}
}

func main() {
	// write a function to sort the movies by release year, return the result, and print the results
	fmt.Println("sorted movies, by release year", sortMovies(Movies))

	// write a function to find the movies in a given release year, return the result, if no movies in given year, return an error and print the results/error.
	foundMovies, err := findMoviesByYear(Movies, 1942)
	fmt.Println("findMoviesByYear(Movies, 1942)=", foundMovies, err)

	foundMovies, err = findMoviesByYear(Movies, 1940)
	fmt.Println("findMoviesByYear(Movies, 1940)=", foundMovies, err)

}
