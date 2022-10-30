package main

import (
	"fmt"

	"github.com/BON4/payment/internal/server"
)

// @title           Payments API
// @version         1.0
// @description     This service provides loading/uploading csv file into/from DB.

// @host      localhost:8080
// @BasePath  /
func main() {
	s, err := server.NewServer(".")
	if err != nil {
		fmt.Printf("INIT ERROR: %s", err.Error())
		return
	}

	if err := s.Run(); err != nil {
		fmt.Printf("RUN ERROR: %s", err.Error())
		return
	}
}
