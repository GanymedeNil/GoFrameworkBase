package main

import (
	"github.com/GanymedeNil/GoFrameworkBase/internal/core"
)

// @host      localhost:8080
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	core.Viper()
	core.Zap()
	core.Gorm()
	core.Service()
}
