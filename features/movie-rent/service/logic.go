package service

import (
	"errors"
	// "fmt"
	// "log"

	"github.com/jackc/pgx/v5"

	rent "stress-test-3-2-go/features/movie-rent"
)

type rentService struct {
	repo rent.RepositoryInterface
}

func NewRentService(repo rent.RepositoryInterface) rent.ServiceInterface {
	return &rentService{
		repo: repo,
	}
}

func (s *rentService) RentBook(input rent.RentMovieRequest) (*rent.ListRentedBooks, error) {
	salutation, err := s.InsertIntoSalutations(input.Salutation)
	if err != nil {
		return nil, err
	}

	customer, err := s.InsertIntoCustomers(input.Fullname, input.PhysicalAddress, salutation.ID)
	if err != nil {
		return nil, err
	}

	movie, err := s.InsertIntoMoviesList(input.MovieRented)
	if err != nil {
		return nil, err
	}

	rentedMovieinput := rent.RentedMovies{
		CustomersID:  customer.ID,
		MoviesListID: movie.ID,
	}
	_, err = s.InsertIntoRentedMovies(rentedMovieinput)

	res := rent.ListRentedBooks{
		Fullname:        customer.Fullname,
		PhysicalAddress: customer.PhysicalAddress,
		MovieRented:     movie.MovieName,
		Salutation:      salutation.Title,
	}

	return &res, err
}

func (s *rentService) InsertIntoSalutations(title string) (res *rent.Salutations, err error) {
	res, err = s.repo.ReadSalutation(title)
	if res != nil {
		return res, nil
	}

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	res, err = s.repo.InsertIntoSalutations(title)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (s *rentService) InsertIntoCustomers(fullname, physicalAddress string, salutationID int) (res *rent.Customers, err error) {
	res, err = s.repo.ReadCustomers(fullname, physicalAddress, salutationID)
	if res != nil {
		return res, nil
	}

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	res, err = s.repo.InsertIntoCustomers(fullname, physicalAddress, salutationID)
	if err != nil {
		return nil, err
	}

	return
}
func (s *rentService) InsertIntoMoviesList(MovieName string) (res *rent.MoviesList, err error) {
	res, err = s.repo.ReadMoviesList(MovieName)
	if res != nil {
		return res, nil
	}

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	res, err = s.repo.InsertIntoMoviesList(MovieName)
	if err != nil {
		return nil, err
	}

	return
}
func (s *rentService) InsertIntoRentedMovies(input rent.RentedMovies) (res *rent.RentedMovies, err error) {
	res, err = s.repo.InsertIntoRentedMovies(input)
	if err != nil {
		return nil, err
	}

	return
}
