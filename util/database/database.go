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
		logg.Printlogger("\t\t Database connection FAILED :: Connection GORM ", err)
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

// Colors
const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

func (sqlLog *SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sqlStatement, rowsAffected := fc()
	logg.Printlogger(fmt.Sprintf("%v***** SQL Statement ***** | ", Green), fmt.Sprintf("%v[ Row Affected : %v%v%v ] -> %v%v \n", BlueBold, YellowBold, rowsAffected, BlueBold, YellowBold, sqlStatement))
}
