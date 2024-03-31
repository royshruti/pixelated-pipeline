package controllers

import (
	"encoding/json"
	"net/http"

	"devstream.in/blogs/models"
	"devstream.in/blogs/repositories"
	"github.com/gorilla/mux"
)

func EditUserData(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	keys := vars["username"]

	var updatedUser models.User
	json.NewDecoder(r.Body).Decode(&updatedUser)

	results, err := repositories.Db.Query(
		"UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4;",
		updatedUser.Name, updatedUser.Email, updatedUser.Password, keys,
	)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		err = results.Scan(
			&updatedUser.Name,
			&updatedUser.Email,
			&updatedUser.Username,
			&updatedUser.Password,
		)

		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(updatedUser)
}