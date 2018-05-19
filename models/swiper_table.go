package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

type Swiper struct {
	Id 			int64
	Cover		string
	ActivityId	int64
	CreatedTime	time.Time	`orm:"type(datetime);auto_now_add"`
	UpdatedTime	time.Time	`orm:"type(datetime);auto_now"`
}

func (this *Swiper) Insert() error {
	fmt.Println(this)
	_, err := orm.NewOrm().Insert(this)
	return err
}

func (this *Swiper) Update() error {
	_, err := orm.NewOrm().Update(this)
	return err
}
func (this *Swiper) Delete() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Swiper) Read() error {
	return orm.NewOrm().QueryTable("swiper").Filter("id", this.Id).RelatedSel().One(this)
}
