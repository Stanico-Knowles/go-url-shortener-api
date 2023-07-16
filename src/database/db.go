package database

import (
	"fmt"
	"go-url-shortener-api/src/database/migrations"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	var datetimePrecision = 2
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_SCHEMA"),
		),
		DefaultStringSize:         256,                // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,               // disable datetime precision support, which not supported before MySQL 5.6
		DefaultDatetimePrecision:  &datetimePrecision, // default datetime precision
		DontSupportRenameIndex:    true,               // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,               // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,              // smart configure based on used version
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database Connected.")

	migrations.RunMigration(db)

	return db
}

func Disconnect(db *gorm.DB) {
	dbConnection, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	dbConnection.Close()
}
