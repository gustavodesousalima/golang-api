package handlers

import (
	"encoding/json"
	"fmt"
	"golang-api/database"
	"golang-api/models"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.Users

	json.NewDecoder(r.Body).Decode(&user)

	timeNow := time.Now()
	
	user.CreatedAt = timeNow.Format("02/01/2006")

	result, err := database.DB.Exec("INSERT INTO users (name, email, createdat, active) VALUES (?, ?, ?, ?)", user.Name, user.Email, user.CreatedAt, user.Active)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.ID = int(id)
	
	json.NewEncoder(w).Encode(user)

	fmt.Println("Usuário inserido com sucesso")
}