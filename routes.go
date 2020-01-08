package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func setStaticFolder(route *mux.Router) {
	fs := http.FileServer(http.Dir("./public/"))
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
}

func addApproutes(route *mux.Router) {

	setStaticFolder(route)

	route.HandleFunc("/user", getUsers).Methods("GET")
	route.HandleFunc("/register",registerUser).Methods("POST")
	route.HandleFunc("/login",loginUser).Methods("POST")
	route.HandleFunc("/delete/{username}", deleteUser).Methods("DELETE")
	route.HandleFunc("/user", updateUser).Methods("PUT")
	route.HandleFunc("/validateToken/{token}", validateToken).Methods("POST")
	fmt.Println("Routes loading is completed")
}
