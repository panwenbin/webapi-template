package mysql

import (
	"app/settings"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

var (
	DbConnection string
	DbUsername   string
	DbPassword   string
	DbHost       string
	DbPort       string
	DbDatabase   string
)

func init() {
	settings.RequireEnvs([]string{
		"DB_CONNECTION", "DB_HOST", "DB_PORT", "DB_DATABASE", "DB_USERNAME", "DB_PASSWORD",
	})

	DbConnection = os.Getenv("DB_CONNECTION")
	DbUsername = os.Getenv("DB_USERNAME")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbDatabase = os.Getenv("DB_DATABASE")

	var err error

	connArgs := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DbUsername,
		DbPassword,
		DbHost,
		DbPort,
		DbDatabase,
	)

	Db, err = gorm.Open(DbConnection, connArgs)

	if err != nil {
		log.Fatalln(err)
	}

	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetMaxIdleConns(20)
	Db.DB().SetConnMaxLifetime(55 * time.Second)

	if settings.Debug {
		Db = Db.Debug()
	}
	Db.AutoMigrate()
}
