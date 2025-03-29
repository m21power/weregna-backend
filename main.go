package main

import (
	"weregna-backend/routes"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	r := routes.NewRouter(route)
	r.RegisterRoute()
	r.Run(":8080")
}
