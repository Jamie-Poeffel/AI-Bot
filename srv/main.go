package main

import (
	"fmt"
	"net/http"

	"srv/aibot/logger"
	"srv/aibot/controllers"
	"srv/aibot/db"
)

func main() {
	db.Connect()
	http.HandleFunc("/newUser", logger.RequestHandler(controllers.Register))
	http.HandleFunc("/login", logger.RequestHandler(controllers.Login))

	fmt.Println("Server läuft auf http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.ErrorLog.Printf("Fehler beim Starten des Servers: %v", err)
	}
}
