# Swagger

Wool handler to automatically generate RESTful API documentation with Swagger 2.0

## Usage

### Start using it

1. Add comments to your API source code, [See Declarative Comments Format](https://github.com/swaggo/swag#declarative-comments-format).
2. Download [Swag](https://github.com/swaggo/swag) for Go by using:
```sh
go install github.com/swaggo/swag/cmd/swag@latest
```
3. Run the [Swag](https://github.com/swaggo/swag) in your Go project root folder which contains `main.go` file, [Swag](https://github.com/swaggo/swag) will parse comments and generate required files(`docs` folder and `docs/doc.go`).
```sh
swag init
```
4. Download [swagger](https://github.com/gowool/swagger) by using:
```sh
go get github.com/gowool/swagger
```

And import following in your code:
```go
import "github.com/gowool/swagger" // wool swagger handler
```

### Canonical example:

```go
package main

import (
    "context"
    "github.com/gowool/swagger"
    _ "github.com/gowool/swagger/example/docs"
    "github.com/gowool/wool"
	"golang.org/x/exp/slog"
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
	logger := slog.New(slog.HandlerOptions{Level: slog.LevelDebug}.NewJSONHandler(os.Stdout))
    w := wool.New(logger.WithGroup("wool"))
    swg := swagger.New(&swagger.Config{})
    w.GET("/swagger/...", swg.Handler)
    
    srv := wool.NewServer(&wool.ServerConfig{
        Address: ":1323",
    }, logger.WithGroup("server"))
    if err := srv.StartC(context.Background(), w); err != nil {
		srv.Log.Error("server error", err)
        os.Exit(1)
    }
}
```

5. Run it, and browser to http://localhost:1323/swagger/index.html, you can see Swagger 2.0 Api documents.

## License

Distributed under MIT License, please see license file within the code for more details.

