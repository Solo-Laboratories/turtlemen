package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	sloghttp "github.com/samber/slog-http"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "turtlemen-server",
		Usage: "Main Server for Turtlemen Application.",
		Action: func(context.Context, *cli.Command) error {
			start()
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		slog.Info(err.Error())
	}
}

func start() {
	http.HandleFunc("/validate", validateHandler)

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

type Session struct {
	Token     string    `json:"token"`
	UserID    string    `json:"userId"`
	ExpiresAt time.Time `json:"expiresAt"`
}

var sessions = make(map[string]Session)

func validateHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token is required", http.StatusBadRequest)
		return
	}

	session, ok := sessions[token]
	if !ok || session.ExpiresAt.Before(time.Now()) {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Token is valid")
}
