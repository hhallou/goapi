package lambda

import (
	"log"

	"github.com/ProovGroup/claim/internal/config"
	"github.com/ProovGroup/claim/internal/logger"
	"github.com/ProovGroup/claim/pkg/cmd/lambda"
)

func main() {
	logger.InitLogger()
	config.InitConfig()

	err := lambda.Serve()
	if err != nil {
		log.Fatal("Claim lambda server", err)
	}
}
