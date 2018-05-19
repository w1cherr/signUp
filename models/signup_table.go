package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

/**
	报名信息表
	Status		0待审核 1审核通过
 */

type SignUp struct {
	Id   		int64
	User		*User		`orm:"rel(fk)"`
	Activity 	*Activity	`orm:"rel(fk)"`
	Status		int			`orm:"default(0)"`
	CreatedTime	time.Time	`orm:"type(datetime);auto_now_add"`
	UpdatedTime	time.Time	`orm:"type(datetime);auto_now"`
}

func (this *SignUp) Insert() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *SignUp) Update() error {
	_, err := orm.NewOrm().Update(this)
	return err
}