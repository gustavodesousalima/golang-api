package main

import (
	"fmt"
	"golang-api/database"
	"golang-api/auth"
	"net/http"
	"golang-api/handlers"
	"github.com/gorilla/mux"
)

func main() {

	database.InitDB()

	r := mux.NewRouter()	

	r.Use(auth.AuthMiddleware)

	r.HandleFunc("/users", handlers.AddUser).Methods("POST")
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	fmt.Println("Servidor rodando na porta 8080")
	
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
		return
	}

}