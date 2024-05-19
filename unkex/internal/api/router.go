package api

import (
	"net/http"
	"render"
	"github.com/go-chi/chi/v5"

	// "unkex/internal/repository"
	"unkex/internal/api"
)



func ApiRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/rate", getRate)
	router.Post("/subscribe", emailSubscribe)
	router.Get("/health", healthCheck)

	return router
}


func getRate(w http.ResponseWriter, r *http.Request) {
	// rateStats := nil
	// render.JSON(w, r, order)	
}


func emailSubscribe(w http.ResponseWriter, r *http.Request) {

}


func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}