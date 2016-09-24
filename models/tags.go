package models

import (
	"github.com/gogather/com/log"
	"time"
)

type Tags struct {
	Id           int
	Tag          string `xorm:"varchar(128)"`
	ProblemNum   int
	RegistorTime time.Time
}

// Add Tag
func (this *Tags) AddTag(tag string) (int64, error) {
	var t Tags
	t.Tag = tag

	return engine.Insert(&t)
}

// List Tags
func (this *Tags) TagList() ([]map[string][]byte, error) {
	sql := "select * from tags order by time desc"

	res, err := engine.Query(sql)
	if err != nil {
		log.Warnln("execute sql error:")
		log.Warnln(err)
		return nil, err
	} else {
		log.Warnln("tags", res)
		return res, err
	}

}
