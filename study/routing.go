// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func main() {
// 	r := mux.NewRouter()

// 	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) { //주소값에서 임의로 데이터 설정하여 전송 가능
// 		vars := mux.Vars(r) //mux.Vars() 메서드를 통해 URL의 Path Parmerer를 가져올 수 있음
// 		title := vars["title"]
// 		page := vars["page"]

// 		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
// 	}).Methods("GET") //이런 식으로 post ,get 인지 구별

// 	http.ListenAndServe(":8081", r)
// }
