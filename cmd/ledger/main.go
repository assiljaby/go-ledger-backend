package main

import (
	"github.com/assiljaby/go-ledger-backend/environments"
	"github.com/assiljaby/go-ledger-backend/internal/router"
	"github.com/assiljaby/go-ledger-backend/internal/server"
	"github.com/assiljaby/go-ledger-backend/pkg/logger"
)

func main() {
	logger.StartLogger()

	environments.LoadEnv()

	e := server.GetInstance()

	router.Route(e)

	server.Start(e)
}
