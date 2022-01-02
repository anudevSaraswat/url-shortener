package main

import (
	_ "github.com/go-sql-driver/mysql"
	"goprojects/urlshortener/config/db"
	"goprojects/urlshortener/config/env"
	"goprojects/urlshortener/route"
)

func init() {
	db.InitDb()
}

func main() {

	echo := route.InitServer()
	echo.Logger.Fatal(echo.Start(env.ServicePort()))

}
