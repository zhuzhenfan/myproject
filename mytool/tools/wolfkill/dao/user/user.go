// user
package user

import (
	"wolfkill/wolfkill/common"
	"wolfkill/wolfkill/common/table"
	"wolfkill/wolfkill/module/user"

	"github.com/go-xorm/xorm"
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
	return table.InitTable(table.TbUser, new(user.User))
}
