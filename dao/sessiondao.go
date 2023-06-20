package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"net/http"
)

//向数据库中添加session
func AddSession(sess *model.Session) error {
	//sql语句
	sqlStr := "insert into sessions values(?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, sess.SessionID, sess.UserName, sess.UserID)
	if err != nil {
		fmt.Println("添加session出错")
		return err
	}
	return nil
}

//删除session
func DeleteSession(sessID string) error {
	sqlStr := "delete from sessions where session_id = ?"
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		fmt.Println("删除session出错")
		return err
	}
	return nil
}

//根据sessionID获取session信息
func GetSessionByID(sessionID string) (*model.Session, error) {
	sqlStr := "select * from sessions where session_id = ?"

	row := utils.Db.QueryRow(sqlStr, sessionID)
	session := &model.Session{}
	row.Scan(&session.SessionID, &session.UserName, &session.UserID)
	return session, nil

}
func IsLogin(r *http.Request) (bool, *model.Session) {
	//根据Cookie的name获取Cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取Cookie的value
		cookieValue := cookie.Value
		//根据cookieValue去数据库中查询与之对应的Session
		session, _ := GetSessionByID(cookieValue)
		if session.UserID > 0 {
			//已经登录
			return true, session
		}
	}
	//没有登录
	return false, nil
}
