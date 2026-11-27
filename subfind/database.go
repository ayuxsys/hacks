package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var databaseConn *Connection

type Connection struct {
	*gorm.DB
}

// model
type SubFinder struct {
	ID     int
	Domain string   `json:"domain"`
	Subs   []string `gorm:"serializer:json" json:"subs"`
	Tag    string   `json:"tag"`
}

func initDBConnection() {
	if databaseURL == "" {
		if err := godotenv.Load(envPath); err != nil {
			errL.Fatal(err)
		}
		databaseURL = os.Getenv("DB_URL")
	}

	if err := Connect(databaseURL); err != nil {
		errL.Fatalf("error while connecting to the psql database: %v", err)
	}

}

func Connect(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Warn)})
	if err != nil {
		return fmt.Errorf("'gorm.Open': %v", err)
	}
	databaseConn = &Connection{DB: db}
	return databaseConn.Migrate()
}

func (c *Connection) JsonOutput(path string) error {
	var results []SubFinder
	if err := c.Select("domain", "subs", "tag").Find(&results).Error; err != nil {
		return err
	}
	data := must(json.Marshal(results))
	file := must(os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644))
	defer file.Close()
	_, err := file.Write(data)
	return err
}

func (SubFinder) TableName() string {
	return "subfinder"
}

func (c *Connection) Migrate() error {
	return c.AutoMigrate(&SubFinder{})
}
