package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"sme-education-backend/internal/pkg/utils"

	"sme-education-backend/cmd/entity-server/models"
)

type UserReq struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FullName  string `json:"full_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`

	PageNum  int
	PageSize int
}
type UserRes struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Avatar    string    `json:"avatar"`
	UserRole  string    `json:"user_role"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

func toUserRes(user *models.User) *UserRes {
	var userRes = &UserRes{
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
		UserRole: user.UserRole,
	}
	return userRes
}

func (obj *UserReq) Login() (*UserRes, error) {
	model := models.User{
		Email: obj.Email,
	}
	user, err := model.Login()
	if err != nil {
		return nil, err
	}
	if !utils.CheckPasswordHash(obj.Password, user.Password) {
		return nil, errors.New("Login failed")
	}
	return toUserRes(user), nil
}

func (obj *UserReq) Register() (bool, error) {
	if strings.TrimSpace(obj.Password) == "" {
		return false, errors.New("Password không được để trống")
	}
	passHash, hashErr := utils.HashPassword(obj.Password)
	if hashErr != nil {
		return false, hashErr
	}
	model := models.User{
		Email:     obj.Email,
		Password:  passHash,
		FullName:  fmt.Sprintf("%s %s", obj.FirstName, obj.LastName),
		FirstName: obj.FirstName,
		LastName:  obj.LastName,
		Phone:     obj.Phone,
		UserRole:  "user",
	}
	isExists, _ := model.IsEmailExist()
	if isExists {
		return false, errors.New("Email đã tồn tại")
	}
	if _, err := model.Register(); err == nil {
		return true, err
	} else {
		return false, err
	}
}
func (obj *UserReq) Get() (*UserRes, error) {
	model := models.User{}
	objRes, err := model.FindOne(obj.ID)
	if err != nil {
		return nil, err
	}
	return toUserRes(objRes), nil
}
func (obj *UserReq) GetAllUser() ([]*UserRes, error) {
	model := models.User{}
	list, err := model.GetAll(obj.PageNum, obj.PageSize)
	if err != nil {
		return nil, err
	}
	var res []*UserRes
	for _, item := range list {
		res = append(res, toUserRes(item))
	}
	return res, nil
}
func (obj *UserReq) GetTotal() (int, error) {
	model := models.User{}
	return model.GetTotal()
}
func (obj *UserReq) AddUser() (bool, error) {
	passHash, hashErr := utils.HashPassword(obj.Password)
	if hashErr != nil {
		return false, hashErr
	}
	model := models.User{
		Email:     obj.Email,
		Password:  passHash,
		FullName:  obj.FullName,
		FirstName: obj.FirstName,
		LastName:  obj.LastName,
		Phone:     obj.Phone,
		Avatar:    obj.Avatar,
	}
	_, err := model.Add()
	if err != nil {
		return false, err
	}
	return true, nil
}
func (obj *UserReq) UpdateUser() (*UserRes, error) {
	model := models.User{
		Email:     obj.Email,
		FullName:  obj.FullName,
		FirstName: obj.FirstName,
		LastName:  obj.LastName,
		Phone:     obj.Phone,
		Avatar:    obj.Avatar,
	}
	objRes, err := model.Update(obj.ID)
	if err != nil {
		return nil, err
	}
	return toUserRes(objRes), nil
}
func (obj *UserReq) DeleteUser() (*UserRes, error) {
	model := models.User{}
	objRes, err := model.Delete(obj.ID)
	if err != nil {
		return nil, err
	}
	return toUserRes(objRes), nil
}
