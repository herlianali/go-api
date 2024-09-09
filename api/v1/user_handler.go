package v1

import (
	"encoding/json"
	"go-api/services"
	"log"
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/api/v1/users", getUsers)
	return router
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to /api/v1/users")
	users, err := services.GetUsers()
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Successfully fetched users")
	json.NewEncoder(w).Encode(users)
}
