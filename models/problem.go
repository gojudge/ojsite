package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/gogather/com/log"
	"strconv"
	"time"
)

type Problem struct {
	Id          int
	Pbid        int
	Title       string
	Type        string
	Description string
	PreCode     string
	IoData      string
	Tags        string
	Level       string
	PassRate    float64
	Time        time.Time
}

// get problen by id or title
func (this *Problem) GetProblem(id int, title string) (Problem, error) {
	o := orm.NewOrm()
	var pro Problem
	var err error

	if id > 0 {
		pro.Id = id
		err = o.Read(&pro, "Id")
	} else if len(title) > 0 {
		pro.Title = title
		err = o.Read(&pro, "Title")
	} else {
		return pro, errors.New("at least one valid param")
	}

	return pro, err

}

// 获取problem列表
// page 页码
// itemsPerPage 每页数量
// level 题目权限级别
func (this *Problem) ListProblem(page int, itemsPerPage int, level string) (problems []orm.Params, hasNext bool, tatalPages int, err error) {
	sql1 := "select * from problem where level=? order by time desc limit ?,?"
	sql2 := "select count(*) as number from problem where level=?"

	var maps, maps2 []orm.Params
	var num int64
	var number int

	o := orm.NewOrm()

	if len(level) <= 0 {
		sql1 = "select * from problem order by time desc limit ?,?"
		sql2 = "select count(*) as number from problem"

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

		num, err = o.Raw(sql1, level, itemsPerPage*(page-1), itemsPerPage).Values(&maps)
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

// put problem into trash
func (this *Problem) TrashProblem(id int) error {
	o := orm.NewOrm()
	ob := orm.NewOrm()
	var pro Problem
	var prob ProblemBank
	var err error
	var num int64

	if id > 0 {
		// read from problem
		pro.Id = id
		err = o.Read(&pro, "Id")
		if err != nil {
			return err
		}

		// read from problem bank
		prob.Id = pro.Pbid
		err = ob.Read(&prob, "Id")
		if err == nil {
			prob.Status = "deleted"
			ob.Update(&prob)
		}

		num, err = o.Delete(&pro)
	} else {
		return errors.New("at least one valid param")
	}

	if num == 0 {
		return errors.New("no row to delete")
	}

	return err
}

// Add Problem into ProblemBank
func (this *Problem) AddProblem(title string, ptype string, description string, precode string, iodata string, tags string) (id int64, err error) {
	o := orm.NewOrm()
	var prob ProblemBank

	prob.Title = title
	prob.Type = ptype
	prob.Description = description
	prob.PreCode = precode
	prob.IoData = iodata
	prob.Tags = tags
	prob.Status = "audit"

	id, err = o.Insert(&prob)
	return id, err
}
