package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

// 初始化一个cookie存储对象
// something-very-secret应该是一个你自己的密匙，只要不被别人知道就行
var store = sessions.NewCookieStore([]byte("something-very-secret"))

//存储session
func saveSession(w http.ResponseWriter, r *http.Request) {

	//　获取一个session对象，session-name是session的名字
	session, err := store.Get(r, "session_name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//在session中存储键值对
	session.Values["name"] = "dylan"
	session.Values["age"] = 20
	session.Save(r, w)
}

//get session
func getSession(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "session_name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := session.Values["name"]
	age := session.Values["age"]
	fmt.Println("name:", name, "age:", age)
}

//将session的最大存储时间设置为小于零的数即为删除
func delSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session_name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options.MaxAge = -1
	session.Save(r, w)
}
func main() {
	http.HandleFunc("/save", saveSession)
	http.HandleFunc("/get", getSession)
	http.HandleFunc("/del", delSession)
	http.ListenAndServe(":8080", nil)

}
