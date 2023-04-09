package server

import (
	handies "github.com/gorilla/handlers"
	"log"
	"net/http"
	_ "pipebomb/docs"
	"pipebomb/handlers"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// InitRouter initializes the router
func InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/api/films/vip/search", handlers.FilmSearch)
	r.HandleFunc("/api/films/vip/source", handlers.FetchFilms)

	// Add the httpSwagger middleware to your router
	r.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)

	return r
}

// Server sets up the server
func Server(r *mux.Router, port string) {
	// setup cors for router
	cors := handies.CORS(
		handies.AllowedOrigins([]string{"*"}),
		handies.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handies.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)
	log.Fatal(http.ListenAndServe(":"+port, cors(r)))
}

func StartServer(port string) {
	r := InitRouter()
	Server(r, port)
}
