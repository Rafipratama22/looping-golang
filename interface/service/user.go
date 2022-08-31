package service

import (
	"fmt"
	"sync"
)

type User struct {
	Nama string
}

type UserService struct {
	ListUser []User
}

type UserIface interface {
	Register(u *User, wg *sync.WaitGroup)
	Get(wg *sync.WaitGroup)
}

func NewUserService(user []User) UserIface {
	return &UserService{
		ListUser: user,
	}
}

func (u *UserService) Register(user *User, wg *sync.WaitGroup) {
	u.ListUser = append(u.ListUser, *user)
	wg.Done()
}

func (u *UserService) Get(wg *sync.WaitGroup) {
	for _, each := range u.ListUser {
		fmt.Println(each.Nama)
	}
	wg.Done()
}