package main

import (
	"github.com/admin-srf/go_crud/config"
	"github.com/admin-srf/go_crud/router"
)

// @title           Auth Example API
// @version         1.0
// @description     This is a sample server celler server.

// @contact.name   Javier Salvador
// @contact.url    http://www.jsrf.com.br
// @contact.email  javier@jsrf.com.br

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  JWT

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://localhost:8080/swagger/index.html
func main() {
	err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	router.Initialize()

}
