package main

import (
	"fmt"
	"minibank/handlers"
	"minibank/models"
	"net/http"
	"os"
	"time"
)

func main() {
	// Connect to the database
	time.Sleep(1000)
	dbConn := fmt.Sprintf("minibank:minibank@tcp(mysql)/minibank")
	models.InitDB(dbConn)
	defer models.Database.Close()
	http.HandleFunc("/api/account/register", handlers.RegisterHandler)
	// http.HandleFunc("/api/account/token", handlers.TokenHandler)
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
