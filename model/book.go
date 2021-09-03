package model

import (
	"gin/book/db"
	"time"
)

// 定义结构体

// Book 定义模型结构体
type Book struct {
	Id        int       `form:"id"`
	Name      string    `form:"name"`
	Price     float64   `form:"price"`
	CreatedAt time.Time `db:"created_at" form:"-"`
	UpdatedAt time.Time `db:"updated_at" form:"-"`
}

// 定义模型方法

// GetBooks 获取所有书籍
func GetBooks() (books []Book, err error) {
	query := `SELECT * FROM books`
	err = db.Db.Select(&books, query)
	return
}

// GetBook 根据ID获取书籍信息
func GetBook(id int) (book Book, err error) {
	query := `SELECT * FROM books WHERE id = ?`
	err = db.Db.Get(&book, query, id)
	if err != nil {
		return
	}
	return
}

// UpdateBook 修改书籍信息
func UpdateBook(book *Book) (err error) {
	query := `UPDATE books SET name=?,price=?,updated_at=? WHERE id = ?`
	_, err = db.Db.Exec(query, book.Name, book.Price, book.UpdatedAt.Format("20060102150405"), book.Id)
	return
}

// AddBook 添加书籍
func AddBook(book *Book) (err error) {
	query := `INSERT INTO books(name,price,created_at,updated_at) VALUES(?,?,?,?)`
	exec, err := db.Db.Exec(query, book.Name, book.Price, book.CreatedAt.Format("20060102150405"), book.UpdatedAt.Format("20060102150405"))
	if err != nil {
		return
	}
	id, err := exec.LastInsertId()
	if err != nil {
		return
	}
	book.Id = int(id)
	return
}

// DelBook 删除书籍
func DelBook(id int) (err error) {
	query := `DELETE FROM books WHERE id = ?`
	_, err = db.Db.Exec(query, id)
	return
}
