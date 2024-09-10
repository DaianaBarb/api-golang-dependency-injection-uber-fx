package mysql

import (
	"database/sql"
	"log"
	"os"
	"time"

	mysql "github.com/go-sql-driver/mysql"
)

func NewConnectDB() *sql.DB {

	netParam := os.Getenv("DB_NET")
	if netParam == "" {
		netParam = "tcp"
	}

	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  netParam,
		Addr:                 os.Getenv("DB_ADDR"),
		DBName:               os.Getenv("DB_NAME"),
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		// log
	}

	//defer db.Close()
	db.SetConnMaxLifetime(time.Minute * 4)
	pingErr := db.Ping()
	if pingErr != nil {
		// log
	}
	query := `CREATE TABLE IF NOT EXISTS user( user_username VARCHAR(50), user_password VARCHAR(100))`
	_, err = db.Exec(query)
	if err != nil {
		log.Printf("Error %s when creating user table", err)

	}
	query2 := `CREATE TABLE IF NOT EXISTS client_cli(client_name VARCHAR(50) PRIMARY KEY NOT NULL UNIQUE, client_tel VARCHAR(50), client_cpf VARCHAR(50) , client_createdAt DATE, client_active boolean DEFAULT false )`
	_, err = db.Exec(query2)
	if err != nil {
		log.Printf("Error %s when creating client_cli table", err)
	}
	// log
	return db
}
