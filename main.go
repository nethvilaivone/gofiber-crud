package main

import (
	"github.com/gofiber/fiber/v2"
	
)

type Book struct {
	Id     int    `josn:"id"`
	Title  string `josn:"title"`
	Auther string `josn:"auther"`
}

var books []Book


func main() {

	books = append(books, Book{Id: 1, Title: "supider man", Auther: "neth"})
	books = append(books, Book{Id: 2, Title: "fast and furiuse", Auther: "ken"})

	app := fiber.New()

	app.Get("/books", getbooks)
	app.Get("/books/:id", getbookbyid)
	app.Post("/books", createBooks)
	app.Put("/books/:id", updateBooks)
    app.Delete("/books/:id", deletebook)
	app.Listen(":8080")
}
