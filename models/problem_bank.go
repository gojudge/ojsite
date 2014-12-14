package models

import (
	// "errors"
	"github.com/astaxie/beego/orm"
	"github.com/gogather/com/log"
	"strconv"
)

type ProblemBank struct {
	Id          int
	Title       string
	Type        string
	Description string
	PreCode     string
	IoData      string
	Tags        string
	Status      string
}

// 获取problem列表
// page 页码
// itemsPerPage 每页数量
// status 题目权限级别
func (this *ProblemBank) ListProblem(page int, itemsPerPage int, status string) (problems []orm.Params, hasNext bool, tatalPages int, err error) {
	sql1 := "select * from problem_bank where status=? order by time desc limit ?,?"
	sql2 := "select count(*) as number from problem_bank where status=?"

	var maps, maps2 []orm.Params
	var num int64
	var number int

	o := orm.NewOrm()

	if len(status) <= 0 {
		sql1 = "select * from problem_bank order by time desc limit ?,?"
		sql2 = "select count(*) as number from problem_bank"

		num, err = o.Raw(sql1, itemsPerPage*(page-1), itemsPerPage).Values(&maps)
		if err != nil {
			log.Warnln("execute sql1 error:")
			log.Warnln(err)
			return nil, false, 0, err
		}

		_, err = o.Raw(sql2).Values(&maps2)
		if err != nil {
			log.Warnln("execute sql2 error:")
			log.Warnln(err)
			return nil, false, 0, err
		}

	} else {

		num, err = o.Raw(sql1, status, itemsPerPage*(page-1), itemsPerPage).Values(&maps)
		if err != nil {
			log.Warnln("execute sql1 error:")
			log.Warnln(err)
			return nil, false, 0, err
		}

		_, err = o.Raw(sql2, status).Values(&maps2)
		if err != nil {
			log.Warnln("execute sql2 error:")
			log.Warnln(err)
			return nil, false, 0, err
		}

	}

	number, err = strconv.Atoi(maps2[0]["number"].(string))

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
