package main

import (
	"gin/book/db"
	"gin/book/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// 初始化数据库链接
	err := db.InitDB()
	if err != nil {
		log.Println("连接数据库失败，error:", err)
		return
	}
	log.Println("连接数据库成功")
	r := gin.Default()
	// 加载html模板
	r.LoadHTMLGlob("views/*")
	// 设置静态资源目录
	r.Static("/static", "./static")
	// 定义路由
	// 书籍列表
	r.GET("/books", func(c *gin.Context) {
		var books []model.Book
		books, _ = model.GetBooks()
		c.HTML(http.StatusOK, "books.html", gin.H{
			"books": books,
		})
	})
	// 添加书籍页面
	r.GET("/books/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add.html", gin.H{})
	})
	// 编辑书籍页面
	r.GET("/books/:id/edit", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		book, _ := model.GetBook(id)
		c.HTML(http.StatusOK, "edit.html", gin.H{
			"book": book,
		})
	})
	// 添加或修改书籍
	r.POST("/books", func(c *gin.Context) {
		var book model.Book
		err := c.ShouldBind(&book)
		if err != nil {
			c.String(http.StatusOK, "修改书籍错误,error："+err.Error())
			return
		}
		book.CreatedAt = time.Now()
		book.UpdatedAt = time.Now()
		if book.Id > 0 {
			// 编辑
			model.UpdateBook(&book)
		} else {
			// 添加
			model.AddBook(&book)
		}
		c.Redirect(302, "/books")
	})
	// 执行删除操作
	r.POST("/books/:id/delete", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if id > 0 {
			model.DelBook(id)
		}
		c.Redirect(302, "/books")
	})
	r.Run(":8080")
}
