package factory

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/julienschmidt/httprouter"

	rentDelivery "stress-test-3-2-go/features/movie-rent/delivery"
	rentRepository "stress-test-3-2-go/features/movie-rent/repository"
	rentService "stress-test-3-2-go/features/movie-rent/service"
)

func InitFactory(router *httprouter.Router, pool *pgxpool.Pool, ctx context.Context) {
	rentRepoInterface := rentRepository.NewRentRepository(pool, ctx)
	rentServiceInterface := rentService.NewRentService(rentRepoInterface)
	rentDelivery.NewMovieRentDelivery(router, rentServiceInterface)
}
