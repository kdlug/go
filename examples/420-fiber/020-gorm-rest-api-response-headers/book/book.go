package book

import (
	"fiber/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn

	var books []Book
	db.Find(&books)

	return c.Status(http.StatusOK).JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.Find(&book, id)

	return c.Status(http.StatusOK).JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	db.Create(&book)

	return c.Status(http.StatusCreated).JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		return c.Status(http.StatusInternalServerError).SendString("No book found")
	}

	db.Delete(&book)
	return c.Status(http.StatusOK).SendString("Book successfully deleted")
}
