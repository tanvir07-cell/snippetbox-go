package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// using object or struct handler

// type home struct{}

// func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request){
// 	w.Write([]byte("Home Page using struct"))

// }



 

func home(w http.ResponseWriter, r *http.Request){
	files:= []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	t,err:= template.ParseFiles(files...)

	if err!=nil{
		log.Print(err.Error())
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w,"base",nil)

	if err!=nil{
		log.Print(err.Error())
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
		return
	}


}


func snippetView(w http.ResponseWriter, r *http.Request){
	// convert string id to number using strconv.Atoi
	id,err:= strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w,r)
		return
	}

	msg:= fmt.Sprintf("Displaying details of snippet %d",id)
	w.Write([]byte(msg))


}

func snippetCreate(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Create a new snippet"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request){

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Create a new snippet for posting..."))
}

 

