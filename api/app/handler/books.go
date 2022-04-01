package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "HW4/api/app/model"
	repositories "HW4/api/app/repo"

	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.GetAllBooks())
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	var book model.Book
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	var bookDb model.Book = repositories.CreateBook(book)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookDb)
}

func GetBook(w http.ResponseWriter, r *http.Request) {

	var book model.Book
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Id hatalı")
		return

	}
	w.WriteHeader(http.StatusOK)
	book = repositories.GetBook(id)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application-json")
	var book model.Book
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&book)
	var bookDb model.Book = repositories.UpdateBook(book)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookDb)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Id hatalı")
		return
	}
	w.WriteHeader(http.StatusOK)
	repositories.DeleteBook(id)
	json.NewEncoder(w).Encode("Deleted")
}
func GetSearchBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	name, err := params["name"]
	if err != true {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Not Found")
		return
	}
	w.WriteHeader(http.StatusOK)
	book = repositories.GetSearchBook(name)
	json.NewEncoder(w).Encode(book)
}
func BuyBookByID(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	number, _ := strconv.Atoi(params["number"])
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Stock Not Found")
		return
	}
	w.WriteHeader(http.StatusOK)
	book, err = repositories.BuyBookByID(id, number)
	json.NewEncoder(w).Encode(book)

}
