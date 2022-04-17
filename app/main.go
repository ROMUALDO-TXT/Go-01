package main

import (
	"github.com/ROMUALDO-TXT/Go-01/app/infrastructure"
)

func main() {
	logger := infrastructure.NewLogger()

	infrastructure.Load(logger)

	infrastructure.HandleRequests(logger)
}
