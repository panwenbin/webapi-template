package clickhouse

import (
	"app/settings"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ClickHouse/clickhouse-go"
)

var ClickHouse *sql.DB

var (
	ClickHouseUri string
)

func init() {
	settings.RequireEnvs([]string{
		"CLICKHOUSE_URI",
	})

	ClickHouseUri = os.Getenv("CLICKHOUSE_URI")
	var err error
	ClickHouse, err = sql.Open("clickhouse", ClickHouseUri)
	if err != nil {
		log.Fatal(err)
	}

	if err := ClickHouse.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return
	}
}
