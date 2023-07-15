package main

import (
	"app/controller"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			controller.Get_Users(w, r)
		} else if r.Method == "PUT" {
			controller.Update_User(w, r)
		} else if r.Method == "DELETE" {
			controller.Delete_User(w, r)
		} else if r.Method == "POST" {
			controller.Create_User(w, r)
		}
	})

	fmt.Println("Listening 4000 port...")

	// http://localhost:4000/user

	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		panic(err)
	}
}
