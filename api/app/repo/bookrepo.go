package repositories

import (
	model "HW4/api/app/model"
	"HW4/api/config"
)

func GetAllBooks() []model.Book {
	var books []model.Book
	config.DB.Preload("Author").Find(&books)
	return books
}
func GetBook(id int) model.Book {
	var book model.Book
	config.DB.First(&book, id)
	return book
}
func CreateBook(book model.Book) model.Book {
	config.DB.Create(&book)
	return book
}
func UpdateBook(book model.Book) model.Book {
	config.DB.Save(&book)
	return book
}
func DeleteBook(id int) model.Book {
	var book model.Book
	config.DB.Where("id=?", id).Delete(&book)
	return book
}
func GetSearchBook(bookName string) model.Book {
	var book model.Book

	config.DB.Where("names ILIKE ? ", "%"+bookName+"%").Find(&book)

	return book
}
func BuyBookByID(id int, number int) (model.Book, error) {
	var book model.Book
	err := config.DB.Where(model.Book{BookID: id}).Find(&book)
	if err == nil {
		log.Fatal("Stock Not Found: ", err)
	} else {
		book.StockNumber = book.StockNumber - int(number)
		result := config.DB.Save(book)
		if result.Error != nil {
			return book, result.Error
		}
	}
	return book, nil
}
