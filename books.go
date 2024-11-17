package main 


import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)


func getbooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getbookbyid(c *fiber.Ctx) error {
	bookid, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, book := range books {
		if book.Id == bookid {
			return c.JSON(book)
		}

	}

	return c.Status(fiber.StatusNotFound).SendString("not found book id")
}

func createBooks(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("could'nt pars string")
	}
	books = append(books, *book)

	return c.JSON(book)

}

func updateBooks(c *fiber.Ctx) error {
	bookid, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	updateNewbook := new(Book)
	if err := c.BodyParser(updateNewbook); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	
	for i, book := range books {
		if book.Id == bookid {
			books[i].Title = updateNewbook.Title
			books[i].Auther = updateNewbook.Auther
			return c.JSON(books[i])
		}

	}

	return c.Status(fiber.StatusNotFound).SendString(err.Error())
}


func deletebook(c *fiber.Ctx) error {
	bookid, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	} 

	for i, book := range books {
		if book.Id == bookid {
			books = append(books[:i],books[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}


	}
	return c.Status(fiber.StatusNotFound).SendString(err.Error())

}