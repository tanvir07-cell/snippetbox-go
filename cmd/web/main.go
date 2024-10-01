package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)
func main() {
mux := http.NewServeMux()

fileServer:= http.FileServer(http.Dir("./ui/static/"))

mux.Handle("/static/",http.StripPrefix("/static",fileServer))


mux.HandleFunc("GET /{$}", home)

mux.HandleFunc("GET /snippet/view/{id}",snippetView)
mux.HandleFunc("GET /snippet/create",snippetCreate)
mux.HandleFunc("POST /snippet/create",snippetCreatePost)

logger:= slog.New(slog.NewJSONHandler(os.Stdout,&slog.HandlerOptions{
	AddSource: true,
}))



err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
	}


port:= os.Getenv("PORT")

logger.Info("Starting server on :"+port)

err = http.ListenAndServe(":4000",mux)
logger.Error(err.Error())

}