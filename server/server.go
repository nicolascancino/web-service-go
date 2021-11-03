package server

import (
	"log"
	"net/http"
	"os"

	"github.com/nicolascancino/web-service-go/routers"
	"github.com/rs/cors"
)

func Start() {

	router := routers.NewRouter()
	handler := cors.AllowAll().Handler((router))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))
}
