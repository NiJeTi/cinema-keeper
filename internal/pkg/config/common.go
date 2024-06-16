package config

import (
	"database/sql"
	"log"

	"github.com/spf13/viper"
)

func Config() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./cmd/service/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("failed to read config file", err)
	}

	viper.SetConfigName("config.private.yaml")
	_ = viper.MergeInConfig()
}

func DB(cs string) *sql.DB {
	db, err := sql.Open("postgres", cs)
	if err != nil {
		log.Fatalln("failed to open database connection", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("failed to ping database", err)
	}

	return db
}
