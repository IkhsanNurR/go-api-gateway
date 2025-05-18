package authdbconfig

import (
	"os"
)

var DbHost = ""
var DbPort = ""
var DbUser = ""
var DbPassword = ""
var DbName = ""
var DbSchema = ""

func InitAuthDbConfig() {
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
	DbSchema = os.Getenv("DB_SCHEMA")
}
