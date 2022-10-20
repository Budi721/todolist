package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config is the required properties to use the database.
type Config struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}

func Open(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Name)
	return gorm.Open(mysql.New(mysql.Config{
		DSN:                  dsn,
		DisableWithReturning: true,
	}), &gorm.Config{})
}
