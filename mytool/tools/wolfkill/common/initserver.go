// initserver
package common

import (
	"gopkg.in/ini.v1"
)

type confServer struct {
	ServerAddr string `ini:"wolfkill_addr"`
	ServerPort string `ini:"wolfkill_port"`
}

type confSql struct {
	DBType     string `ini:"wolfkill_db_type"`
	DBAddr     string `ini:"wolfkill_db_addr"`
	DBPort     string `ini:"wolfkill_db_port"`
	DBName     string `ini:"wolfkill_db_name"`
	DBUserName string `ini:"wolfkill_db_username"`
	DBPassWord string `ini:"wolfkill_db_password"`
}

func confServers(confName string) (*confServer, error) {
	conf, err := ini.Load(confName)
	if err != nil {
		return nil, err
	}
	bean := confServer{}
	err = conf.MapTo(&bean)
	return &bean, err
}

func confSQL(confName string) (*confSql, error) {
	conf, err := ini.Load(confName)
	if err != nil {
		return nil, err
	}
	bean := confSql{}
	err = conf.MapTo(&bean)
	return &bean, err
}

func initServer(confName, linxuConfName,windowsConfName string) (*confServer, error) {
	server, err := confServers(confName)
	if err != nil {
		server, err = confServers(linxuConfName)
		if err != nil{
			server, err = confServers(windowsConfName)
			return server, err
		}
	}
	return server, nil
}

func initSQL(confName, linxuConfName,windowsConfName string) (*confSql, error) {
	sql, err := confSQL(confName)
	if err != nil {
		sql, err = confSQL(linxuConfName)
		if err != nil{
			sql, err = confSQL(windowsConfName)
			return sql, err
		}
	}
	return sql, nil
}
func InitFunc(confName, linxuConfName,windowsConfName string) error {
	server, err := initServer(confName, linxuConfName,windowsConfName)
	if err != nil {
		return err
	}
	sql, err := initSQL(confName, linxuConfName,windowsConfName)
	if err != nil {
		return err
	}
	ServerAddr = server.ServerAddr
	ServerPort = server.ServerPort

	dbType = sql.DBType
	dbAddr = sql.DBAddr
	dbPort = sql.DBPort
	dbName = sql.DBName
	dbUserName = sql.DBUserName
	dbPassWord = sql.DBPassWord

	return nil
}
