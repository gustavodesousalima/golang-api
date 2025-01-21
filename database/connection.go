package database

import (
	"database/sql"
	"fmt"
	"os"

	"golang-api/config"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	config.LoadConfig()

	dsn := os.Getenv("DATABASE")

    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        return fmt.Errorf("falha ao conectar ao banco: %v", err)
    }

    if err = DB.Ping(); err != nil {
        return fmt.Errorf("falha ao pingar banco: %v", err)
    }

    fmt.Println("Conexão com banco estabelecida")
    return nil
}