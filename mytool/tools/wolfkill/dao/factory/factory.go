// factory
package factory

import (
	"wolfkill/wolfkill/dao/assemble/game"
	"wolfkill/wolfkill/dao/assemble/game/gameinfo"
	"wolfkill/wolfkill/dao/role"
	"wolfkill/wolfkill/dao/room"
	"wolfkill/wolfkill/dao/user"
)

func TableFactory() error {
	err := user.Init()
	if err != nil {
		return err
	}
	err = role.Init()
	if err != nil {
		return err
	}
	err = room.Init()
	if err != nil {
		return err
	}
	err = game.Init()
	if err != nil {
		return err
	}
	err = gameinfo.Init()
	return nil
}
