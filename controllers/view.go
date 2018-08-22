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
	c.HTML(http.StatusOK,"book_list.html",gin.H{})
}

func AddBook(c *gin.Context) {
	book_name := c.PostForm("book_id")
	price := c.PostForm("price")
	author := c.PostForm("author")
	publish_time := c.PostForm("publish_time")
	add_num := c.PostForm("add_num")
	book_id := c.PostForm("book_name")
	
	fmt.Println(book_id)
	fmt.Println(book_name)
	fmt.Println(price)
	fmt.Println(publish_time)
	fmt.Println(add_num)
	fmt.Println(author)

	var book models.Book

	book.BookId = book_id
	err := book.InsertBook()
	fmt.Println(err)
}
