package model

type Book struct {
	Id          string `bson:"_id,omitempty" json:"id"`
	Title       string `bson:"title" json:"title"`
	ReleaseYear int    `bson:"release_year" json:"release_year"`
}
