package main

import (
	"github.com/assiljaby/go-ledger-backend/internal/router"
	"github.com/assiljaby/go-ledger-backend/internal/server"
	"github.com/assiljaby/go-ledger-backend/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	logger.StartLogger()

	_ = godotenv.Load()

	e := server.GetInstance()

	router.Route(e)

	server.Start(e)
}
