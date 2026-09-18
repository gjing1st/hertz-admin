# CORS Hertz's middleware

Hertz middleware/handler to enable CORS support.

This repo forks from [gin cors](https://github.com/gin-contrib/cors) and adapt it to Hertz.

## Usage

### Start using it

Download and install it:

```sh
go get github.com/hertz-contrib/cors
```

Import it in your code:

```go
import "github.com/hertz-contrib/cors"
```

### Canonical example

```go
package main

import (
	"context"
	"time"

	"github.com/hertz-contrib/cors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	h.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"}, // Allowed domains, need to bring schema
		AllowMethods:     []string{"PUT", "PATCH"},    // Allowed request methods
		AllowHeaders:     []string{"Origin"},          // Allowed request headers
		ExposeHeaders:    []string{"Content-Length"},  // Request headers allowed in the upload_file
		AllowCredentials: true,                        // Whether cookies are attached
		AllowOriginFunc: func(origin string) bool { // Custom domain detection with lower priority than AllowOrigins
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour, // Maximum length of upload_file-side cache preflash requests (seconds)
	}))

	h.GET("/cors", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "Hello hertz!")
	})

	h.Spin()
}
```

### Using DefaultConfig as start point

```go
func main() {
  h := server.Default()
  // - No origin allowed by default
  // - GET,POST, PUT, HEAD methods
  // - Credentials share disabled
  // - Preflight requests cached for 12 hours
  config := cors.DefaultConfig()
  config.AllowOrigins = []string{"http://google.com"}
  // config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
  // config.AllowAllOrigins = true

  h.Use(cors.New(config))
  h.Spin()
}
```
note: while Default() allows all origins, DefaultConfig() does not and you will still have to use AllowAllOrigins

### Default() allows all origins

```go
func main() {
  h := server.Default()
  // same as
  // config := cors.DefaultConfig()
  // config.AllowAllOrigins = true
  // h.Use(cors.New(config))
  h.Use(cors.Default())
  h.Spin()
}
```
