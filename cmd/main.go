package main

import (
	// "crypto/tls"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/fsmardani/go-for-example/database"
	"github.com/fsmardani/go-for-example/models"

)


func main() {

	viper.SetConfigType("yaml")
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
	if err := viper.ReadInConfig();  err != nil {
        return
    }
    var config models.Config
    if err := viper.Unmarshal(&config); err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(viper.Get("DB_USER"))
	database.ConnectDb()

	app := fiber.New()
	prometheus := fiberprometheus.New("go-for-example")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	setupRoutes(app)

	// app.Listen(":3000") 

	app.ListenTLS(":443", "certs/certificate.pem", "certs/key.pem")


	// cer, _:= tls.LoadX509KeyPair("certs/certificate.pem", "certs/key.pem")

	// ln, _ := tls.Listen("tcp","127.0.0.1:443", &tls.Config{Certificates: []tls.Certificate{cer}})

	// app.Listener(ln)



}
