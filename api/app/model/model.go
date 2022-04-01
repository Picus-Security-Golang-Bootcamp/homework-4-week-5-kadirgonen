package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

type Books struct {
	Books []Book `json:books`
}
type Book struct {
	BookID      int     `json:"book_id" gorm:"primaryKey"`
	Names       string  `json:"names"`
	PaperNumber int     `json:"paper_number"`
	StockNumber int     `json:"stock_number"`
	Cost        float64 `json:"cost"`
	StockCode   int     `json:"stock_code"`
	Isbn        int     `json:"ISBN"`
	AuthorID    int
	Author      Author `json:"author" gorm:"foreignKey:AuthorID;references:ID"`
	Deleted     gorm.DeletedAt
}

type Author struct {
	ID   int32  `json:"id"  gorm:"primary_key"`
	Name string `json:"name"`
}
// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Book{})
	return db
}
func (b *Book) ToString() string {

	return fmt.Sprintf("BookID : %d, Names : %s, PaperNumber : %d,StockNumber : %d, Cost : %.f, StockCode : %d, ISBN : %d, Author : %s", b.BookID, b.Names, b.PaperNumber, b.StockNumber, b.Cost, b.StockCode, b.Isbn, b.Author.Name)
}

func GetAllBooksFromJson() []Book {
	list := []Book{}
	jsonFile, err := os.Open("books.json")
	if err != nil {
		log.Fatal("opening json: ", err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Println(byteValue)
	err = json.Unmarshal(byteValue, &list)
	if err != nil {
		log.Fatal(" unmarshal json: ", err)
	}
	return list
}
func InsertSampleData(db *gorm.DB) {
	jsonFile, err := os.Open("books.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	values, _ := ioutil.ReadAll(jsonFile)
	books := []Book{}
	json.Unmarshal(values, &books)

	for _, book := range books {
		db.FirstOrCreate(&book)
	}
}
