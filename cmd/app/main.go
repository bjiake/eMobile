package main

import (
	"eMobile/internal/config"
	"eMobile/internal/di"
	"eMobile/pkg/logger"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	cfg := config.LoadConfig()
	log := logger.New()
	config.SetSwaggerDefaultInfo(cfg)

	if server, err := di.InitializeAPI(cfg, log); err == nil {
		server.Start(cfg, log)
	}
}
