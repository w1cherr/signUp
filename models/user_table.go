package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

/**
	用户表
	CanSignUp 0不可报名 1可报名
 */

type User struct {
	Id 			int64
	Name		string
	Password	string
	CanSignUp 	int			`orm:"default(1)"`
	CreatedTime	time.Time	`orm:"type(datetime);auto_now_add"`
	UpdatedTime	time.Time	`orm:"type(datetime);auto_now"`
}

func (this *User) InsertUser() error {
	fmt.Println(this)
	_, err := orm.NewOrm().Insert(this)
	return err
}

func (this *User) Update() error {
	_, err := orm.NewOrm().Update(this)
	return err
}

func (this *User) Read() error {
	return orm.NewOrm().QueryTable("User").Filter("id", this.Id).RelatedSel().One(this)
}