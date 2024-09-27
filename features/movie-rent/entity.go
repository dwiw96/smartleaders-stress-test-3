package movie_rent

type RentMovieRequest struct {
	Fullname        string `json:"fullname" validate:"required,min=1"`
	PhysicalAddress string `json:"physical_address" validate:"required,min=1"`
	MovieRented     string `json:"movie_rented" validate:"required,min=1"`
	Salutation      string `json:"salutation" validate:"required,oneof=Mr. Ms."`
}

type Salutations struct {
	ID    int
	Title string
}

type Customers struct {
	ID              int
	Fullname        string
	PhysicalAddress string
	SalutationsID   int
}

type MoviesList struct {
	ID        int
	MovieName string
}

type RentedMovies struct {
	ID           int
	CustomersID  int
	MoviesListID int
}

type ListRentedBooks struct {
	Fullname        string `json:"fullname" validate:"required,min=1"`
	PhysicalAddress string `json:"physical_address" validate:"required,min=1"`
	MovieRented     string `json:"movie_rented" validate:"required,min=1"`
	Salutation      string `json:"salutation" validate:"required"`
}

type ServiceInterface interface {
	RentBook(input RentMovieRequest) (res *ListRentedBooks, err error)
	InsertIntoSalutations(title string) (res *Salutations, err error)
	InsertIntoCustomers(fullname, physicalAddress string, salutationID int) (res *Customers, err error)
	InsertIntoMoviesList(MovieName string) (res *MoviesList, err error)
	InsertIntoRentedMovies(input RentedMovies) (res *RentedMovies, err error)
	// ListOfRentBook() (res *ListRentedBooks, err error)
}

type RepositoryInterface interface {
	ReadSalutation(title string) (res *Salutations, err error)
	ReadCustomers(fullname, physicalAddress string, salutationID int) (res *Customers, err error)
	ReadMoviesList(MovieName string) (res *MoviesList, err error)

	InsertIntoSalutations(title string) (*Salutations, error)
	InsertIntoCustomers(fullname, physicalAddress string, salutationID int) (*Customers, error)
	InsertIntoMoviesList(MovieName string) (*MoviesList, error)
	InsertIntoRentedMovies(input RentedMovies) (*RentedMovies, error)
}
