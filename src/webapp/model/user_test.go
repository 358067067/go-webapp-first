package model

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("新增用户测试")
	user := &User{}
	user.AddUser()
}
