package repo

import (
	"fmt"
	"reflect"
)

type UserServiceInterface interface {
	CreateUser()
	DeleteUser()
	UpdateUser()
	GetUser()
	ListsUser()
}

type users struct{}

var us *users

func Users() UserServiceInterface {
	return us
}

func (u *users) CreateUser(u interface{}) error {

	return nil
}

/*
func (u *UserRepo) CreateUser(u interface{}) error
*/
