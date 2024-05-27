package main

import (
	// "crypto/tls"
	// "os"
	// "fmt"
	// "net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	// "github.com/rs/zerolog"

	// "github.com/fsmardani/go-for-example/database"
	"github.com/fsmardani/go-for-example/config"
	"github.com/fsmardani/go-for-example/database"
)

func init() {
	config.GetViperConfig()

	//DB connection
	// database.ConnectDb()
	recordMetrics() 
}



func main() {
	
	database.Connect()

	app := fiber.New()

	app.Get("/metrics", adaptor.HTTPHandler(MetricsHandler))

	// http.Handle("/", adaptor.FiberHandler(greet))

	setupRoutes(app)

	go createThumbnail()


	// fmt.Println(config.ConfigValues.GetString("NAME"))
	
	app.Listen(":3000")
	// http.ListenAndServe(":3000", nil)

	// app.ListenTLS(":443", "certs/cert.pem", "certs/key.pem")

	
	// cer, _:= tls.LoadX509KeyPair("certs/cert.crt", "certs/key.key")

	// ln, _ := tls.Listen("tcp","127.0.0.1:443", &tls.Config{Certificates: []tls.Certificate{cer}})

	// app.Listener(ln)

}

// func greet(c *fiber.Ctx) error {
// 	return c.SendString("Hello World!")
// }