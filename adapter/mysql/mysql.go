package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	DB *sql.DB
}

func New(user, password, host, database string, port int) (*Mysql, error) {
	conn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, password, host, port, database)
	println(conn)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}

	return &Mysql{DB: db}, nil
}
