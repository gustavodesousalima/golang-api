package handlers

import(
	"net/http"
	"fmt"
	"golang-api/database"
	"golang-api/models"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
)

func DeleteUser (w http.ResponseWriter, r * http.Request){
	var params = mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var user models.Users

	json.NewDecoder(r.Body).Decode(&user)

	user.ID = id

	_, err := database.DB.Exec("DELETE FROM users WHERE id = ?", user.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
	fmt.Println("Usuário deletado com sucesso")
}