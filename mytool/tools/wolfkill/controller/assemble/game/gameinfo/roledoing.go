package gameinfo

import (
	dao "wolfkill/wolfkill/dao/assemble/game/gameinfo"
	"wolfkill/wolfkill/module/assemble/game/gameinfo"
)

func Sheriff()float32{
	return 1.5
}

func Villager()float32{
	return 1
}
// 预言家(sql版)
func Prophet(in gameinfo.GameOpIn)error{
	bc:=dao.BeanConnect{}
	bc.Engine = dao.GetEngine()
	//exist,roleInfo,err:=bc.GetRoleById(in.RoleId)
	//if err !=nil{
	//	return err
	//}
	//if exist == false{
	//	return errs.ErrFunc(errs.Role_NotExist)
	//}
	//event:=gameinfo.GameEvent{}
	//event.GameEventId = utilstring.SimpleUUID()
	//event.RoomId = in.RoomId
	//event.EventTime = gameinfo.NightTime
	//event.Number = in.Number
	//
	//who:=gameinfo.GameWho{}
	//who.GameEventId = event.GameEventId
	//who.Skill =
	for _,val:=range in.WhoAfterIds{
		bc.GetGameByUserId(val)
	}
	return nil
}

func Witch(){}

func Hunter(){}

func Idiot(){}

func Guard(){}

func Bear(){}

func Stalker(){}

func Magician(){}

func Elder(){}

func RustySwordKnight(){}

func BlackMarketBusinessman(){}

func TombGuard(){}

func Wolf(){}

func WhiteWolfKing(){}

func WolfBeauty(){}

func InvisibleWolf(){}

func Gargoyle(){}

func Cupid(){}

func Thief(){}