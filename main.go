package main

import (
	"encoding/json"
	"net/http"
	"pipebomb/server"
)

func getFilm(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	data, err := cacheData("films_"+q, FilmAPI.getFilmData, q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getFilmSources(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	data, err := cacheData("films_sources_"+q, FilmAPI.getSources, q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	server.Serve()
}