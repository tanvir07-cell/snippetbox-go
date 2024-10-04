package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"snippetbox.tanvirRifat.io/internal/models"
)

// using object or struct handler

// type home struct{}

// func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request){
// 	w.Write([]byte("Home Page using struct"))

// }



 

func (app *application) home(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Server","Go")



snippets, err := app.snippets.Latest()
if err != nil {
 app.serverError(w,r,err)
}

	files:= []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	t,err:= template.ParseFiles(files...)

	

	if err!=nil{
		app.serverError(w,r,err)
		return
	}

	

	data:= templateData{
		Snippets: &snippets,
	}

	err = t.ExecuteTemplate(w,"base",data)

	if err!=nil{
		 app.serverError(w,r,err)
		 return
	}



}


func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
id, err := strconv.Atoi(r.PathValue("id"))
if err != nil || id < 1 {
http.NotFound(w, r)
return
}

snippet, err := app.snippets.Get(id)
if err != nil {
if errors.Is(err, models.ErrNoRecord) {
http.NotFound(w, r)
} else {
	app.serverError(w,r,err)

}
return
}

files:= []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/view.tmpl.html",
}

t,err:= template.ParseFiles(files...)

if err!=nil{
	app.serverError(w,r,err)
	return
}

data := templateData{
Snippet: &snippet,
}
err = t.ExecuteTemplate(w,"base",data)

if err!=nil{
	app.serverError(w,r,err)
	return
}




}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Create a new snippet"))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request){

	title := "O snail"
content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
expires := 7
// Pass the data to the SnippetModel.Insert() method, receiving the
// ID of the new record back.
id, err := app.snippets.Insert(title, content, expires)
if err != nil {
	app.serverError(w,r,err)
return
}
// Redirect the user to the relevant page for the snippet.
http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}

 

