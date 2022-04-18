package main

import (
	"github.com/ROMUALDO-TXT/Go-01/app/infrastructure"
)

func main() {
	logger := infrastructure.NewLogger()
	infrastructure.Load(logger)

	sqlHandler, err := infrastructure.NewSQLHandler()
	if err != nil {
		logger.LogError("%s", err)
	}

	infrastructure.HandleRequests(sqlHandler, logger)
}
