package database

import (
	"fmt"
	"lightyear/core/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

type Database interface {
	Dsn() (string, error)
}

type MySqlDB struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}
type PostgreSqlDB struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}
type sqliteDB struct {
	Location string
}

func (db *PostgreSqlDB) Dsn() (string, error) {
	// TODO: validate database config
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", db.Host, db.Port, db.Username, db.Password, db.Database), nil
}

func (db *MySqlDB) Dsn() (string, error) {
	// TODO: validate database config
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=skip-verify", db.Username, db.Password, db.Host, db.Port, db.Database), nil
}

func (db *sqliteDB) Dsn() (string, error) {
	// TODO: validate database config
	return fmt.Sprintf("%s", db.Location), nil
}

func InitDatabase() (err error) {

	var db *gorm.DB
	switch global.Config.Database.Type {
	case "postgres":
		postgresConfig := &PostgreSqlDB{
			Host:     global.Config.Database.Host,
			Port:     global.Config.Database.Port,
			Username: global.Config.Database.Username,
			Password: global.Config.Database.Password,
			Database: global.Config.Database.Database,
		}
		dsn, err := postgresConfig.Dsn()
		if err != nil {
			panic(fmt.Errorf("error: %s", err))
		}
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("connect to database failed: %s", err))
		}

	case "mysql":
		mysqlConfig := &MySqlDB{
			Host:     global.Config.Database.Host,
			Port:     global.Config.Database.Port,
			Username: global.Config.Database.Username,
			Password: global.Config.Database.Password,
			Database: global.Config.Database.Database,
		}
		dsn, err := mysqlConfig.Dsn()
		if err != nil {
			panic(fmt.Errorf("connect to database failed: %s", err))
		}
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("connect to database failed: %s", err))
		}

	case "sqlite":
		sqliteConfig := &sqliteDB{
			Location: global.Config.Database.Location,
		}
		dsn, err := sqliteConfig.Dsn()
		if err != nil {
			panic(fmt.Errorf("connect to database failed: %s", err))
		}
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("connect to database failed: %s", err))
		}

	default:
		panic(fmt.Errorf("invalid database type: %s", global.Config.Database.Type))

	}

	sqlDB, err := db.DB()
	if err != nil {
		global.Logger.Warnln(fmt.Sprintf("get sqlDB failed: %s", err))
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)

	global.DB = db
	global.Logger.Infoln("connect to database successfully")

	return err
}
