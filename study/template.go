// package main

// import (
// 	"net/http"
// 	"text/template"
// )

// type Todo struct {
// 	Title string
// 	Done  bool
// }

// type TodoPageData struct {
// 	PageTitle string
// 	Todos     []Todo
// }

// func main() {
// 	tmpl := template.Must(template.ParseFiles("templates/layout.html"))
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //요청 핸들러 등록
// 		data := TodoPageData{
// 			PageTitle: "데이터 베이스 연동 예제_이진구",
// 			Todos: []Todo{
// 				{Title: "회원정보", Done: false},
// 				{Title: "Task 2", Done: true},
// 				{Title: "Task 3", Done: true},
// 			},
// 		}
// 		tmpl.Execute(w, data)
// 	})

// 	http.ListenAndServe(":8081", nil) //포트 설정
// }
