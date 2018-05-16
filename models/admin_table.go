package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

type Admin struct {
	Id 			int64
	Name		string
	Password	string
	CreatedTime	time.Time	`orm:"type(datetime);auto_now_add"`
	UpdatedTime	time.Time	`orm:"type(datetime);auto_now"`
}

func (this *Admin) InsertAdmin() error {
	fmt.Println(this)
	_, err := orm.NewOrm().Insert(this)
	return err
}