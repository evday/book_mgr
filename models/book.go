package models

import (
	"fmt"
	"database/sql"
	"time"
	db "../database"
)

type Book struct{
	Id int64 `db:"id"`
	BookName string `db:"bookname"`
    Price float64 `db:"price"`
    Author string `db:"author"`
	PublishTime time.Time `db:"publishtime"`
	StockNum uint  `db:"stocknum"`
	BookId string 	`db:"bookid"`
}

func (book *Book)InsertBook()(err error){
	sqlstr := `select bookid from book where bookid = ?`
	var BookId string
	err = db.Db.Get(&BookId,sqlstr,book.BookId)
	if err == sql.ErrNoRows{
		//插入
		err = fmt.Errorf("记录不存在:%v\n",err)
		return
	}else if err != nil{
		err = fmt.Errorf("查询失败:%v\n",err)
		return 
	}else{
		//更新
		err = fmt.Errorf("更新记录：%v\n",err)
		return
	}
	

}