package main

import (
	"context"
	"github.com/gowool/swagger"
	_ "github.com/gowool/swagger/example/docs"
	"github.com/gowool/wool"
	"os"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /api/v1
func main() {
	w := wool.New()
	swg := swagger.New(&swagger.Config{})
	w.GET("/swagger/...", swg.Handler)

	srv := wool.NewServer(&wool.ServerConfig{
		Address: ":1323",
	})
	if err := srv.StartC(context.Background(), w); err != nil {
		wool.Logger().Error("server error", err)
		os.Exit(1)
	}
}
