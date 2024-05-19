package main

import (
	"crypto/tls"
	"net"
	"fmt"
	// "io/ioutil"
	// "log"
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
	// cert, err := ioutil.ReadFile("certs/local.crt")
	// if err != nil {
	//   fmt.Println("Failed to read certificate file: %v", err)
	// }
	// fmt.Println(cert)
	// app.Listen(":3000") 

	// app.ListenTLS("127.0.0.1:80", "certs/local.crt", "certs/local.key")

	ln, _ := net.Listen("tcp", "127.0.0.1:80")

	cer, _:= tls.LoadX509KeyPair("certs/local.crt", "certs/local.key")

	ln = tls.NewListener(ln, &tls.Config{Certificates: []tls.Certificate{cer}})

	app.Listener(ln)

}
