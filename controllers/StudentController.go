package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"encoding/json"
	"firstAPI/models"
)

// Operations about students
type StudentController struct {
	beego.Controller
}


func (u *StudentController) GetAll() {
	ss := models.GetAllStudents()
	u.Data["json"] = ss
	fmt.Println(ss)
	u.ServeJSON()
}


func (u *StudentController) GetById() {
	id, _ := u.GetInt(":id")
	if id != 0 {
		fmt.Println(id)
		s,err := models.GetStudentById(id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = s
		}
	}
	u.ServeJSON()
}

func (u *StudentController) Post() {
	var s models.Student
	json.Unmarshal(u.Ctx.Input.RequestBody, &s)
	uid := models.AddStudent(&s)
	u.Data["json"] = uid
	fmt.Println(uid)
	u.ServeJSON()
}


// @Title 修改用户
// @Description 修改用户的内容
// @Param      body          body   models.Student true          "body for user content"
// @Success 200 {int} models.Student
// @Failure 403 body is empty
// @router / [put]
func (u *StudentController) Update() {
	var s models.Student
	json.Unmarshal(u.Ctx.Input.RequestBody, &s)
	models.UpdateStudent(&s)
	u.Data["json"] = s
	u.ServeJSON()
}

// @Title 删除一个学生
// @Description 删除某学生数据
// @Param      id            path   int    true          "The key for staticblock"
// @Success 200 {object} models.Student
// @router /:id [delete]
func (u *StudentController) Delete() {
	id ,_:= u.GetInt(":id")
	models.DeleteStudent(id)
	u.Data["json"] = true
	u.ServeJSON()
}
