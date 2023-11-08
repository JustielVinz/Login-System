// Package database provides gorm connection
package database

import (
	"fmt"
	"sample/algorithm"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	_ "github.com/go-sql-driver/mysql" // initate go sql driver
)

var (
	// DBConn Current Database connection
	DBConn *gorm.DB

	// Err Database connection error
	Err error
)

// MySQLConnect Connect to a MySQL driver-based database
func MySQLConnect(username, password, host, databaseName string) {
	username = algorithm.DecodeBase64(username)
	password = algorithm.DecodeBase64(password)
	host = algorithm.DecodeBase64(host)
	databaseName = algorithm.DecodeBase64(host)
	if host != "" {
		host = fmt.Sprintf("tcp(%s)", host)
	}

	DBConn, Err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@%s/%s?parseTime=true", username, password, host, databaseName)), &gorm.Config{})
}

// SQLiteConnect ...
func SQLiteConnect(filename string) {
	DBConn, Err = gorm.Open(sqlite.Open(filename), &gorm.Config{})
}

// PostgreSQLConnect Connect to a PostgreSQL database
func PostgreSQLConnect(username, password, host, databaseName, port, sslMode, timeZone string) {
	username = algorithm.DecodeBase64(username)
	password = algorithm.DecodeBase64(password)
	host = algorithm.DecodeBase64(host)
	databaseName = algorithm.DecodeBase64(databaseName)
	port = algorithm.DecodeBase64(port)
	DBConn, Err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, username, password, databaseName, port, sslMode, timeZone)),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "public.", // schema name
				SingularTable: false,
			}})
}
