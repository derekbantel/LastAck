package main

import (
	"fmt"
	"net/http"

	"github.com/derekbantel/LastAck/pkgs"
)

func main() {
	router := pkgs.NewRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router.Routes(),
	}
	fmt.Println("Server started on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
