package mysql

import (
	"database/sql"
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

	// log
	return db
}
