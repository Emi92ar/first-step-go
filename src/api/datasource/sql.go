package datasource

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type DBConfig struct {
	Host string
	User string
	Password string
	DBName string
	PoolConfig *poolConfig
	EnableLogs bool
}

type poolConfig struct {
	MaxIdleConnections int
	MaxOpenConnections int
	ConnMaxLifetime time.Duration
}

func GetDBConnection() *sql.DB {
	config := getDBConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True",
		config.User, config.Password, config.Host, config.DBName)
	connection, err := sql.Open("mysql", dsn)
	if err != nil {
		println("ERROR OPENING CONNECTION TO DB")
		panic(errors.New("error opening connection"))
	}

	err = connection.Ping()
	if err != nil {
		println("Error connection to db")
		panic(errors.New("error opening connection"))
	}
	return connection
}

func getDBConfig() *DBConfig {

	return &DBConfig{
		Host: "127.0.0.1:3306",
		User: "root",
		Password: "DataBase2017",
		DBName: "estate_agency",
		PoolConfig: &poolConfig{
			MaxIdleConnections: 100,
			MaxOpenConnections: 100,
			ConnMaxLifetime: 50,
		},
		EnableLogs: false,
	}
}