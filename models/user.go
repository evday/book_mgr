package models


import (
	"fmt"
	"encoding/hex"
	"crypto/md5"
	"database/sql"
	db "../database"
)


type User struct {
	Id int64 `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

func (user *User)Md5()(md5Str string){
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(user.Password))
	cipherStr := md5Ctx.Sum(nil)
	md5Str = hex.EncodeToString(cipherStr)
	return

}


func (user *User)InsertUser() (err error) {
	sqlstr := `select username from user where username=?`
	var Username string
	err = db.Db.Get(&Username,sqlstr,user.Username)
	if err == sql.ErrNoRows{
		password := user.Md5()
		sqlstr = `insert into user(username,password)values(?,?)`
		_,err = db.Db.Exec(sqlstr,user.Username,password)
		if err != nil{
			err = fmt.Errorf("保存用户失败,err:%v\n",err)
			return
		}
		return 
	}else if err != nil{
		err = fmt.Errorf("查询失败,请检查你的查询语句是否正确,err:%v\n",err)
		return
	}else{
		err = fmt.Errorf("数据已存在，不能重复插入,err:%v\n",err)
		return 
	}
}