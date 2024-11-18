package main

import (
	"fmt"
	"koriebruh/EWallet/cnf"
	"log/slog"
	"net/http"
)

func main() {
	config, err := cnf.LoadConfig()
	if err != nil {
		slog.BoolValue(false)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /hi", HandleHi)

	fmt.Println("<=== SERVER RUNNING ===>")
	if err := http.ListenAndServe(config.Server, mux); err != nil {
		slog.Error("connection down")
		return
	}
}

func HandleHi(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Guest"
	}
	response := fmt.Sprintf("Welcome %s", name)
	w.Write([]byte(response))
}
