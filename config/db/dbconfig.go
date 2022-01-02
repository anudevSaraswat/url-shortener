package db

import (
	"database/sql"
	"github.com/joho/godotenv"
	"goprojects/urlshortener/config/env"
	"sync"
)

var Db *sql.DB
var oncee sync.Once

func InitDb() {

	oncee.Do(func() {

		err := godotenv.Load(env.FilePath)
		if err != nil {
			panic(err)
		}

		d, err := sql.Open("mysql", env.ShortURLDBConString())
		if err != nil {
			panic(err)
		}

		Db = d
		d.SetMaxOpenConns(50)
		d.SetMaxIdleConns(10)

	})

}
