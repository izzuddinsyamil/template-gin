package model

type Book struct {
	Title       string `bson:"title"`
	ReleaseYear int    `bson:"release_year"`
}
