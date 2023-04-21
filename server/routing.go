// pipebomb/server/routing.go
package server

import (
	"pipebomb/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// InitRouter initializes the router
func InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Home)
	r.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	r.HandleFunc("/api/films/vip/search", handlers.FilmSearch)
	r.HandleFunc("/api/films/vip/servers", handlers.FetchFilms)
	r.HandleFunc("/api/films/vip/sources", handlers.FetchFilmSources)
	r.HandleFunc("/api/series/vip/search", handlers.ShowSearch)
	r.HandleFunc("/api/series/vip/seasons", handlers.ShowSeason)
	r.HandleFunc("/api/series/vip/servers", handlers.FetchShows)
	r.HandleFunc("/api/series/vip/sources", handlers.FetchShowSources)

	return r
}