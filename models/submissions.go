package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/duguying/ojsite/judger"
	"github.com/gogather/com/log"
	"time"
)

type Submissions struct {
	Id       int
	Pid      int
	Uid      int
	Type     string
	Language string
	Code     string
	Judger   string

	//执行状态:'TA'(Task Add),'AC'(Accept),'WA'(Waiting),
	//'TLE'(Time Limit Error),'OLE'(Output Limit Error),
	//'MLE'(Memory Limit Error),'RE'(Runtime Error),
	//'PE'(Process Exit Normally),'CE'(Compile Error),'UK'(Unknown)
	Status        string
	BuildLog      string
	ExecuterDebug string
	SubmitTime    time.Time
	JudgeTime     time.Time
}

// add submission
func (this *Submissions) Add(pid int, uid int, ptype string, language string, code string, judgerName string) (int64, error) {

	o := orm.NewOrm()
	var subm Submissions
	subm.Pid = pid
	subm.Uid = uid

	if len(ptype) == 0 {
		ptype = "assert"
	}
	subm.Type = ptype
	subm.Language = language
	subm.Code = code
	subm.Judger = judgerName
	subm.Status = "TA"
	subm.SubmitTime = time.Now()
	subm.JudgeTime = time.Now()

	id, err := o.Insert(&subm)

	if err != nil {
		return id, err
	}

	jdg := judger.Get("default")
	jdg.AddTask(id, language, code)

	return id, err
}

// update submission status
func (this *Submissions) UpdateSubmissionStatus(id int, status string) error {
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

func (this *Submissions) GetSubmissionStatus(id int) (string, error) {

	jdg := judger.Get("default")
	response, err := jdg.GetStatus(int64(id))

	if err != nil {
		log.Warnln("send Request failed:", err)
	}

	info := response["info"].(map[string]interface{})

	o := orm.NewOrm()
	var subm Submissions
	subm.Id = id
	err = o.Read(&subm, "Id")

	if err != nil {
		log.Warnf("记录[%d]不存在\n", id)
		return "", err
	} else {

		status := info["run_result"].(string)

		if status == "PEN" {
			subm.Status = "AC"
		} else if status == "PRE" {
			subm.Status = "RE"
		} else if status == "POM" {
			subm.Status = "MLE"
		} else if status == "POT" {
			subm.Status = "TLE"
		} else if status == "POL" {
			subm.Status = "OLE"
		} else if status == "PSF" {
			subm.Status = "PSF"
		} else if status == "EC" {
			subm.Status = "CE"
		} else {
			subm.Status = "UK"
		}

		subm.BuildLog = info["build_log"].(string)
		exe_debug, ok := info["executer_debug"].(string)
		if ok {
			subm.ExecuterDebug = exe_debug
		}
		o.Update(&subm)

		return subm.Status, err
	}
}
