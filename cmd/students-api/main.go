package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JayShenkar07/students-api/internal/config"
)

func main() {
	fmt.Println("Hello")

	// Load configuration
	cfg := config.MustLoad()

	// // Print loaded config values for confirmation
	// fmt.Printf("Running in environment: %s\n", cfg.Env)
	// fmt.Printf("Storage Path: %s\n", cfg.Storage)
	fmt.Printf("HTTP Server Address: %s\n", cfg.Addr)

	router := http.NewServeMux()
	router.HandleFunc("GET /api/student/home", userHome)
	router.HandleFunc("GET /api/student/profile", userProfile)
	
	server := http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}

	fmt.Println("Server Started!")
	err := server.ListenAndServe()
	if err != nil{
		log.Fatal("Failed to start server!")
	}
	
}

func userHome(writer http.ResponseWriter, request *http.Request) {
	var message string = "api/student/home"
	fmt.Fprint(writer, message)
}

func userProfile(writer http.ResponseWriter, request *http.Request) {
	var message string = "api/student/profile"
	fmt.Fprint(writer, message)
}
