package main

import (
  "net/http"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "fmt"
  "strconv"

)

type Book struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
	Price float32 `json:"price"`
}


var mp = make(map[int]Book)
var counter = 1

func hello(c echo.Context) error{
	return c.String(http.StatusOK, "Hello World!")
}

func books(c echo.Context) error{
	// author := c.QueryParam("author")
	book := new(Book)
	// book.name = "Influence"
	// book.author = "Dr. Robert Cialdini"
	// book.price = 546.00

	if err := c.Bind(book); err!=nil{
		return err
	}

	return c.JSON(http.StatusOK, book)
}


func createBook(c echo.Context) error {
	book := new(Book)

	if err := c.Bind(book);err!=nil{
		return err
	}
	// mp[counter] = book
	mp[book.Id] = *book
	return c.JSON(http.StatusOK, book)

}

func getAllBooks(c echo.Context) error {

	fmt.Println("we are in get All books:", mp)
	return c.JSON(http.StatusOK, mp)
}

func bookDetails(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK,  mp[id])
}

// func getAllBooks(c echo.Context) error{

// 	book := &Book{
// 		Name: "Influence",
// 		Author: "Cialdini",
// 		Price: 356.00,
// 	}
// 	return c.JSON(http.StatusOK, book)
// }

func deleteBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	delete(mp, id)

	return c.NoContent(http.StatusNoContent)
}

func updateBook(c echo.Context) error {
	bk := new(Book)

	if err := c.Bind(bk); err!=nil{
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))

	mp[id] = *bk

	return c.JSON(http.StatusOK, mp[id])
}

func main(){

	// Echo instance
	e := echo.New()


	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Routes
	e.GET("/", hello)
	e.GET("/books", getAllBooks)
	 e.POST("/books", createBook)
	e.GET("/books/:id", bookDetails)
	e.DELETE("/books/:id", deleteBook)
	e.PUT("/books/:id", updateBook)

 	//Start server
 	e.Logger.Fatal(e.Start(":1234"))

}

