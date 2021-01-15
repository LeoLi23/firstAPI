package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Student struct {
	Id int `orm:"column(Id)"`
	Name string `orm:"column(Name)"`
	Gender bool `orm:"column(Gender)"`
	Score int `orm:"column(Score)"`
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

func UpdateStudent(student *Student) {
	o := orm.NewOrm()
	_ = o.Using("default")
	_, _ = o.Update(student)
}

func DeleteStudent(id int) {
	o := orm.NewOrm()
	o.Using("default")
	o.Delete(&Student{Id:id})
}


func init() {
	orm.RegisterModel(new(Student))
}
