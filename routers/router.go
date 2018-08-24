package routers

import (
	"github.com/gin-gonic/gin"
	"../controllers"
)

func InitRouter() *gin.Engine{

	router := gin.Default()

	router.LoadHTMLGlob("views/*")

	//登录
	router.GET("/login",view.GetLogin)
	router.POST("/login/post",view.PostLogin)
	//注册
	router.GET("/register",view.GetRegister)
	router.POST("/register/post",view.PostRegister)

	//书籍列表
	router.GET("/book_list",view.GetBookList)
	router.POST("/book_add",view.AddBook)

	//修改书籍
	router.GET("/edit_book",view.EditOneBook)
	router.POST("/update_book",view.BookUpdate)

	//删除书籍
	router.GET("/delete_book",view.DeleteBook)

	return router

}