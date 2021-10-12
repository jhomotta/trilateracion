package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	// Get Controller instance
	//us := NewUserController()

	//r.HandleFunc("/api/satellite", GetSatellites).Methods("GET")
	r.HandleFunc("/api/satellite", CreateSatellite).Methods("POST")
	/*r.HandleFunc("/api/satellite/{id}", GetUser).Methods("GET")
	r.HandleFunc("/api/satellite/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/api/satellite/{id}", DeleteUser).Methods("DELETE")
	*/

	port := ":9000"
	log.Println("Server started at", port)
	log.Fatal(http.ListenAndServe(port,
		handlers.CORS(handlers.AllowedHeaders([]string{
			"X-Requested-With",
			"Content-Type",
			"Authorization"}),
			handlers.AllowedMethods([]string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
				"HEAD",
				"OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(r)))

}
