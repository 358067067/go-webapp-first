package main

import (
	"encoding/json"
	"net/http"
	"text/template"
	"webapp/model"
)

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	// t, _ := template.ParseFiles("pages/index.html")
	// Must函数调用会自动处理异常
	t := template.Must(template.ParseFiles("pages/index.html", "pages/index2.html"))
	// t.Execute(w, "占位符")
	// 这样写，响应数据会在index2中显示
	user := model.User{}
	u, _ := user.GetUsers()
	json, _ := json.Marshal(u)
	t.ExecuteTemplate(w, "index2.html", json)
}

func main() {
	http.HandleFunc("/index", htmlHandler)
	http.ListenAndServe(":8080", nil)
}
