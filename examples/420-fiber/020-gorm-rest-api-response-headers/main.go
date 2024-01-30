// curl http://localhost:3000/api/v1/book/1
// curl http://localhost:3000/api/v1/book
// curl -X POST http://localhost:3000/api/v1/book
// curl -X DELETE http://localhost:3000/api/v1/book
// curl -X POST  -d '{"ID":1,"title":"1984","author":"George Orwell","rating":5}' http://localhost:3000/api/v1/book
// curl -X DELETE http://localhost:3000/api/v1/book/1
// https://www.youtube.com/watch?v=Iq2qT0fRhAA
package main

import (
	"fiber/book"
	"fiber/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"gorm.io/gorm"
	"net/http"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", api)
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	// You can also use file::memory:?cache=shared instead of a path to a file.
	// This will tell SQLite to use a temporary database in system memory.
	database.DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to a database")
	}
	fmt.Println("Database connection sucessfully stablished")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func api(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("Book Api")
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)

	app.Listen(":3000")

}
