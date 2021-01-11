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

//AddUser 插入User数据
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
	_, err2 := inStmt.Exec(sqlStr, user.Name, user.Pwd, user.Email)
	if err2 != nil {
		fmt.Println("新增用户失败", err2)
		return err2
	}
	return nil
}

//GetUserByName 查询User数据
func (user *User) GetUserByName() (*User, error) {
	sqlStr := "select id, name, pwd ,email from users where name = ?"
	row := utils.Db.QueryRow(sqlStr, user.Name)
	// 变量赋值
	var id int
	var name string
	var pwd string
	var email string
	err := row.Scan(&id, &name, &pwd, &email)
	if err != nil {
		fmt.Println("查询异常", err)
		return nil, err
	}
	u := &User{
		ID:    id,
		Name:  name,
		Pwd:   pwd,
		Email: email,
	}
	return u, nil
}

//GetUsers 查询所有用户
func (user *User) GetUsers() ([]*User, error) {
	sqlStr := "select id, name, pwd, email from users"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		fmt.Println("查询异常", err)
		return nil, err
	}
	var users []*User
	for rows.Next() {
		// 变量赋值
		var id int
		var name string
		var pwd string
		var email string
		err := rows.Scan(&id, &name, &pwd, &email)
		if err != nil {
			fmt.Println("查询异常", err)
			return nil, err
		}
		u := &User{
			ID:    id,
			Name:  name,
			Pwd:   pwd,
			Email: email,
		}
		users = append(users, u)
	}
	return users, nil
}
