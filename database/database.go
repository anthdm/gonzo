package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func Connect() (*Repo, error) {
	var (
		user     = os.Getenv("DB_USER")
		dbname   = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
	)

	connectionString := fmt.Sprintf(
		"sslmode=disable user=%s dbname=%s password=%s", user, dbname, password)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return &Repo{db}, nil
}

type Repo struct {
	gorm.DB
}
