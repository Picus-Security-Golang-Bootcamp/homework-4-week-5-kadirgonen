## Homework 4 | Week 5
In this application, we aim to make a RESTful API by making a book application and providing a few status checks of the books in the application with the help of database.

## Purpose of application

We have a list of books.
The book areas are as follows;
```
- Book ID
- Book Name
- Paper Number
- Stock Number
- Cost
- Stock Code
- ISBN
- Author (ID and Name)
```
1. Book and Author information read from a file.
2. It record in the database.
3. Database queries writte.(GORM Queries)
4. GORM Queries writte.
5. Creating a handle functions.
6. As a result also complete RESTful API.

 
## Used Handle Functions

### for book
```
	HandleFunc("/books", GetAllBooks).Methods("GET")
	HandleFunc("/books", CreateBook).Methods("POST")
	HandleFunc("/books/{id}", GetBook).Methods("GET")
  HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")
```
These functions allow us to create a RESTful API with the books we define in the application.

### for author
```
	HandleFunc("/authors", GetAllAuthors).Methods("GET")
	HandleFunc("/authors/{id}", GetAuthor).Methods("GET")
	HandleFunc("/authors", CreateAuthor).Methods("POST")
	HandleFunc("/authors/{id}", DeleteAuthor).Methods("DELETE")
	HandleFunc("/authors/{id}", UpdateAuthor).Methods("PUT")
```
These functions allow us to create a RESTful API with the authors we define in the application.

## Requirements

* Go Language
* Git
* Go Module
* GORM
* Gorilla/Mux

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc
