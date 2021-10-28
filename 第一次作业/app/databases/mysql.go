package databases

import (
	"database/sql"
	"github.com/pkg/errors"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Mysql_Init() error  {
	db, err := sql.Open("mysql", "root:Lixun886521.@tcp(127.0.0.1:3306)/book?charset=utf8")
	Db = db
	return errors.Wrap(err,"mysql.Init error")
}
