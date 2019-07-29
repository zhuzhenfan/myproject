package gameinfo

import (
	"github.com/go-xorm/xorm"
	"wolfkill/wolfkill/common"
	"wolfkill/wolfkill/common/table"
	"wolfkill/wolfkill/module/assemble/game/gameinfo"
)

type BeanConnect struct {
	Engine  *xorm.Engine
	Session *xorm.Session
}

func GetEngine() *xorm.Engine {
	return common.PGClient
}
func GetSession() *xorm.Session {
	return common.PGClient.NewSession()
}
func Init() error {
	err := table.InitTable(table.TbGameRecord, new(gameinfo.GameRecord))
	if err != nil{
		return err
	}
	err = table.InitTable(table.TbGameEvent, new(gameinfo.GameEvent))
	if err != nil{
		return err
	}
	err = table.InitTable(table.TbGameWho, new(gameinfo.GameWho))

	return nil
}