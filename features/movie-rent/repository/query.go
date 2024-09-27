package repository

import (
	"context"
	"fmt"

	rent "stress-test-3-2-go/features/movie-rent"

	"github.com/jackc/pgx/v5/pgxpool"
)

type rentRepository struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

func NewRentRepository(pool *pgxpool.Pool, ctx context.Context) rent.RepositoryInterface {
	return &rentRepository{
		pool: pool,
		ctx:  ctx,
	}
}

func (r *rentRepository) ReadSalutation(title string) (res *rent.Salutations, err error) {
	query := "SELECT * FROM salutations WHERE title=$1"

	var temp rent.Salutations
	err = r.pool.QueryRow(r.ctx, query, title).Scan(&temp.ID, &temp.Title)
	if err != nil {
		errMsg := fmt.Errorf("failed read salutation, err: %v", err)
		return nil, errMsg
	}

	return &temp, nil
}

func (r *rentRepository) ReadCustomers(fullname, physicalAddress string, salutationID int) (res *rent.Customers, err error) {
	query := "SELECT * FROM customers WHERE fullnames = $1 AND physical_address = $2"

	var customers rent.Customers
	row := r.pool.QueryRow(r.ctx, query, fullname, physicalAddress)
	err = row.Scan(&customers.ID, &customers.Fullname, &customers.PhysicalAddress, &customers.SalutationsID)
	if err != nil {
		errMsg := fmt.Errorf("failed read customers, err: %v", err)
		return nil, errMsg
	}

	return &customers, nil
}
func (r *rentRepository) ReadMoviesList(MovieName string) (res *rent.MoviesList, err error) {
	query := "SELECT * FROM movies_list WHERE movies_name = $1"

	var movies rent.MoviesList
	err = r.pool.QueryRow(r.ctx, query, MovieName).Scan(&movies.ID, &movies.MovieName)
	if err != nil {
		errMsg := fmt.Errorf("failed read movies list, err: %v", err)
		return nil, errMsg
	}

	return &movies, nil
}

func (r *rentRepository) InsertIntoSalutations(title string) (*rent.Salutations, error) {
	query := "INSERT INTO salutations(title) VALUES($1) RETURNING *"

	var res rent.Salutations
	err := r.pool.QueryRow(r.ctx, query, title).Scan(&res.ID, &res.Title)
	if err != nil {
		errMsg := fmt.Errorf("failed insert into salutation, err: %v", err)
		return nil, errMsg
	}

	return &res, nil
}
func (r *rentRepository) InsertIntoCustomers(fullname, physicalAddress string, salutationID int) (*rent.Customers, error) {
	query := "INSERT INTO customers(fullnames, physical_address, salutations_id) VALUES($1, $2, $3) RETURNING *"

	var res rent.Customers
	row := r.pool.QueryRow(r.ctx, query, fullname, physicalAddress, salutationID)
	err := row.Scan(&res.ID, &res.Fullname, &res.PhysicalAddress, &res.SalutationsID)
	if err != nil {
		errMsg := fmt.Errorf("failed insert into customers, err: %v", err)
		return nil, errMsg
	}

	return &res, nil
}
func (r *rentRepository) InsertIntoMoviesList(MovieName string) (*rent.MoviesList, error) {
	query := "INSERT INTO movies_list(movies_name) VALUES($1) RETURNING *"

	var res rent.MoviesList
	err := r.pool.QueryRow(r.ctx, query, MovieName).Scan(&res.ID, &res.MovieName)
	if err != nil {
		errMsg := fmt.Errorf("failed insert into movies list, err: %v", err)
		return nil, errMsg
	}

	return &res, nil
}

func (r *rentRepository) InsertIntoRentedMovies(input rent.RentedMovies) (*rent.RentedMovies, error) {
	query := "INSERT INTO rented_movies(customers_id, movies_list_id) VALUES($1, $2) RETURNING *"

	var res rent.RentedMovies
	row := r.pool.QueryRow(r.ctx, query, input.CustomersID, input.MoviesListID)
	err := row.Scan(&res.ID, &res.CustomersID, &res.MoviesListID)
	if err != nil {
		errMsg := fmt.Errorf("failed insert into rented movie, err: %v", err)
		return nil, errMsg
	}

	return &res, nil
}

func (r *rentRepository) ListOfRentedmovies() (*[]rent.ListRentedBooks, error) {
	query := `
		SELECT
			c.fullnames AS "FULLNAMES",
			c.physical_address AS "PHYSICAL_ADDRESS",
			STRING_AGG(ml.movies_name, ', ' ORDER BY ml.id) AS "MOVIES_RENTED",
			s.title AS "SALUTATION"
		FROM
			rented_movies rm
		INNER JOIN
			customers c ON rm.customers_id = c.id
		INNER JOIN
			movies_list ml ON rm.movies_list_id = ml.id
		INNER JOIN
			salutations s ON c.salutations_id = s.id
		GROUP BY
			c.fullnames, c.physical_address, s.title
		ORDER BY c.fullnames ASC;
	`

	var res []rent.ListRentedBooks
	rows, err := r.pool.Query(r.ctx, query)
	if err != nil {
		errMsg := fmt.Errorf("failed get list of rented movie, err: %v", err)
		return nil, errMsg
	}

	for rows.Next() {
		var temp rent.ListRentedBooks
		err := rows.Scan(&temp.Fullname, &temp.PhysicalAddress, &temp.MovieRented, &temp.Salutation)
		if err != nil {
			errMsg := fmt.Errorf("failed get list of rented movie, err: %v", err)
			return nil, errMsg
		}

		res = append(res, temp)
	}

	return &res, nil
}
