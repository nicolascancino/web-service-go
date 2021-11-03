package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicolascancino/web-service-go/controller"
	"github.com/nicolascancino/web-service-go/middleware"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/test", middleware.CheckDB(controller.HolaMundo)).Methods(http.MethodGet)
	router.HandleFunc("/registro", middleware.CheckDB(controller.Registro)).Methods(http.MethodPost)
	router.HandleFunc("/login", controller.Login).Methods(http.MethodPost)

	return router

}
