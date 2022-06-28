package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/newrelic/go-agent/_integrations/nrmysql"
)

func GetDB(batch bool) (*sqlx.DB, error) {
	mysqlConfig := mysql.NewConfig()
	mysqlConfig.Net = "tcp"
	mysqlConfig.Addr = GetEnv("MYSQL_HOSTNAME", "127.0.0.1") + ":" + GetEnv("MYSQL_PORT", "3306")
	mysqlConfig.User = GetEnv("MYSQL_USER", "isucon")
	mysqlConfig.Passwd = GetEnv("MYSQL_PASS", "isucon")
	mysqlConfig.DBName = GetEnv("MYSQL_DATABASE", "isucholar")
	mysqlConfig.Params = map[string]string{
		"time_zone": "'+00:00'",
	}
	mysqlConfig.ParseTime = true
	mysqlConfig.MultiStatements = batch

	db, err := sql.Open("nrmysql", mysqlConfig.FormatDSN())
	if err != nil {
		log.Fatalf("failed to connect to DB: %s.", err.Error())
	}
	// defer db.Close()
	return sqlx.NewDb(db, "nrmysql"), nil

	// return sqlx.Open("mysql", mysqlConfig.FormatDSN())
}
