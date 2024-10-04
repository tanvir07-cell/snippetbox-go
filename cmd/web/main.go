package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"snippetbox.tanvirRifat.io/internal/models"

	_ "github.com/go-sql-driver/mysql" // New import
)

// for the dependecy injection

type application struct{
	logger *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	



logger:= slog.New(slog.NewJSONHandler(os.Stdout,&slog.HandlerOptions{
	AddSource: true,
}))

// db connection:


 db,dbErr:=openDB()

if dbErr != nil {
logger.Error(dbErr.Error())
os.Exit(1)
}

defer db.Close()


// dependency injection: passing logger to application struct

app:= &application{
  logger: logger,	
	snippets: &models.SnippetModel{DB: db},
}








err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
	}


port:= os.Getenv("PORT")

logger.Info("Starting server on :"+port)

err = http.ListenAndServe(fmt.Sprintf(":%s",port),app.routes())
logger.Error(err.Error())

os.Exit(1)

}

func openDB() (*sql.DB,error){
	db,err:= sql.Open("mysql","web:pass@/snippetbox?parseTime=true")
	if err != nil {
		return nil,err
	}
	if err:= db.Ping(); err != nil {
		return nil,err
	}
	return db,nil
}