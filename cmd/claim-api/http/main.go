package main

import (
	"log"

	"github.com/ProovGroup/claim/internal/config"
	"github.com/ProovGroup/claim/internal/logger"
	"github.com/ProovGroup/claim/pkg/cmd/router"
)

func main() {
	logger.InitLogger()
	config.InitConfig()

	err := router.Serve()
	if err != nil {
		log.Fatal("Claim API server", err)
	}
}
