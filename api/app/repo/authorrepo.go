package repositories

import (
	model "HW4/api/app/model"
	"HW4/api/config"
)

func GetAllAuthors() []model.Author {
	var authors []model.Author
	config.DB.Preload("Book").Find(&authors)
	return authors
}
func GetAuthor(id int) model.Author {
	var author model.Author
	config.DB.First(&author, id)
	return author
}
func CreateAuthor(author model.Author) model.Author {
	config.DB.Create(&author)
	return author
}
func UpdateAuthor(author model.Author) model.Author {
	config.DB.Save(&author)
	return author
}
func DeleteAuthor(id int) model.Author {
	var author model.Author
	config.DB.Where("id=?", id).Delete(&author)
	return author
}
