package database

import (
	"context"
	"fmt"
	"time"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/config"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB

type SqlLogger struct {
	logger.Interface
}

func InitDatabase() bool {

	gorm_db, err := initDatabaseGORM()
	if err != nil {
		logg.Printlogger("\t\t Database connection FAILED :: ", "Connection GORM", err)
		return false
	}
	Conn = gorm_db

	return true

}

func initDatabaseGORM() (db *gorm.DB, err error) {

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v)/%v?parseTime=True&loc=Local",
		config.Env.MYSQL_USERNAME,
		config.Env.MYSQL_PASSWORD,
		config.Env.MYSQL_HOSTNAME,
		config.Env.MYSQL_DB_NAME,
	)
	gorm_db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			Logger: &SqlLogger{},
			// DryRun: true,
		},
	)
	if err != nil {
		return gorm_db, err
	}

	return gorm_db, nil
}

func ConnectionClose() {
	dbInstance, _ := Conn.DB()
	_ = dbInstance.Close()
}

func (sqlLog *SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sqlStatement, _ := fc()
	logg.Printlogger("\t\t\t ***** Generate GORM SQL Statement *****", "SQL Statement", sqlStatement)
}
