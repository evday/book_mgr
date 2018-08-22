package data_layer

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	Db *sqlx.DB
)

func init(){
	var err error
	dns := "root:@tcp(localhost:3306)/golang?parseTime=True"
	Db,err = sqlx.Connect("mysql",dns)

	if err != nil {
		panic(err)
	}

	err = Db.Ping()
	if err != nil {
		panic(err)	
	}
}