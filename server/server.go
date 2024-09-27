package api

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func SetupRouter() *httprouter.Router {
	log.Println("<- SetupRouter()")

	router := httprouter.New()

	log.Println("-> SetupRouter()")
	return router
}

func StartServer(port string, router *httprouter.Router) {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Adjust according to your needs
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Authorization"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	log.Println("start server at localhost", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
