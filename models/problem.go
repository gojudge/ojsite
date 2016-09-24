package models

import (
	"errors"
	"github.com/gogather/com/log"
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
	var pro Problem
	var err error

	if id > 0 {
		pro.Id = id
		_, err = engine.Get(&pro)
	} else if len(title) > 0 {
		pro.Title = title
		_, err = engine.Get(&pro)
	} else {
		return pro, errors.New("at least one valid param")
	}

	return pro, err

}

// 获取problem列表
// page 页码
// itemsPerPage 每页数量
// level 题目权限级别
func (this *Problem) ListProblem(page int, itemsPerPage int, level string) (problems []*Problem, hasNext bool, tatalPages int64, err error) {
	var num int64
	var p Problem

	if len(level) <= 0 {
		err := engine.Desc("time").Limit(itemsPerPage, itemsPerPage*(page-1)).Find(&problems)
		if err != nil {
			log.Warnln("execute sql1 error:")
			log.Warnln(err)
			return nil, false, 0, err
		}

		num, err = engine.Count(&p)
		if err != nil {
			log.Warnln("execute sql2 error:")
			log.Warnln(err)
			return nil, false, 0, err
		}

	} else {

		err = engine.Where("level = ?", level).Limit(itemsPerPage, itemsPerPage*(page-1)).Find(&problems)
		if err != nil {
			log.Warnln("execute sql1 error:")
			log.Warnln(err)
			return nil, false, 0, err
		}

		num, err = engine.Where("level = ?", level).Count(&p)
		if err != nil {
			log.Warnln("execute sql2 error:")
			log.Warnln(err)
			return nil, false, 0, err
		}

	}

	log.Warnln("test engine")

	var addFlag int64
	if 0 == (num % int64(itemsPerPage)) {
		addFlag = 0
	} else {
		addFlag = 1
	}

	tatalPages = num/int64(itemsPerPage) + addFlag

	if tatalPages == int64(page) {
		hasNext = false
	} else {
		hasNext = true
	}

	if err == nil && num > 0 {
		return problems, hasNext, tatalPages, nil
	} else {
		return nil, false, tatalPages, err
	}
}

// get top 10 problem
func (this *Problem) GetTop10() ([]map[string][]byte, error) {
	sql := `SELECT problem.id as id, problem.title as title, count(*) AS count FROM submissions,problem where submissions.pid=problem.id GROUP BY pid ORDER BY count DESC limit 10`

	maps, err := engine.Query(sql)
	if err != nil {
		log.Warnln("execute sql error:")
		log.Warnln(err)
		return nil, err
	} else {
		return maps, err
	}
}

/*
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
*/
