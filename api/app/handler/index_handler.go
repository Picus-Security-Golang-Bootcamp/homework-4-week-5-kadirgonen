package handler

import (
	"github.com/gorilla/mux"
)

func IndexRouting() *mux.Router {
	r := mux.NewRouter()
	//? <--------------------Books-------------------------->
	r.HandleFunc("/books", GetAllBooks).Methods("GET")
	r.HandleFunc("/books", CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")
	r.HandleFunc("/bookname/{name}", GetSearchBook).Methods("GET")
	//? <--------------------Authors-------------------------->
	r.HandleFunc("/authors", GetAllAuthors).Methods("GET")
	r.HandleFunc("/authors/{id}", GetAuthor).Methods("GET")
	r.HandleFunc("/authors", CreateAuthor).Methods("POST")
	r.HandleFunc("/authors/{id}", DeleteAuthor).Methods("DELETE")
	r.HandleFunc("/authors/{id}", UpdateAuthor).Methods("PUT")
	return r
}
