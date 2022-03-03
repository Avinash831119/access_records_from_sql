package util

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

func GetConfiguration() (Configuration, error) {
	config := Configuration{}
	file, err := os.Open("connection.json")
	if err != nil {
		return config, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}
	defer file.Close()
	return config, nil
}

func Connection() (*sql.DB, error) {

	config, err := GetConfiguration()

	if err != nil {
		fmt.Println("dbHandler:Connection, unable to read config file ", err.Error())
		return nil, err
	}

	//port, err := strconv.Atoi(config.Port)

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Server, config.Port, config.Database)

	database, dberr := sql.Open("mysql", connString)

	if dberr != nil {
		fmt.Println("dbHandler:Connection, unable to create database connection ", dberr.Error())
		return nil, dberr
	}
	fmt.Println("dbHandler:Connection, created database connection")
	return database, nil
}
