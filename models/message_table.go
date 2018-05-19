package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

/**
	报名信息表
	Status		0未读 1已读
 */

type Message struct {
	Id   		int64
	User		*User		`orm:"rel(fk)"`
	Activity 	*Activity	`orm:"rel(fk)"`
	Status		int			`orm:"default(0)"`
	CreatedTime	time.Time	`orm:"type(datetime);auto_now_add"`
	UpdatedTime	time.Time	`orm:"type(datetime);auto_now"`
}

func (this *Message) Insert() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Message) Update() error {
	_, err := orm.NewOrm().Update(this)
	return err
}