package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

//检查用户名和密码,根据用户名和密码在数据库中查询和记录
func CheckUserNameAndPassword(username, password string) (*model.User, error) {
	//sql语句
	sqlStr := "select id,username,password,email from users where username = ? and password = ?"
	//执行 //扫描
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

//验证用户名
func CheckUserName(username string) (*model.User, error) {
	//sql语句
	sqlStr := "select id,username,password,email from users where username = ?"
	//执行 //扫描
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil

}

// 注册
func SaveUser(username, password, email string) error {
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	// 执行//Exec执行一次命令（包括查询、删除、更新、插入等），不返回任何执行结果。参数args表示query中的占位参数。
	_, err := utils.Db.Exec(sqlStr, username, password, email)
	if err != nil {
		fmt.Println("SaveUser方法插入用户失败！")
		panic(err.Error())
	}
	return nil

}
