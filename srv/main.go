package main

import (
	"fmt"
	"net/http"
	"os"


	"srv/aibot/logger"
	"srv/aibot/controllers"
	"srv/aibot/db"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("PRODUCTION") == "" {
		if err := godotenv.Load(); err != nil {
			logger.Fatalf("Fehler beim Laden der .env-Datei: %v", err)
		}
	}
	
	db.Connect()

	
	router := mux.NewRouter()

	
	router.HandleFunc("/newUser", logger.RequestHandler(controllers.Register)).Methods("POST")
	router.HandleFunc("/login", logger.RequestHandler(controllers.Login)).Methods("POST")
	router.HandleFunc("/prompt", logger.RequestHandler(controllers.Auth(controllers.AI))).Methods("POST")

	
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, 
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	
	handler := corsMiddleware.Handler(router)

	
	fmt.Println("Server läuft auf http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		logger.ErrorLog.Printf("Fehler beim Starten des Servers: %v", err)
	}
}
