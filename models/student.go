package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)


type Student struct {
	Id int `json:"id" orm:"column(Id);auto"`
	Username string `json:"Username" orm:"column(Username);size(128)"`
	Password string `json:"Password" orm:"column(Password);size(128)"`
	Salt string `json:"Salt" orm:"column(Salt);size(128)"`
}

// LoginRequest defines login request format
type LoginRequest struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

// LoginResponse defines login response
type LoginResponse struct {
	Username    string             `json:"Username"`
	UserID      int                `json:"userID"`
	Token       string             `json:"token"`
}

//CreateRequest defines create user request format
type CreateRequest struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

//CreateResponse defines create user response
type CreateResponse struct {
	UserID   int    `json:"userID"`
	Username string `json:"Username"`
}

func GetAllStudents() []*Student{
	o := orm.NewOrm()
	_ = o.Using("default")
	var students []*Student 
	query := o.QueryTable("student")
	_, _ = query.All(&students)
	
	return students
}

func GetStudentById(id int) (Student,error) {
	u := Student{Id:id}
	o := orm.NewOrm()
	_ = o.Using("default")
	err := o.Read(&u)
	
	if err == orm.ErrNoRows {
		fmt.Println("Can't find it!")
	} else if err == orm.ErrMissPK {
		fmt.Println("No Primary key!")
	}
	
	return u,err
}

func AddStudent(student *Student) Student {
	o := orm.NewOrm()
	_ = o.Using("default")
	_, _ = o.Insert(student)
	
	return *student 
}

func UpdateStudent(student *Student) (Student, error ){
	o := orm.NewOrm()
	_ = o.Using("default")
	_, err := o.Update(student)
	return *student, err
}

func DeleteStudent(id int) {
	o := orm.NewOrm()
	o.Using("default")
	o.Delete(&Student{Id:id})
}

func DoLogin(username, password string) bool {
	o := orm.NewOrm()
	o.Using("default")
	toread := &Student{Username:username, Password: password}
	err := o.Read(toread,"Username","Password")

	if err != nil {
		return false
	}
	return true
}

