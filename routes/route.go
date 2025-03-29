package routes

import (
	"log"
	"net/http"
	"weregna-backend/controllers/handlers"
	"weregna-backend/controllers/repository"
	"weregna-backend/db"
	"weregna-backend/usecases"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Router struct {
	route *mux.Router
}

func NewRouter(r *mux.Router) *Router {
	return &Router{route: r}
}

func (r *Router) RegisterRoute() {
	// Connect to DB
	db, err := db.ConnectDb()
	if err != nil {
		log.Println("Cannot connect to db")
		return
	}

	// Initialize repository, usecase, and handler
	studentRepository := repository.NewStudentRepoImpl(db)
	studentUsecase := usecases.NewStudentUsecases(studentRepository)
	studentHandler := handlers.NewStudentHandler(studentUsecase)

	// Define route prefix
	studentRoutes := r.route.PathPrefix("/api/v1").Subrouter()
	studentRoutes.Handle("/student/create", http.HandlerFunc(studentHandler.CreateStudent)).Methods("POST")
	studentRoutes.Handle("/student/get-by-email/{email}", http.HandlerFunc(studentHandler.GetStudentByEmail)).Methods("GET")
	studentRoutes.Handle("/student/get-by-id/{id:.}", http.HandlerFunc(studentHandler.GetStudentByID)).Methods("GET")
	studentRoutes.Handle("/student/update/{email}", http.HandlerFunc(studentHandler.UpdateStudent)).Methods("PUT")
	studentRoutes.Handle("/student/delete/{id}", http.HandlerFunc(studentHandler.DeleteStudent)).Methods("DELETE")
	studentRoutes.Handle("/students", http.HandlerFunc(studentHandler.GetStudents)).Methods("GET")

	log.Println("Routes registered:")
}

func (r *Router) Run(addr string) error {
	// CORS configuration to allow all origins
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                                       // Allow all origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed methods
		AllowedHeaders: []string{"Content-Type", "Authorization"},           // Allowed headers
	})

	// Wrap the mux router with CORS middleware
	handler := corsHandler.Handler(r.route)

	// Run the server with CORS enabled
	log.Println("Server running on port: ", addr)
	return http.ListenAndServe(addr, handler)
}
