package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"srv/aibot/models"
	"srv/aibot/db"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)


func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := db.UserCollection.CountDocuments(ctx, bson.M{"email": input.Email})
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	if count > 0 {
		http.Error(w, "Email already in use", http.StatusBadRequest)
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Could not hash password", http.StatusInternalServerError)
		return
	}
	input.Password = string(hashed)

	_, err = db.UserCollection.InsertOne(ctx, input)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered"))
}

func Login(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
        return
    }

    var input models.User
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var user models.User
    err := db.UserCollection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&user)
    if err != nil {
        http.Error(w, "User not found", http.StatusUnauthorized)
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
    if err != nil {
        http.Error(w, "Invalid password", http.StatusUnauthorized)
        return
    }

    sessionToken := generateSessionToken()

	user.UUID = sessionToken

    http.SetCookie(w, &http.Cookie{
        Name:     "session_token",
        Value:    sessionToken,
        Path:     "/",
        HttpOnly: true,
        Secure:   true, 
        SameSite: http.SameSiteStrictMode,
        MaxAge:   3600,
    })

    w.Write([]byte("Login successful"))
}

func generateSessionToken() string {
    return uuid.NewString()
}


func auth(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("session_token")
        if err != nil {
            http.Error(w, "Unauthorized - No session token", http.StatusUnauthorized)
            return
        }

        sessionUUID := cookie.Value
        if sessionUUID == "" {
            http.Error(w, "Unauthorized - Empty session token", http.StatusUnauthorized)
            return
        }

        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        var user models.User
        err = db.UserCollection.FindOne(ctx, bson.M{"uuid": &sessionUUID}).Decode(&user)
        if err != nil {
            http.Error(w, "Unauthorized - Invalid session", http.StatusUnauthorized)
            return
        }

        ctx = context.WithValue(r.Context(), "user", user)
        r = r.WithContext(ctx)

        next(w, r)
    }
}

