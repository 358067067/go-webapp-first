package model

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	// fmt.Println("新增用户测试")
	// user := User{
	// Name: "admin",
	// }
	// user.AddUser()
	// fmt.Println("查询用户测试")
	// u, _ := user.GetUserByName()
	// fmt.Println(u)
	u := User{}
	users, _ := u.GetUsers()
	for k, v := range users {
		fmt.Printf("%v--------%v\n", k+1, v)
	}
}
