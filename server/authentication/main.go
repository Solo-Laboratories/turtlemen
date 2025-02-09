package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	sloghttp "github.com/samber/slog-http"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "auth-server",
		Usage: "Authentication Server for Turtlemen Application.",
		Action: func(context.Context, *cli.Command) error {
			start()
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		slog.Info(err.Error())
	}
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"` // We will need to hash this
}

type Session struct {
	Token     string    `json:"token"`
	UserID    string    `json:"userId"`
	ExpiresAt time.Time `json:"expiresAt"`
}

var users = map[string]User{
	"testuser": {Username: "testuser", Password: "password"}, // This will be stored in a database?
}

var sessions = make(map[string]Session)

func start() {
	http.HandleFunc("/login", loginHandler)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	config := sloghttp.Config{
		WithSpanID:  true,
		WithTraceID: true,
	}

	mux := http.NewServeMux()

	// Middleware
	handler := sloghttp.Recovery(mux)
	handler = sloghttp.NewWithConfig(logger, config)(handler)

	// Start server
	http.ListenAndServe(":8080", handler)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	storedUser, ok := users[user.Username]
	if !ok || storedUser.Password != user.Password { // In real app, compare hashes!
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate Session Token
	token := uuid.New().String()
	expiresAt := time.Now().Add(time.Hour * 24) // Example: 24-hour expiry

	session := Session{
		Token:     token,
		UserID:    user.Username, // Use a proper user ID in a real app
		ExpiresAt: expiresAt,
	}

	sessions[token] = session

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
