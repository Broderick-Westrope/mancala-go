package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func main() {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("PLACEHOLDER"))
		if err != nil {
			log.Error().Err(err).Msg("")
		}
	})

	err := http.ListenAndServe(":4000", router)
	if err != nil {
		log.Err(err)
	}
}
