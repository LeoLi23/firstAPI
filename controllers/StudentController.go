package controllers

import (
	"encoding/json"
	"firstAPI/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"net/http"
	"strconv"
)

// Operations about students
type UserController struct {
	beego.Controller
}

func (u *UserController) unmarshalPayload(v interface{}) error {
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, v); err != nil {
		logs.Error("unmarshal payload of %s error: %s", u.Ctx.Request.URL.Path, err)
	}
	return nil
}

func (u *UserController) respond(code int, message string, data ...interface{}) {
	u.Ctx.Output.SetStatus(code)
	var d interface{}
	if len(data) > 0 {
		d = data[0]
	}
	u.Data["json"] = struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}{
		Code:    code,
		Message: message,
		Data:    d,
	}
	u.ServeJSON()
}

func (u *UserController) GetAll() {
	ss := models.GetAllStudents()
	u.Data["json"] = ss
	fmt.Println(ss)
	u.ServeJSON()
}


//func (u *UserController) GetById() {
//	id, _ := u.GetInt(":id")
//	if id != 0 {
//		fmt.Println(id)
//		s,err := models.GetStudentById(id)
//		if err != nil {
//			u.Data["json"] = err.Error()
//		} else {
//			u.Data["json"] = s
//		}
//	}
//	u.ServeJSON()
//}

func (u *UserController) GetById(){
	gr := new(models.GetRequest)
	if err := u.unmarshalPayload(gr);err != nil {
		u.respond(http.StatusBadRequest,err.Error())
		return
	}

	grs, statuscode,err := models.GetStudentById(gr)
	if err != nil {
		u.respond(statuscode,err.Error())
		return
	}

	u.Ctx.Output.Header("Authorization",grs.Token)
	u.respond(http.StatusOK,"",grs)
}

func (u *UserController) Update() {
	id,_ := u.GetInt(":id")
	if id != 0 {
		var s models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &s)
		ss, err := models.UpdateStudent(&s)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = ss
		}
	}
	u.ServeJSON()
}

func (u *UserController) Delete() {
	id ,_:= u.GetInt(":id")
	models.DeleteStudent(id)
	u.Data["json"] = map[string]string{"status": "delete success", "data": strconv.Itoa(id)}
	u.ServeJSON()
}

func (u *UserController) Login(){
	lr := new(models.LoginRequest)

	if err := u.unmarshalPayload(lr); err != nil {
		u.respond(http.StatusBadRequest, err.Error())
		return
	}

	lrs, statusCode, err := models.DoLogin(lr)
	if err != nil {
		u.respond(statusCode,err.Error())
		return
	}

	u.Ctx.Output.Header("Authorization",lrs.Token)//set token into header
	u.respond(http.StatusOK,"",lrs)

}

func (u *UserController) CreateUser() {
	cu := new(models.CreateRequest)

	if err := u.unmarshalPayload(cu); err != nil {
		u.respond(http.StatusBadRequest, err.Error())
	}

	createUser, statusCode, err := models.DoCreateUser(cu)
	if err != nil {
		u.respond(statusCode, err.Error())
		return
	}

	u.respond(http.StatusOK, "", createUser)
}

