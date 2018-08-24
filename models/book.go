package models

import (
	"fmt"
	"database/sql"
	db "../database"


)

type Book struct{
	Id int64 `json:"id" form:"id"`
	BookName string `json:"bookname" form:"bookname"`
    Price string `json:"price" form:"price"`
    Author string `json:"author" form:"author"`
	PublishTime string `json:"publishtime" form:"publishtime"`
	StockNum string  `json:"stocknum" form:"stocknum"`
	BookId string 	`json:"bookid" form:"bookid"`
}

func(book *Book)GetAllBook()(books []Book,err error){
	books = make([]Book,0)
	sqlstr := `select bookid,bookname,bookprice,stocknum,publishtime,author from book`
	rows,err := db.Db.Query(sqlstr)
	defer rows.Close()
	if err != nil{
		err = fmt.Errorf("获取书籍列表失败，err:%v\n",err)
		return
	}
	for rows.Next(){
		var book Book
		rows.Scan(&book.BookId,&book.BookName,&book.Price,&book.StockNum,&book.PublishTime,&book.Author)
		
		books = append(books,book)	
	}
	if err = rows.Err(); err != nil {
		return 
	}
	return

}

func (book *Book)GetOneBook()(bookinfo map[string]string,err error){
	
	sqlstr := `select bookid,bookname,bookprice,stocknum,publishtime,author from book where bookid=?`
	row := db.Db.QueryRow(sqlstr,book.BookId)
	err = row.Scan(&book.BookId,&book.BookName,&book.Price,&book.StockNum,&book.PublishTime,&book.Author)
	if err != nil{
		err = fmt.Errorf("没有找到您要的书籍,err:%v\n",err)
		return

	}
	bookinfo = make(map[string]string)
	bookinfo["book_name"] = book.BookName
	bookinfo["book_id"] = book.BookId
	bookinfo["book_price"] = book.Price
	bookinfo["book_author"] = book.Author
	bookinfo["book_stocknum"] = book.StockNum
	bookinfo["book_publishtime"] = book.PublishTime

	return 
}

func(book *Book)InsertBook()(err error){
	sqlstr := `select bookid from book where bookid = ?`
	var BookId string
	err = db.Db.Get(&BookId,sqlstr,book.BookId)
	if err == sql.ErrNoRows{
		//插入
		sqlx := `insert into book(bookid,bookname,bookprice,publishtime,stocknum,author)
		values(?,?,?,?,?,?)`
		_, err = db.Db.Exec(sqlx, book.BookId, book.BookName, book.Price, book.PublishTime, book.StockNum, book.Author)

		if err != nil {
			err = fmt.Errorf("添加书籍失败,err:%v\n", err)
			return
		}
		return
	}else if err != nil{
		err = fmt.Errorf("查询失败:%v\n",err)
		return 
	}else{
		//更新
		err = fmt.Errorf("记录已存在不能重复插入:%v\n",err)
		return err
	}
	

}

func (book *Book)UpdateBook()(err error){
	sqlstr := "UPDATE book SET bookname = ?,bookprice = ?,author = ?,publishtime = ?,stocknum = ? WHERE bookid= ? "
	result,err := db.Db.Exec(sqlstr,book.BookName,book.Price,book.Author,book.PublishTime,book.StockNum,book.BookId)
	
	if err != nil{
		err = fmt.Errorf("更新书籍失败,err:%v\n", err)
		return err
	}
	affects, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affects == 0 {
		err = fmt.Errorf("update book failed, bookid:%s, not found", book.BookId)
		return err
	}
	return
}

func(book *Book)DeleteOneBook()(err error){
	sqlstr := `delete from book where bookid=?`
	result,err := db.Db.Exec(sqlstr,book.BookId)
	if err != nil{
		err = fmt.Errorf("删除书籍失败,err:%v\n", err)
		return err
	}
	affects, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affects == 0 {
		err = fmt.Errorf("delete book failed, bookid:%s, not found", book.BookId)
		return err
	}
	return
}
