package models

import (
	// "errors"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/duguying/ojsite/utils"
	// "github.com/gogather/com"
	"github.com/gogather/com/log"
	"time"
)

type Submissions struct {
	Id         int
	Pid        int
	Uid        int
	Type       string
	Language   string
	Code       string
	Judger     string
	Status     string
	BuildLog   string
	RunLog     string
	SubmitTime time.Time
	JudgeTime  time.Time
}

// add submission
func (this *Submissions) Add(pid int, uid int, ptype string, language string, code string, judger string) (int64, error) {
	o := orm.NewOrm()
	var subm Submissions
	subm.Pid = pid
	subm.Uid = uid
	subm.Type = ptype
	subm.Language = language
	subm.Code = code
	subm.Judger = judger

	return o.Insert(&subm)
}

// update submission status
func UpdateSubmissionStatus(id int, status string) error {
	o := orm.NewOrm()
	var subm Submissions
	subm.Id = id
	err := o.Read(&subm, "Id")

	if err != nil {
		log.Warnf("记录[%d]不存在\n", id)
		return err
	} else {
		subm.Status = status
		if _, err := o.Update(&subm); err != nil {
			log.Warnln("状态更新失败")
			return err
		} else {
			return nil
		}
	}
}
