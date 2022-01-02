package env

import (
	"fmt"
	"os"
)

const (
	kDbPort              = "DB_PORT"
	kDbUser              = "DB_USER"
	kDbPassword          = "DB_PASSWORD"
	kDbHost              = "DB_HOST"
	kDbCharacterEncoding = "DB_CHR_ENC"
	kDbTimeParsing       = "DB_TIME_PARSING"
	kDbShortUrl          = "DB_SHORTURL"
	kServicePort         = "SERVICE_PORT"
	FilePath             = ".\\.env"
)

func ShortURLDBConString() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s", os.Getenv(kDbUser)+
		":", os.Getenv(kDbPassword)+"@tcp(", os.Getenv(kDbHost)+
		":", os.Getenv(kDbPort)+")/", os.Getenv(kDbShortUrl)+
		"?", os.Getenv(kDbCharacterEncoding)+"&",
		os.Getenv(kDbTimeParsing))
}

func ServicePort() string {
	return os.Getenv(kServicePort)
}
