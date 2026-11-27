package main

import (
	"fmt"

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
	Domain string
	Subs   []string `gorm:"serializer:json"`
	Tag    string
}

func Connect(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Warn)})
	if err != nil {
		return fmt.Errorf("'gorm.Open': %v", err)
	}
	databaseConn = &Connection{DB: db}
	return databaseConn.Migrate()
}

func (SubFinder) TableName() string {
	return "subfinder"
}

func (c *Connection) Migrate() error {
	return c.AutoMigrate(&SubFinder{})
}
