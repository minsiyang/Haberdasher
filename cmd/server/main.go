package main

import (
	"net/http"

	"github.com/twirp-play/internal/haberdasherserver"
	"github.com/twirp-play/rpc/haberdasher"
)

func main() {
	server := &haberdasherserver.Server{} // implements Haberdasher interface
	twirpHandler := haberdasher.NewHaberdasherServer(server)

	http.ListenAndServe(":8080", twirpHandler)
}
