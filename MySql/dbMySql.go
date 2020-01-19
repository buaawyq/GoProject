package MySql

import (
	"GoProject/Domain"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	USERNAME = "root"
	PASSWORD = "wang19951014"
	NETWORK  = "tcp"
	//SERVER   = "10.135.69.100"
	SERVER = "127.0.0.1"
	PORT   = 3306
	//DATABASE = "test"
	DATABASE = "satellite"
)

func InitDataBase() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		print(err)
		return nil
	}
	DB.SetConnMaxLifetime(100 * time.Second)
	DB.SetMaxIdleConns(100)
	DB.SetMaxOpenConns(16)
	return DB
}

func QueryOne(DB *sql.DB, user string) *Domain.Card {
	card := new(Domain.Card)
	row := DB.QueryRow("select * from card1 where password=?", user)
	if err := row.Scan(&card.Id, &card.Number, &card.User, &card.Password); err != nil {
		fmt.Printf("scan fa iled, err:%v", err)
		return nil
	}
	return card
}
