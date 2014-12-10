package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/gogather/com/log"
	"time"
)

type Tags struct {
	Id           int
	Tag          string
	ProblemNum   int
	RegistorTime time.Time
}

// Add Tag
func (this *Tags) AddTag(tag string) (int64, error) {
	o := orm.NewOrm()
	var t Tags
	t.Tag = tag

	return o.Insert(&t)
}

// List Tags
func (this *Tags) TagList() ([]orm.Params, error) {
	sql := "select * from tags order by time desc"

	var maps []orm.Params
	o := orm.NewOrm()
	_, err := o.Raw(sql).Values(&maps)
	if err != nil {
		log.Warnln("execute sql error:")
		log.Warnln(err)
		return nil, err
	} else {
		return maps, err
	}

}
