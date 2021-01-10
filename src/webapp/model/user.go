package model

import (
	"fmt"
	"webapp/utils"
)

//User 结构体
type User struct {
	ID    int
	Name  string
	Pwd   string
	Email string
}

//插入User数据
func (user *User) AddUser() error {
	// sql语句
	sqlStr := "insert into users(name, pwd, email) values(?,?,?)"
	// 预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常", err)
		return err
	}
	// 执行插入语句
	_, err2 := inStmt.Exec("admin", "123456", "xxhundan@qq.com")
	// _, err2 := inStmt.Exec(sqlStr, "admin", "123456", "xxhundan@qq.com")
	if err2 != nil {
		fmt.Println("新增用户失败", err2)
		return err2
	}
	return nil
}
