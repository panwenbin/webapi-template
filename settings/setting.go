package settings

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Debug bool

var EnableMysql bool
var EnableMongo bool

func boolString(s string) bool {
	if s != "" && s != "false" && s != "FALSE" && s != "0" {
		return true
	}

	return false
}

func init() {
	_ = godotenv.Load()

	Debug = boolString(os.Getenv("DEBUG"))
	EnableMysql = boolString(os.Getenv("ENABLE_MYSQL"))
	EnableMongo = boolString(os.Getenv("ENABLE_MONGO"))
}

func RequireEnvs(needEnvs []string) {
	for _, envKey := range needEnvs {
		if os.Getenv(envKey) == "" {
			log.Fatalf("env %s missed", envKey)
		}
	}
}
