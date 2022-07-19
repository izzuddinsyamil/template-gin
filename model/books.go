package model

type Book struct {
	Title       string `bson:"title" json:"title"`
	ReleaseYear int    `bson:"release_year" json:"release_year"`
}
