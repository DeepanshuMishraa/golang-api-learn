package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
"errors"
)

type book struct {
	ID   string  `json:"id"`
	Title string  `json:"title"`
	Author string  `json:"author"`
	Quantity int   `json:"quantity"`
}


var books = []book{
	{ID: "1", Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Quantity: 3},
	{ID: "2", Title: "The Hobbit", Author: "J.R.R. Tolkien", Quantity: 7},
	{ID: "3", Title: "The Silmarillion", Author: "J.R.R. Tolkien", Quantity: 5},
}

func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books);
}

func getBookById(id string)(*book,error){
	for i, b := range books{
		if b.ID == id{
			return &books[i],nil;
		}
	}

	return nil,errors.New("book not found");
}

func bookById(c *gin.Context){
	id:= c.Param("id");
	b,err:= getBookById(id);

	if err!=nil{
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book not found"});
		return ;
	}

	c.IndentedJSON(http.StatusOK,b);
}

func checkOutBook(c *gin.Context){
	id,ok := c.GetQuery("id");

	if ok == false{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"id is required"});
		return ;
	}

	book,err := getBookById(id);

	if err!=nil{
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book not found"});
		return ;
	}

	if book.Quantity == 0{
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book not available"});
		return ;
	}

	book.Quantity--;

	c.IndentedJSON(http.StatusOK,book);
}

func createBook(c *gin.Context){
	var newBook book;

	if err:= c.BindJSON(&newBook);err!=nil{
		return ;
	}

	books = append(books,newBook);
	c.IndentedJSON(http.StatusCreated,newBook);
}

func main(){
	router:= gin.Default();
	router.GET("/books",getBooks);
	router.POST("/create",createBook);
	router.GET("/books/:id",bookById);
	router.PATCH("/checkout",checkOutBook);
	router.Run("localhost:8080");
}
