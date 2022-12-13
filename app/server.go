package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DbConfig struct {
	dbHost string
	dbPort string
	dbUser string
	dbPass string
	dbName string
}

func (server *Server) Initialize(appconfig AppConfig) {
	fmt.Println("Welcome to toko ", appconfig.AppName)

	server.initializeRoute()

}

func (server *Server) InitializeDB(dbConfig DbConfig) {
	var err error

	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.dbUser, dbConfig.dbPass, dbConfig.dbHost, dbConfig.dbPort, dbConfig.dbName)

	server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	} else {
		fmt.Println("DB COnnected")
	}

}

func (server *Server) Run(addr string) {
	fmt.Println(".. Listenin on port ,", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {

	var server = Server{}
	var appConfig = AppConfig{}

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error on loading .env")
	}

	appConfig.AppName = getEnv("APP_NAME", "Toko Online Merdeka") // "TOko Online Go"
	appConfig.AppEnv = getEnv("APP_ENV", "development")           // "development"
	appConfig.AppPort = getEnv("APP_PORT", "9000")                // "9001"

	var dbConfig = DbConfig{}
	dbConfig.dbHost = getEnv("DB_HOST", "localhost")
	dbConfig.dbPort = getEnv("DB_PORT", "3306")
	dbConfig.dbUser = getEnv("DB_USER", "root")
	dbConfig.dbPass = getEnv("DB_PASS", "rahasia")
	dbConfig.dbName = getEnv("DB_NAME", "gormdb")

	server.Initialize(appConfig)
	server.InitializeDB(dbConfig)
	server.Run(":" + appConfig.AppPort)
}
