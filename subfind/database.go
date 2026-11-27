package main

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Conn = &Connection{}
)

type Connection struct {
	*gorm.DB
}

func init() {
	var err error
	Conn, err = Connect()
	if err != nil {
		errL.Printf("failed to connect to psql database saving will be skipped: %v", err)
	}
	if err := Conn.Migrate(); err != nil {
		errL.Fatal(err)
	}
}

func Connect() (*Connection, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}
	conn, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	return &Connection{DB: conn}, err
}

// model
type SubFinder struct {
	ID     int
	Domain string
	Subs   []string `gorm:"serializer:json"`
}

func (SubFinder) TableName() string {
	return "subfinder"
}

func (c *Connection) Migrate() error {
	return c.AutoMigrate(&SubFinder{})
}
