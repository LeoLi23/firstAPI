package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Student struct {
	Id int `orm:"column(Id)"`
	UserName string `orm:"column(UserName)"`
	Password string `orm:"column(Password)"`
	Gender bool `orm:"column(Gender)"`
	Score int `orm:"column(Score)"`
}

var StudentList map[string]*Student

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

func StudentLogin(username, password string) bool {
	o := orm.NewOrm()
	o.Using("default")
	toread := &Student{UserName:username, Password: password}
	err := o.Read(toread,"UserName","Password")

	if err != nil {
		return false
	}
	return true
}


func init() {
	orm.RegisterModel(new(Student))
}
