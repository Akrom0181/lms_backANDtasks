package main

import (
	"fmt"
	"net/http"
	"lms_back/config"
	"lms_back/storage"
	"lms_back/controller"
)

func main() {
	cfg := config.Load()
	store, err := storage.New(cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.DB.Close()

	con := controller.NewController(store)

	http.HandleFunc("/branch", con.Branch)
	http.HandleFunc("/teacher", con.Teacher)
	http.HandleFunc("/student", con.Student)
	http.HandleFunc("/group", con.Group)

	fmt.Println("programm is running on localhost:8080...")
	http.ListenAndServe(":8080", nil)

}
