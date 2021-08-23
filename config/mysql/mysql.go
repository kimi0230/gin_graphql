package mysql

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var GormDB *gorm.DB

func init() {
	if len(os.Args) > 1 {
		env := os.Args[1]
		switch env {
		case "app":
			// fmt.Println("----- run app env -----")
			godotenv.Load(".env")
		case "dev":
			// fmt.Println("----- run develop env -----")
			godotenv.Load("./.env.dev")
		case "qa":
			// fmt.Println("----- run qa env -----")
			godotenv.Load("./.env.qa")
		default:
			// fmt.Println("----- run default env (dev) -----")
			godotenv.Load("./.env.dev")
		}
	} else {
		// fmt.Println("----- run default env (dev)-----")
		godotenv.Load("./.env.dev")
	}

	fmt.Printf("===== DB: %s ===== \n", os.Getenv("DB_DATABASE"))

	connection := os.Getenv("DB_CONNECTION")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	// 使用go-sql-driver 来連接mysql，獲取的時區默認是UTC +0的
	url := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true"

	var err error

	// 使用ORM套件: gorm 連線 DB
	GormDB, err = gorm.Open(connection, url)
	if err != nil {
		log.Error(err.Error())
	}
	err = GormDB.DB().Ping()
	if err != nil {
		log.Error(err.Error())
	}

	if err != nil {
		log.Error(err.Error())
	}
}
