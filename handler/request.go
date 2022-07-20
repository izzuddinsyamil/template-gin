package handler

type insertBookRequest struct {
	Title       string `json:"title"`
	ReleaseYear int    `json:"release_year"`
}

type getBooksRequest struct {
	Title     string `form:"title"`
	StartYear int    `form:"start_year"`
	EndYear   int    `form:"end_year"`
}
