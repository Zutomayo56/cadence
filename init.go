package main

import (
	"database/sql"
	"fmt"

	env "github.com/deanishe/go-env"
	"github.com/kenellorando/clog"
)

// Config - Primary configuration object holder
type Config struct {
	server CConfig
	db     DBConfig
}

// CConfig - CServer configuration
type CConfig struct {
	Port     string
	MusicDir string
}

// DBConfig - Database configuration
type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

// Grabs all environment variable configurations
func getConfigs() Config {
	clog.Info("getConfig", "Getting configuration")
	// Full config object
	config := Config{}

	// Get server-related configs
	server := CConfig{}
	server.Port = env.GetString("CSERVER_WEB_PORT", ":8000")
	server.MusicDir = env.GetString("CSERVER_MUSIC_DIR", "/Default/Fake/Music/Dir")
	config.server = server

	// Get database related configs
	db := DBConfig{}
	db.Host = env.GetString("CSERVER_DB_HOST", "Default_FakeDBHost")
	db.Port = env.GetString("CSERVER_DB_PORT", "Default_FakeDBPort")
	db.User = env.GetString("CSERVER_DB_USER", "Default_FakeDBUser")
	db.Pass = env.GetString("CSERVER_DB_PASS", "Default_FakeDBPass")
	db.Name = env.GetString("CSERVER_DB_NAME", "Default_FakeDBName")
	config.db = db

	return config
}

// Establishes database connection using configuration
func connectDatabase(db DBConfig) (*sql.DB, error) {
	clog.Debug("connectDatabase", "Attempting connection to database...")

	// Form a connection with the database using config
	connectInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", db.Host, db.Port, db.User, db.Pass, db.Name)
	database, err := sql.Open("postgres", connectInfo)
	if err != nil {
		clog.Error("connectDatabase", "Connection to the database failed!", err)
	} else {
		clog.Info("connectDatabase", "Connected to the database.")
	}

	return database, err
}
