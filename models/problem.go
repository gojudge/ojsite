package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/gogather/com/log"
	"strconv"
)

type Problem struct {
	Id          int
	Title       string
	Type        string
	Description string
	PreCode     string
	IoData      string
	Tags        string
	Level       string
}

// 获取problem列表
// page 页码
// itemsPerPage 每页数量
// level 题目权限级别
func (this *Problem) ListProblem(page int, itemsPerPage int, level string) (problems []orm.Params, hasNext bool, tatalPages int, err error) {
	sql1 := "select * from problem where level=? order by time desc limit ?,?"
	sql2 := "select count(*) as number from problem where level=?"
	var maps, maps2 []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(sql1, level, itemsPerPage*(page-1), itemsPerPage).Values(&maps)
	if err != nil {
		log.Warnln("execute sql1 error:")
		log.Warnln(err)
		return nil, false, 0, err
	}

	_, err = o.Raw(sql2, level).Values(&maps2)
	if err != nil {
		log.Warnln("execute sql2 error:")
		log.Warnln(err)
		return nil, false, 0, err
	}

	number, err := strconv.Atoi(maps2[0]["number"].(string))

	var addFlag int
	if 0 == (number % itemsPerPage) {
		addFlag = 0
	} else {
		addFlag = 1
	}

	tatalPages = number/itemsPerPage + addFlag

	if tatalPages == page {
		hasNext = false
	} else {
		hasNext = true
	}

	if err == nil && num > 0 {
		return maps, hasNext, tatalPages, nil
	} else {
		return nil, false, tatalPages, err
	}
}

// get top 10 problem
func (this *Problem) GetTop10() ([]orm.Params, error) {
	sql := `SELECT problem.id as id, problem.title as title, count(*) AS count FROM submissions,problem where submissions.pid=problem.id GROUP BY pid ORDER BY count DESC limit 10`

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
