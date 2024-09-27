package repository

import (
	"context"
	"os"
	"testing"
	"time"

	cfg "stress-test-3-2-go/config"
	rent "stress-test-3-2-go/features/movie-rent"
	pg "stress-test-3-2-go/utils/driver/postgresql"

	// generator "stress-test-3-2-go/utils/generator"

	"github.com/jackc/pgx/v5/pgxpool"
	// "github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require"
)

var (
	repoTest rent.RepositoryInterface
	pool     *pgxpool.Pool
	ctx      context.Context
)

func TestMain(m *testing.M) {
	os.Setenv("DB_USERNAME", "dwiw")
	os.Setenv("DB_PASSWORD", "secret")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "rental_store")

	envConfig := &cfg.EnvConfig{
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}

	pool = pg.ConnectToPg(envConfig)

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	repoTest = NewRentRepository(pool, ctx)

	os.Exit(m.Run())
}
