package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig(conf *Config) *DBConfig {
	dbConfig := DBConfig{
		Host:     conf.DB.Host,
		Port:     conf.DB.Port,
		User:     conf.DB.Username,
		DBName:   conf.DB.Dbname,
		Password: conf.DB.Password,
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func InitDatabase(conf *Config) *gorm.DB {
	DB, err := gorm.Open(
		"postgres",
		DbURL(
			BuildDBConfig(conf),
		),
	)
	if err != nil {
		log.Printf("database connection failed with error: %v ", err)
		panic("database failed!")
	}

	return DB
}
