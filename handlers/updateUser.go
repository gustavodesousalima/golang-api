package handlers

import (
	"encoding/json"
	"fmt"
	"golang-api/database"
	"golang-api/models"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func UpdateUser(w http.ResponseWriter, r *http.Request){
	var params = mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var user models.Users

	json.NewDecoder(r.Body).Decode(&user)	

	user.ID = id

	_, err := database.DB.Exec("UPDATE users SET name = ?, email = ?, active = ? WHERE id = ?", user.Name, user.Email, user.Active, user.ID)

	if err != nil {
		 http.Error(w, err.Error(), http.StatusInternalServerError)
		 return
	}

	json.NewEncoder(w).Encode(user)
	fmt.Println("Usuário atualizado com sucesso")
}