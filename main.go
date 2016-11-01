package main

import (
	"net/http"

	"github.com/gorilla/mux"

	UserRoutes "./app/routes/user"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	UserRoutes.AssignRoutes(router)
	http.ListenAndServe(":3000", router)
}
