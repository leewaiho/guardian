package dao

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"time"
)

const (
	_driverName = "mysql"
)

type Dao struct {
	DB *sql.DB
}

func New() *Dao {
	return &Dao{
		DB: newMYSQL(),
	}
}

func newMYSQL() *sql.DB {
	mysqlCfg := mysql.Config{
		User:      viper.GetString("db.username"),
		Passwd:    viper.GetString("db.password"),
		Net:       "tcp",
		Addr:      fmt.Sprintf("%s:%d", viper.GetString("db.host"), viper.GetInt("db.port")),
		DBName:    viper.GetString("db.database"),
		ParseTime: true,
		Loc:       time.Local,
	}
	dsn := mysqlCfg.FormatDSN()
	db, err := sql.Open(_driverName, dsn)
	if err != nil {
		panic(err)
	}
	if e := db.Ping(); e != nil {
		panic(e)
	}
	return db
}
