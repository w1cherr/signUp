package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

type User struct {
	Id 			int64
	Name		string
	Password	string
	CanSignUp 	int8
	CreatedTime	time.Time	`orm:"type(datetime);auto_now_add"`
	UpdatedTime	time.Time	`orm:"type(datetime);auto_now"`
}

func (this *User) InsertUser() error {
	fmt.Println(this)
	_, err := orm.NewOrm().Insert(this)
	return err
}