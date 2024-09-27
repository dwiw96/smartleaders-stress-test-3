package delivery

type rentMovieResponse struct {
	Fullname        string `json:"fullname" validate:"required,min=1"`
	PhysicalAddress string `json:"physical_address" validate:"required,min=1"`
	MovieRented     string `json:"movie_rented" validate:"required,min=1"`
	Salutation      string `json:"salutation" validate:"required"`
}
