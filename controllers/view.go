package view

import (

	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"../models"
)

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK,"login.html",gin.H{})
}

func PostLogin(c *gin.Context){
	c.String(http.StatusOK,"提交信息收到")
}

func GetRegister(c *gin.Context){
	c.HTML(http.StatusOK,"register.html",gin.H{})
}

func PostRegister(c *gin.Context){

	username := c.PostForm("username")
	password := c.PostForm("password")
	re_password:= c.PostForm("re_password")

	if password != re_password{
		c.String(http.StatusOK,"两次密码不一致")
		return
	}
	var user models.User
	user.Username = username
	user.Password = password
	
	err := user.InsertUser()
	fmt.Println(err)	
}


func GetBookList(c *gin.Context) {
	var book models.Book

	books,err := book.GetAllBook()
	if err != nil{
		fmt.Println(err)
	}

	var book_list = make([]interface{},0)
	for _,book = range books{
		var book_map = make(map[string]string)
		book_map["book_name"] = book.BookName
		book_map["book_id"] = book.BookId
		book_map["book_price"] = book.Price
		book_map["book_author"] = book.Author
		book_map["book_stocknum"] = book.StockNum
		book_map["book_publishtime"] = book.PublishTime

		book_list = append(book_list,book_map)
	}
	c.HTML(http.StatusOK,"book_list.html",book_list)
}

func AddBook(c *gin.Context) {
	book_name := c.PostForm("book_name")
	price := c.PostForm("price")
	author := c.PostForm("author")
	publish_time := c.PostForm("publish_time")
	add_num := c.PostForm("add_num")
	book_id := c.PostForm("book_id")

	
	var book models.Book

	book.BookId = book_id
	book.BookName = book_name
	book.Price = price
	book.Author = author
	book.PublishTime = publish_time
	book.StockNum = add_num


	err := book.InsertBook()
	if err != nil{
		fmt.Println(err)
		return
	}
	c.Redirect(302,"/book_list")
}

func EditOneBook(c *gin.Context){
	book_id := c.Query("book_id")
	fmt.Println(book_id)
	var book models.Book
	book.BookId = book_id
	bookinfo,err := book.GetOneBook()
	if err != nil{
		fmt.Println(err)
	}
	c.JSON(200,bookinfo)
}

func BookUpdate(c *gin.Context) {
	book_name := c.PostForm("book_name")
	price := c.PostForm("price")
	author := c.PostForm("author")
	publish_time := c.PostForm("publish_time")
	add_num := c.PostForm("add_num")
	book_id := c.PostForm("book_id")

	fmt.Println(book_id)

	
	var book models.Book

	book.BookId = book_id
	book.BookName = book_name
	book.Price = price
	book.Author = author
	book.PublishTime = publish_time
	book.StockNum = add_num


	err := book.UpdateBook()
	if err != nil{
		fmt.Println(err)
		return
	}
	c.Redirect(302,"/book_list")
}


func DeleteBook(c *gin.Context){
	book_id := c.Query("book_id")
	var book models.Book

	book.BookId = book_id

	err := book.DeleteOneBook()

	if err != nil{
		fmt.Println(err)
		return
	}
	c.Redirect(302,"/book_list")

	
}