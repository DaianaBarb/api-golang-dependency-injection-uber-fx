package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {

	netParam := os.Getenv("tcp")
	if netParam == "" {
		netParam = "tcp"
	}

	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER_TAG"),
		Passwd:               os.Getenv("DB_PASS_TAG"),
		Net:                  netParam,
		Addr:                 os.Getenv("DB_ADDR_TAG"),
		DBName:               os.Getenv("DB_NAME_TAG"),
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Sprintf("mysql database open error: %v", err)
	}

	//defer db.Close()
	db.SetConnMaxLifetime(time.Minute * 4)
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Sprintf("mysql database ping error: %v", pingErr)
	}

	fmt.Println("MySQL connected!")
	return db
}

func Close(db *sql.DB) {
	db.Close()
}
