package main

//type ContactDetails struct {
// 	Email   string
// 	Subject string
// 	Message string
// }

//func main() {
// 	r := mux.NewRouter()
// 	tmpl := template.Must(template.ParseFiles("templates/form.html"))

// 	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		tmpl.Execute(w, nil)
// 	}).Methods("GET")

// 	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		details := ContactDetails{
// 			Email:   r.FormValue("email"),
// 			Subject: r.FormValue("subject"),
// 			Message: r.FormValue("message"),
// 		}

// 		fmt.Printf("Email: %s, Subject: %s Message: %s", details.Email, details.Subject, details.Message)

// 		tmpl.Execute(w, struct{ Success bool }{true})
// 	}).Methods("POST")

// 	http.ListenAndServe(":8081", r)
// }
