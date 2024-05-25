package main

import (
	// "crypto/tls"
	// "os"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	// "github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	// "github.com/rs/zerolog"


	"github.com/fsmardani/go-for-example/database"
	"github.com/fsmardani/go-for-example/models"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func main() {


	//viper config
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return
	}
	var config models.Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		return
	}
	// viper.AutomaticEnv()

	fmt.Println(viper.Get("TEST_ENV"))

	//DB connection
	database.ConnectDb()

	app := fiber.New()

	//prometheus config
	recordMetrics()

	prometheus.Register(opsProcessed)
	metricsHandler := promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{})
	app.Get("/metrics", adaptor.HTTPHandler(metricsHandler))

	//set routers
	setupRoutes(app)


	
	// app.Listen(":3000")

	app.ListenTLS(":443", "certs/cert.pem", "certs/key.pem")

	
	// cer, _:= tls.LoadX509KeyPair("certs/cert.crt", "certs/key.key")

	// ln, _ := tls.Listen("tcp","127.0.0.1:443", &tls.Config{Certificates: []tls.Certificate{cer}})

	// app.Listener(ln)

}
