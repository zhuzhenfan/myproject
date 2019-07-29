package room

import (
	"github.com/go-xorm/xorm"
	"wolfkill/wolfkill/common"
	"wolfkill/wolfkill/common/table"
	"wolfkill/wolfkill/module/room"
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
	return table.InitTable(table.TbRoom, new(room.Room))
}