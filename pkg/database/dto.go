package database

type MovieDto struct {
	Title    string `json:"title"`
	Year     int    `json:"year"`
	Genre    string `json:"genre"`
	Desc     string `json:"description"`
	Director string `json:"director"`
	Actors   string `json:"actors"`
}

func (m *MovieDto) ToMovie() Movie {
	var movie Movie
	movie.Title = m.Title
	movie.Year = m.Year
	movie.Genre = m.Genre
	movie.Desc = m.Desc
	movie.Director = m.Director
	movie.Actors = m.Actors
	return movie
}

type ReviewDto struct {
	Title   string `json:"title"`
	Review  string `json:"review"`
	Rating  int    `json:"rating"`
	Author  string `json:"author"`
	MovieId uint   `json:"movieId"`
}

func (r *ReviewDto) ToReview() Review {
	var review Review
	review.Title = r.Title
	review.Review = r.Review
	review.Rating = r.Rating
	review.Author = r.Author
	review.MovieId = r.MovieId
	return review
}

func (r *ReviewDto) IsRatingValid() bool {
	return r.Rating >= 1 && r.Rating <= 10
}
