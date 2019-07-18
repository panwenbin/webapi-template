package databases

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

func init() {
	var err error

	connArgs := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	Db, err = gorm.Open(os.Getenv("DB_CONNECTION"), connArgs)

	if err != nil {
		log.Fatalln(err)
	}

	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetMaxIdleConns(20)
	Db.DB().SetConnMaxLifetime(55 * time.Second)

	debug := os.Getenv("DEBUG")
	if debug != "" && debug != "false" && debug != "0" {
		Db = Db.Debug()
	}
}
