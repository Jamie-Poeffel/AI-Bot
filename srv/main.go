package main

import (
	"fmt"
	"net/http"

	"srv/aibot/logger"
)

func main() {
	http.HandleFunc("/", logger.RequestHandler)
	http.HandleFunc("/error", logger.ErrorHandler)

	fmt.Println("Server läuft auf http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.ErrorLog.Printf("Fehler beim Starten des Servers: %v", err)
	}
}
