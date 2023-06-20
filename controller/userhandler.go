package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"html/template"
	"net/http"
)

//login 处理用户登陆函数
func Login(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登陆
	flag, _ := dao.IsLogin(r)
	if flag {
		//去首页
		GetPageBooksByPrice(w, r)
	} else {
		//获取用户名和密码
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		//调用userdao的CheckUserNameAndPassword方法
		user, _ := dao.CheckUserNameAndPassword(username, password)
		if user.ID != 0 {
			//用户名密码正确
			//生成UUID传给session
			uuid := utils.CreateUUID()
			//创建session
			session := &model.Session{
				SessionID: uuid,
				UserName:  username,
				UserID:    user.ID,
			}
			//将session保存到数据库
			err := dao.AddSession(session)
			if err != nil {
				fmt.Printf("登陆操作中插入session出错: %v\n", err)
			}
			//创建cookie与session关联
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)

			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			//执行
			t.Execute(w, user)

		} else {
			//用户名密码不正确
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			//执行
			t.Execute(w, "用户名或密码不正确！")
		}

	}

}

func Logout(w http.ResponseWriter, r *http.Request) {
	//获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		dao.DeleteSession(cookieValue)
		//设置cookie失效
		cookie.MaxAge = -1 //maxge<0表示立即删除cookie
		//发送到浏览器
		http.SetCookie(w, cookie)
	}
	//去首页
	GetPageBooksByPrice(w, r)
}

//regist 处理用户注册函数
func Regist(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	//调用userdao的CheckUserName方法
	user, _ := dao.CheckUserName(username)
	if user.ID != 0 {
		//用户名已存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在！")

	} else {
		//注册 调用SaveUser方法
		dao.SaveUser(username, password, email)
		//注册成功页面
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "注册成功")
	}

}

//通过ajax验证用户名是否可用
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	fmt.Println("传入的用户名是：")
	user, _ := dao.CheckUserName(username)
	if user.ID != 0 {
		//用户名已存在
		w.Write([]byte("用户名已存在！"))

	} else {
		// 用户名可用
		w.Write([]byte("<font style='color:green'>用户名可用！</font>"))

	}

}
