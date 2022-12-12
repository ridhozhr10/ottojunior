package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBHost string
	DBUser string
	DBPass string
	DBPort string
	DBName string
}

// ConnectGORM open connection using golang orm package
func ConnectGORM(c Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		c.DBHost,
		c.DBUser,
		c.DBPass,
		c.DBName,
		c.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
