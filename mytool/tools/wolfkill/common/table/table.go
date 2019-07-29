// table
package table

import "wolfkill/wolfkill/common"

const (
	TbUser = "user"
	TbRole = "role"
	TbRoom = "room"
	TbGame = "game"
	TbGameNumber = "game_number"
	TbGameRecord = "game_record"
	TbGameEvent = "game_event"
	TbGameWho = "game_who"

)

func InitTable(tableName string, tableStruct interface{}) error {
	engine := common.GetEngine()

	has, err := engine.IsTableExist(tableName)
	if err != nil {
		return err
	}
	if !has {
		err = engine.Sync2(tableStruct)
		if err != nil {
			return err
		}
	}
	return nil
}
