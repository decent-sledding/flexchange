package main


import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"excan/config"
	"excan/internal/api"
)


func main() {
	var conf = config.NewConfig();
	serve_addr := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)

	mainRouter := chi.NewRouter()
	mainRouter.Mount("/api", api.ApiRouter());

	fmt.Printf("Listening on: http://%s\n", serve_addr)
	http.ListenAndServe(serve_addr, mainRouter)
}