package main

import (
	"log"
	"os"
	
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	jwtware "github.com/gofiber/jwt/v2"
	
)

type Book struct {
	Id     int    `josn:"id"`
	Title  string `josn:"title"`
	Auther string `josn:"auther"`
}
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var userMember = User{Email: "neth@gmail.com", Password: "123123"}
var books []Book

func main() {

	books = append(books, Book{Id: 1, Title: "supider man", Auther: "neth"})
	books = append(books, Book{Id: 2, Title: "fast and furiuse", Auther: "ken"})

	if err := godotenv.Load(); err != nil {
		log.Fatal("env not found")
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Post("/login", login)

	// app.Use(checkMideeleWare)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is not set")
	}
	
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	}))

	app.Use(checkMideeleWare)

	app.Get("/books", getbooks)
	app.Get("/books/:id", getbookbyid)
	app.Post("/books", createBooks)
	app.Put("/books/:id", updateBooks)
	app.Delete("/books/:id", deletebook)
	app.Post("/upload", updoadfile)
	app.Get("/testHTML", gettemplathtmlfunc)
	app.Get("/config", getEVN)
	app.Listen(":8080")
}
