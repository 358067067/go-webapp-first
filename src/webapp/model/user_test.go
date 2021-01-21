package model

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	u := User{}
	users, _ := u.GetUsers()
	for k, v := range users {
		fmt.Printf("%v--------%v\n", k+1, v)
	}
}

func TestAddUser(t *testing.T) {
	fmt.Println("新增用户测试")
	user := User{
		Name: "admin",
	}
	user.AddUser()
}
