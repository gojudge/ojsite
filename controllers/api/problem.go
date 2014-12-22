package api

import (
	// "github.com/astaxie/beego"
	"github.com/duguying/ojsite/controllers"
	"github.com/duguying/ojsite/models"
	"github.com/duguying/ojsite/utils"
	"github.com/gogather/com/log"
	"strconv"
)

// get problem list
type ProblemListController struct {
	controllers.BaseController
}

// uri /api/problem/list/:page
func (this *ProblemListController) Get() {
	page, err := this.GetInt("page")
	if nil != err || page < 0 {
		page = 0
	}

	s := this.Ctx.Input.Param(":page")
	pageParm, err := strconv.Atoi(s)
	if nil != err || pageParm < 0 {
		pageParm = 0
	} else {
		page = pageParm
	}

	if 0 == page {
		page = 1
	}

	status := this.GetString("status")
	if len(status) <= 0 {
		status = "ok"
	}

	if status == "ok" {
		pro := models.Problem{}
		data, hasNext, totalPage, err := pro.ListProblem(page, 20, "")

		if err != nil {
			this.Data["json"] = map[string]interface{}{
				"result":     false,
				"msg":        "get list failing",
				"list":       nil,
				"has_next":   false,
				"total_page": 0,
				"refer":      nil,
			}
		} else {
			log.Blueln(data)

			this.Data["json"] = map[string]interface{}{
				"result":     true,
				"msg":        "get list success",
				"list":       data,
				"has_next":   hasNext,
				"total_page": totalPage,
				"refer":      nil,
			}
		}
	} else {
		if status != "audit" && status != "ok" && status != "deleted" {
			this.Data["json"] = map[string]interface{}{
				"result":     false,
				"msg":        "invalid status",
				"list":       nil,
				"has_next":   false,
				"total_page": 0,
				"refer":      nil,
			}
		} else {

			pro := models.ProblemBank{}
			data, hasNext, totalPage, err := pro.ListProblem(page, 20, status)

			if err != nil {
				this.Data["json"] = map[string]interface{}{
					"result":     false,
					"msg":        "get list failing",
					"list":       nil,
					"has_next":   false,
					"total_page": 0,
					"refer":      nil,
				}
			} else {
				log.Blueln(data)

				this.Data["json"] = map[string]interface{}{
					"result":     true,
					"msg":        "get list success",
					"list":       data,
					"has_next":   hasNext,
					"total_page": totalPage,
					"refer":      nil,
				}
			}
		}
	}

	this.ServeJson()
}

// check problem title exist
type ProblemTitleExistController struct {
	controllers.BaseController
}

func (this *ProblemTitleExistController) Get() {
	this.Data["json"] = map[string]interface{}{
		"result": true,
		"msg":    "get list success",
		"list":   nil,
		"refer":  nil,
	}

	this.ServeJson()
}

// problem delete
type ProblemDeleteController struct {
	controllers.BaseController
}

// uri /api/problem/delete/:id
func (this *ProblemDeleteController) Get() {
	id, err := this.GetInt("id")
	if nil != err || id < 0 {
		id = 0
	}

	s := this.Ctx.Input.Param(":id")
	idParm, err := strconv.Atoi(s)
	if nil != err || idParm < 0 {
		idParm = 0
	} else {
		id = idParm
	}

	if 0 == id {
		id = 1
	}

	pro := models.Problem{}
	err = pro.TrashProblem(id)

	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"result": false,
			"msg":    "get list failed",
			"debug":  err,
			"refer":  nil,
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"result": true,
			"msg":    "get list success",
			"refer":  nil,
		}
	}

	this.ServeJson()
}

// problem add
type ProblemAddController struct {
	controllers.BaseController
}

func (this *ProblemAddController) Get() {
	this.Data["json"] = map[string]interface{}{
		"result": false,
		"msg":    "only post method support",
		"refer":  nil,
	}

	this.ServeJson()
}

func (this *ProblemAddController) Post() {
	title := this.GetString("title")
	description := this.GetString("description")
	ptype := this.GetString("type")
	precode := this.GetString("precode")
	input := this.GetString("input")
	output := this.GetString("output")
	tags := this.GetString("tags")
	tags = utils.TagsCheck(tags)

	email := this.GetString("email")

	// decode iodata
	_, iodata, _ := utils.DecodeIoData(input, output)

	// decode description
	description = utils.Markdown2HTML(description)

	// store into database
	pro := models.Problem{}
	id, err := pro.AddProblem(title, ptype, description, precode, iodata, tags)

	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"result": false,
			"msg":    "add problem failed",
			"refer":  nil,

			"title":       title,
			"description": description,
			"ptype":       ptype,
			"precode":     precode,
			"input":       input,
			"output":      output,
			"email":       email,

			"debug": err,
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"result": true,
			"msg":    "add problem success",
			"refer":  nil,

			"id":          id,
			"title":       title,
			"description": description,
			"ptype":       ptype,
			"precode":     precode,
			"input":       input,
			"output":      output,
			"email":       email,
		}
	}

	this.ServeJson()
}
