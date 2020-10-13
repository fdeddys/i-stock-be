package database

import (
	"fmt"
	"log"

	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kelseyhightower/envconfig"
)

// Env ..
type Env struct {
	DbUser  string `envconfig:"POSTGRES_USER" default:"resto"`
	DbPass  string `envconfig:"POSTGRES_PASS" default:"Resto#123"`
	DbName  string `envconfig:"POSTGRES_DB_NAME" default:"restodb"`
	DbHost  string `envconfig:"POSTGRES_HOST" default:"156.67.214.228"`
	DbPort  string `envconfig:"POSTGRES_PORT" default:"5432"`
	DbDebug bool   `envconfig:"POSTGRES_DEBUG" default:"true"`
	DbType  string `envconfig:"POSTGRES_TYPE" default:"POSTGRES"`
	SslMode string `envconfig:"POSTGRES_SSL_MODE" default:"disable"`
}

var (
	DbCon *gorm.DB
	DbErr error
	dbEnv Env
)

func init() {

	log.Println("DB POSTGRES")
	err := envconfig.Process("", &dbEnv)
	if err != nil {
		fmt.Println("Failed to get DB env:", err)
	}

	if DbOpen() != nil {
		//panic("DB Can't Open")
		fmt.Println("Can Open database Postgres")
	}
	DbCon = GetDbCon()
	DbCon = DbCon.LogMode(true)

}

// DbOpen ..
func DbOpen() error {
	args := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbEnv.DbHost, dbEnv.DbPort, dbEnv.DbUser, dbEnv.DbPass, dbEnv.DbName, dbEnv.SslMode)
	DbCon, DbErr = gorm.Open("postgres", args)

	if DbErr != nil {
		logs.Error("open database Err ", DbErr)
		return DbErr
	}

	if errping := DbCon.DB().Ping(); errping != nil {
		return errping
	}
	return nil
}

// GetDbCon ..
func GetDbCon() *gorm.DB {
	//TODO looping try connection until timeout
	// using channel timeout
	if errping := DbCon.DB().Ping(); errping != nil {
		logs.Error("Db Not Connect test Ping :", errping)
		errping = nil
		if errping = DbOpen(); errping != nil {
			logs.Error("try to connect again but error :", errping)
		}
	}
	DbCon.LogMode(dbEnv.DbDebug)
	return DbCon
}
