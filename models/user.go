package models

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"log"
	"net/http"
)


type User struct {
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
	UserID      int                `json:"UserID"`
	Token       string             `json:"token"`
}

//CreateRequest defines create user request format
type CreateRequest struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

//CreateResponse defines create user response
type CreateResponse struct {
	UserID   int    `json:"UserID"`
	Username string `json:"Username"`
}

type GetRequest struct {
	Username string `json:"Username"`
}

type GetResponse struct {
	UserID int `json:"UserID"`
	Username string `json:"Username"`
	Token string `json:"token"`
}

func GetAllStudents() []*User{
	o := orm.NewOrm()
	_ = o.Using("default")
	var students []*User
	query := o.QueryTable("User")
	_, _ = query.All(&students)
	
	return students
}

func GetStudentByName(gr *GetRequest) (*User, int, error){
	log.Println("GetStudentByName starts")
	name := gr.Username
	o := orm.NewOrm()
	user := &User{Username:name}
	err := o.Read(user, "Username")

	if err != nil {
		return nil, http.StatusBadRequest, errors.New("error: username doesn't exist")
	}
	log.Println("GetStudentByName ends")
	return user, http.StatusOK, nil
}

func UpdateStudent(student *User) (User, error ){
	o := orm.NewOrm()
	_ = o.Using("default")
	_, err := o.Update(student)
	return *student, err
}

func DeleteStudent(id int) {
	o := orm.NewOrm()
	o.Using("default")
	o.Delete(&User{Id:id})
}

func DoLogin(lr *LoginRequest) (*LoginResponse, int, error){
	// get username and password
	username := lr.Username
	password := lr.Password

	// validate user name and password is they are empty
	if len(username) == 0 || len(password) == 0 {
		return nil, http.StatusBadRequest,errors.New("error: username or password is empty")
	}

	o := orm.NewOrm()

	// check if the username exists
	user := &User{Username: username}
	err := o.Read(user,"Username")
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("error: username doesn't exist")
	}

	// generate the password hash
	hash, err := GeneratePassHash(password,user.Salt)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	if hash != user.Password {
		return nil, http.StatusBadRequest,errors.New("error: password is error")
	}

	// generate token
	tokenString, err := GenerateToken(lr, user.Id, 0)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return &LoginResponse{
		Username: user.Username,
		UserID: user.Id,
		Token: tokenString,
	},http.StatusOK,nil
}

func DoCreateUser(cr *CreateRequest)(*CreateResponse,int,error){
	o := orm.NewOrm()

	// check if username exists
	userNameCheck := User{Username: cr.Username}
	err := o.Read(&userNameCheck,"Username")
	if err == nil {
		return nil, http.StatusBadRequest, errors.New("username has already existed")
	}

	//generate salt
	saltKey, err := GenerateSalt()
	if err != nil {
		logs.Info(err.Error())
		return nil, http.StatusBadRequest, err
	}

	// generate password hash
	hash, err := GeneratePassHash(cr.Password,saltKey)
	if err != nil {
		logs.Info(err.Error())
		return nil, http.StatusBadRequest,err
	}

	// create user
	user := User{}
	user.Username = cr.Username
	user.Password = hash
	user.Salt = saltKey

	_, err = o.Insert(&user)
	if err != nil {
		logs.Info(err.Error())
		return nil, http.StatusBadRequest,err
	}

	return &CreateResponse{
		UserID:user.Id,
		Username: user.Username,
	}, http.StatusOK,nil
}







