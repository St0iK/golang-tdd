package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, GetPlayerScore(player))
}

func GetPlayerScore(player string) string {
	if player == "Pepper" {
		return "20"
	}

	if player == "Floyd" {
		return "100"
	}
	return ""
}
