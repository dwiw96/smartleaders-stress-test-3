package main

import (
	"context"
	"log"
	config "stress-test-3-2-go/config"
	server "stress-test-3-2-go/server"

	factory "stress-test-3-2-go/factory"
	postgres "stress-test-3-2-go/utils/driver/postgresql"
	"time"
)

func main() {
	log.Println("-- run stress test 3 - 2 --")
	env := config.GetEnvConfig()
	pgPool := postgres.ConnectToPg(env)
	defer pgPool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	router := server.SetupRouter()

	factory.InitFactory(router, pgPool, ctx)

	server.StartServer(env.SERVER_PORT, router)
}
