package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "HW4/api/app/model"
	repositories "HW4/api/app/repo"

	"github.com/gorilla/mux"
)

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.GetAllAuthors())
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {

	var author model.Author
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&author); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	var authorDb model.Author = repositories.CreateAuthor(author)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authorDb)
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {

	var author model.Author
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Id hatalı")
		return

	}
	w.WriteHeader(http.StatusOK)
	author = repositories.GetAuthor(id)
	json.NewEncoder(w).Encode(author)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application-json")
	var author model.Author
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&author)
	var authorDb model.Author = repositories.UpdateAuthor(author)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authorDb)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Id hatalı")
		return
	}
	w.WriteHeader(http.StatusOK)
	repositories.DeleteAuthor(id)
	json.NewEncoder(w).Encode("Deleted")
}
