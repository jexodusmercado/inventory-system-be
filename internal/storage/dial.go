package storage

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Dial(vi *viper.Viper) (*gorm.DB) {
	DSN:= vi.GetString("DB_URL")

	fmt.Println(DSN)
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatalln("database connection failed")
		return nil
	}

	return db
}