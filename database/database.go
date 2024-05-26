package database

import (
	"fmt"
	"log"
	"os"

	// "github.com/fsmardani/go-for-example/models"
	"github.com/spf13/viper"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
	// "gorm.io/gorm/logger"
	"github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql"
)

type Dbinstance struct {
	Db *sqlx.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.Get("DB_USER"),
		viper.Get("DB_PASSWORD"),
		viper.Get("DB_HOST"),
		viper.Get("DB_PORT"),
		viper.Get("DB_NAME"),
	)

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })
    db, err := sqlx.Connect("mysql", dsn)

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	// log.Println("connected")
	// db.Logger = logger.Default.LogMode(logger.Info)

	// log.Println("running migrations")
	// db.AutoMigrate(&models.User{})

	DB = Dbinstance{
		Db: db,
	}
}
