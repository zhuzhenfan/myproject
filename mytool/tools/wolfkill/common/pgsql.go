// pgsql
package common

import (
	"errors"
	"fmt"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var PGClient *xorm.Engine

// default database is postgres
const (
	defaultDB = "postgres"
)
type BeanConnect struct {
	Engine  *xorm.Engine
	Session *xorm.Session
}
func GetEngine() *xorm.Engine {
	return PGClient
}
func GetSession() *xorm.Session {
	return PGClient.NewSession()
}

func CreatePGEngine() error {
	// 先连接目标数据
	dbSource := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable",
		dbName, dbUserName, dbPassWord, dbAddr, dbPort)
	pgClient, err := xorm.NewEngine(dbType, dbSource)
	if err == nil {
		_, err := pgClient.DBMetas()
		if err == nil {
			PGClient = pgClient
			return nil
		}
	}
	// 若连接失败，则连接默认的数据库
	if err := pgClient.Close(); err != nil {
		return err
	}
	dbSource = fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable",
		defaultDB, dbUserName, dbPassWord, dbAddr, dbPort)
	pgClient, err = xorm.NewEngine(dbType, dbSource)
	if pgClient == nil {
		return errors.New("sql engine create fail")
	}
	// 之后创建目标数据库
	_, err = pgClient.Exec("CREATE DATABASE " + dbName + ";")
	if err != nil {
		return err
	}
	if err := pgClient.Close(); err != nil {
		return err
	}
	// 再重新连接
	dbSource = fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable",
		dbName, dbUserName, dbPassWord, dbAddr, dbPort)
	pgClient, err = xorm.NewEngine(dbType, dbSource)
	if err != nil {
		return err
	}
	_, err = pgClient.DBMetas()
	if err == nil {
		PGClient = pgClient
		return nil
	}
	return errors.New("sql engine init fail")
}
