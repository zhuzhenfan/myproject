package game

import (
	"wolfkill/wolfkill/common/table"
	"wolfkill/wolfkill/module/assemble/game"
	"wolfkill/wolfkill/module/assemble/game/gameinfo"
	"wolfkill/wolfkill/module/room"
)
// create room
func (bc *BeanConnect)InsertRoom(bean ...*room.Room)(int64,error){
	return bc.Session.Table(table.TbRoom).Insert(&bean)
}

// connect room and role
func (bc *BeanConnect)InsertGame(bean *[]game.Game)(int64,error){
	return bc.Session.Table(table.TbGame).Insert(bean)
}

//
func (bc *BeanConnect)UpdateGameByRoomId(bean *game.Game, field []string)(int64,error){
	//cond:=builder.Cond(builder.And(builder.Eq{table.TbGame+"."+game.RoomId:bean.RoomId}))
	//cond = cond.And(builder.Eq{table.TbGame+"."+game.RoleId:bean.RoleId})
	//cond = cond.And(builder.Eq{table.TbGame+"."+game.Select:false})
	//Where(cond)

	for _,val:=range field{
		bc.Session = bc.Session.Cols(val)
	}
	// 非主键不能用id
	return bc.Session.Table(table.TbGame).
		In(table.TbGame+"."+game.RoomId,bean.RoomId).Update(bean)
}

func (bc *BeanConnect)UpdateGameByGameId(bean *game.Game, field []string)(int64,error){
	for _,val:=range field{
		bc.Session = bc.Session.Cols(val)
	}
	return bc.Session.Table(table.TbGame).ID(bean.GameId).Update(bean)
}

func (bc *BeanConnect)UpdateRoomById(bean *room.Room)(int64,error){
	return bc.Session.Table(table.TbRoom).ID(bean.RoomId).Update(bean)
}

func (bc *BeanConnect)InsertGameRecord(bean *gameinfo.GameRecord)(int64,error){
	return bc.Session.Table(table.TbGameRecord).Insert(bean)
}

func (bc *BeanConnect)DeleteGameRecord(condi map[string]interface{})(int64,error){
	for key,val:=range condi{
		bc.Session = bc.Session.In(table.TbGameRecord+"."+key,val)
	}
	bean:=gameinfo.GameRecord{}
	return bc.Session.Table(table.TbGameRecord).Delete(&bean)
}

func (bc *BeanConnect)DeleteGameEvent(condi map[string]interface{})(int64,error){
	for key,val:=range condi{
		bc.Session = bc.Session.In(table.TbGameEvent+"."+key,val)
	}
	bean:=gameinfo.GameEvent{}
	return bc.Session.Table(table.TbGameEvent).Delete(&bean)
}

func (bc *BeanConnect)DeleteGameWho(condi map[string]interface{})(int64,error){
	for key,val:=range condi{
		bc.Session = bc.Session.In(table.TbGameWho+"."+key,val)
	}
	bean:=gameinfo.GameWho{}
	return bc.Session.Table(table.TbGameWho).Delete(&bean)
}

func (bc *BeanConnect)DeleteRoom(condi map[string]interface{})(int64,error){
	for key,val:=range condi{
		bc.Session = bc.Session.In(table.TbRoom+"."+key,val)
	}
	bean:=room.Room{}
	return bc.Session.Table(table.TbRoom).Delete(&bean)
}

func (bc *BeanConnect)DeleteGame(condi map[string]interface{})(int64,error){
	for key,val:=range condi{
		bc.Session = bc.Session.In(table.TbGame+"."+key,val)
	}
	bean:=game.Game{}
	return bc.Session.Table(table.TbGame).Delete(&bean)
}
