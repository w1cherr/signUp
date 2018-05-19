package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"time"
)

/**
	活动表
	ImgUrl 		逗号拼接图片url成字符串
	UserCount	已报名人数
	UserLimit 	限制报名人数
	Status		0不可报名 1可报名
 */

type Activity struct {
	Id 				int64
	Title			string
	Introduction	string
	Cover			string
	ImgUrl			string
	UserCount		int64
	UserLimit 		int64
	Status			int			`orm:"default(1)"`
	StartTime 		string
	EndTime 		string
	CreatedTime		time.Time	`orm:"type(datetime);auto_now_add"`
	UpdatedTime		time.Time	`orm:"type(datetime);auto_now"`
}

func (this *Activity) Insert() error {
	_, err := orm.NewOrm().Insert(this)
	return err
}

func (this *Activity) Update() error {
	_, err := orm.NewOrm().Update(this)
	return err
}

func (this *Activity) Delete() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Activity) Read() error {
	return orm.NewOrm().QueryTable("activity").Filter("id", this.Id).RelatedSel().One(this)
}

