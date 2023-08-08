package main

type movie struct {
	Title string
	Year  string
}
type star struct {
	Name      string
	Photo     string
	JobTitle  string
	BirthDate string
	Bio       string
	TopMovies []movie
}
