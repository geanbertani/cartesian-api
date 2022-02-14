package postgres

import (
	"fmt"

	"github.com/geanbertani/cartesian-api/libs/json"
	"github.com/jmoiron/sqlx"
)

type dbConfig struct {
	FilePath string
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DataBase string `json:"database"`
	SSLMode  string `json:"ssl_mode"`
}

var connections = map[string]*sqlx.DB{}

func get(filePath string) (db *sqlx.DB, err error) {
	db, found := connections[filePath]
	if found {
		return
	}

	err = fmt.Errorf(`error database not found. File path: %s`, filePath)

	return
}

func connect(config dbConfig) (db *sqlx.DB, err error) {
	db, err = sqlx.Connect("postgres", fmt.Sprintf("user=%s port=%d password=%s host=%s dbname=%s sslmode=%s",
		config.User,
		config.Port,
		config.Password,
		config.Host,
		config.DataBase,
		config.SSLMode,
	))

	if err != nil {
		err = fmt.Errorf(`error connecting to database because of: %s`, err.Error())
		return
	}

	db.SetMaxIdleConns(0)
	db.SetConnMaxLifetime(0)

	connections[config.FilePath] = db
	return
}

// GetByFile Create a database connection through
// the path of a file
func GetByFile(filePath string) (db *sqlx.DB, err error) {
	var config dbConfig

	db, err = get(filePath)
	if err == nil {
		return
	}

	err = json.UnmarshalFile(filePath, &config)
	if err != nil {
		return
	}

	config.FilePath = filePath

	return connect(config)
}

// MustGetByFile Create a database connection through
// the path of a file and generates a panic in case of error
func MustGetByFile(filePath string) (db *sqlx.DB) {
	db, err := GetByFile(filePath)
	if err != nil {
		panic(err)
	}

	return
}
