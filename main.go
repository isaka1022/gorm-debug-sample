package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	gorm.Model
	ID    int `gorm:"AUTO_INCREMENT"`
	Code  string
	Price uint
}

type dbConfig struct {
	Username string
	Userpass string
	Host     string
	Port     string
	Name     string
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	var config dbConfig
	err = envconfig.Process("db", &config)

	if err != nil {
		log.Fatal(err.Error())
	}

	dsn := config.Username + ":" + config.Userpass + "@(" + config.Host + ":" + config.Port + ")/" + config.Name + "?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Logger = db.Logger.LogMode(logger.Info)
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})
	db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	db.First(&product)

	db.Model(&product).Update("Price", 200)
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})

	db.Delete(&product, product.ID)
}
