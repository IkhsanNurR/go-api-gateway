package appconfig

import (
	"log"
	"os"
)

var AppPort = ":8080"
var AppEnv = "dev"
var APIKey = "1234"
var APIKeyMobile = "1234"
var MaxRefreshToken = "3"
var MaxLoginAttempt = "3"
var LoginAttempt = "5" // minutes
var JwtSecretKey = ""
var JwtSecretKeyMobile = ""
var JwtExpiredTime = "8"                    // hours
var JwtExpiredTimeModile = "8"              // hours
var JwtRefreshTokenExpiredTime = "24"       // hours
var JwtRefreshTokenExpiredTimeMobile = "24" // hours
var EndPoint = ""
var AccessKey = ""
var SecretKey = ""
var UseSsl = false
var Bucket = ""

func InitAppConfig() {

	if os.Getenv("APP_PORT") == "" {
		AppPort = ":8080"
		log.Default().Println("AppPort is not set. Defaulting to :8080")
	} else {
		AppPort = os.Getenv("APP_PORT")
	}

	if os.Getenv("APP_ENV") == "" {
		AppEnv = "dev"
		log.Default().Println("AppEnv is not set. Defaulting to development")
	} else {
		AppEnv = os.Getenv("APP_ENV")
	}

	if os.Getenv("API_KEY") == "" {
		APIKey = "1234"
		log.Default().Println("APIKey is not set. Defaulting to 1234")
	} else {
		APIKey = os.Getenv("API_KEY")
	}

	if os.Getenv("API_KEY_MOBILE") == "" {
		APIKeyMobile = "1234"
		log.Default().Println("API_KEY_MOBILE is not set. Defaulting to 1234")
	} else {
		APIKeyMobile = os.Getenv("API_KEY_MOBILE")
	}
	MaxRefreshToken = os.Getenv("MaxRefreshToken")
	MaxLoginAttempt = os.Getenv("MaxLoginAttempt")
	LoginAttempt = os.Getenv("LoginAttempt")
	JwtSecretKey = os.Getenv("JwtSecretKey")
	JwtSecretKeyMobile = os.Getenv("JwtSecretKeyMobile")
	JwtExpiredTime = os.Getenv("JwtExpiredTime")
	JwtExpiredTimeModile = os.Getenv("JwtExpiredTimeModile")
	JwtRefreshTokenExpiredTime = os.Getenv("JwtRefreshTokenExpiredTime")
	JwtRefreshTokenExpiredTimeMobile = os.Getenv("JwtRefreshTokenExpiredTimeMobile")
	EndPoint = os.Getenv("MINIO_ENDPOINT")
	AccessKey = os.Getenv("MINIO_ACCESS_KEY")
	SecretKey = os.Getenv("MINIO_SECRETKEY")
	UseSsl = os.Getenv("MINIO_USE_SSL") == "true"
	Bucket = os.Getenv("MINIO_BUCKET")
}
