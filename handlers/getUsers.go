package handlers

import (
	"encoding/json"
	"fmt"
	"golang-api/database"
	"golang-api/models"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT * FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var users []models.Users

	for rows.Next() {
		var user models.Users

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.Active)
		
		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)

	fmt.Println("Usuários retornados com sucesso")
}