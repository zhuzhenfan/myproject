package game

import (
	"wolfkill/wolfkill/common/table"
	modelGame "wolfkill/wolfkill/module/assemble/game"
	modelGameInfo"wolfkill/wolfkill/module/assemble/game/gameinfo"
	modelRole "wolfkill/wolfkill/module/role"
	modelRoom "wolfkill/wolfkill/module/room"
	modelUser "wolfkill/wolfkill/module/user"
)

func (bc *BeanConnect) GetUserForRoom(userId string) (bool, *modelRoom.Room, error) {
	bean := modelRoom.Room{}
	bean.RoomOwner = userId
	exist, err := bc.Engine.Table(table.TbRoom).Get(&bean)
	return exist, &bean, err
}

func (bc *BeanConnect) GetUserForGame(userId string) (bool, *modelGame.Game, error) {
	bean := modelGame.Game{}
	bean.UserId = userId
	exist, err := bc.Engine.Table(table.TbGame).Get(&bean)
	return exist, &bean, err
}

func (bc *BeanConnect) ExistRoomByCode(roomCode string) (bool, error) {
	return bc.Engine.Table(table.TbRoom).
		In(table.TbRoom+"."+modelRoom.RoomCode, roomCode).Exist()
}

func (bc *BeanConnect) FindGameByCondi(
	condi map[string]interface{},page,perpage int) (int64, []*modelGame.Game, error) {
	session := bc.Engine.NewSession()
	defer session.Close()
	for key, val := range condi {
		session = session.In(table.TbGame+"."+key, val)
	}
	bean := make([]*modelGame.Game, 0)
	if perpage ==-1{
		total, err := bc.Engine.Table(table.TbGame).Asc(modelGame.UserNumber).
			Limit(perpage,(page-1)*perpage).FindAndCount(&bean)
		return total, bean, err
	}
	total, err := bc.Engine.Table(table.TbGame).Asc(modelGame.UserNumber).FindAndCount(&bean)
	return total, bean, err
}

func (bc *BeanConnect) GetRoleById(roleId string) (bool, *modelRole.Role, error) {
	bean := modelRole.Role{}
	exist, err := bc.Engine.Table(table.TbRole).
		In(table.TbRole+"."+modelRole.RoleId, roleId).Get(&bean)
	return exist, &bean, err
}

func (bc *BeanConnect) GetRoomByCode(roomCode string) (bool, *modelRoom.Room, error) {
	bean := modelRoom.Room{}
	bean.RoomCode = roomCode
	exist, err := bc.Engine.Table(table.TbRoom).Get(&bean)
	return exist, &bean, err
}

func (bc *BeanConnect) GetRoomById(roomId string) (bool, *modelRoom.Room, error) {
	bean := modelRoom.Room{}
	bean.RoomId = roomId
	exist, err := bc.Engine.Table(table.TbRoom).ID(bean.RoomId).Get(&bean)
	return exist, &bean, err
}

func (bc *BeanConnect) FindGameByRoomId(roomId string,page, perpage int) (int64,[]*modelGame.Game, error) {
	bean := make([]*modelGame.Game, 0)
	if perpage == -1{
		err := bc.Engine.Table(table.TbGame).
			In(table.TbGame+"."+modelGame.RoomId, roomId).Find(&bean)
		return int64(len(bean)),bean, err
	}
	total,err := bc.Engine.Table(table.TbGame).
		In(table.TbGame+"."+modelGame.RoomId, roomId).
		Limit(perpage,(page-1)*perpage).FindAndCount(&bean)
	return total,bean, err
}

func (bc *BeanConnect) FindUserByCondi(condi map[string]interface{}) ([]*modelUser.User, error) {
	session := bc.Engine.NewSession()
	defer session.Close()
	for key, val := range condi {
		session = session.In(table.TbUser+"."+key, val)
	}
	bean := make([]*modelUser.User, 0)
	err := session.Table(table.TbUser).Find(&bean)
	return bean, err
}

func (bc *BeanConnect)FindRoleByCondi(condi map[string]interface{})([]*modelRole.Role,error){
	session := bc.Engine.NewSession()
	defer session.Close()
	for key, val := range condi {
		session = session.In(table.TbRole+"."+key, val)
	}
	bean := make([]*modelRole.Role, 0)
	err := session.Table(table.TbRole).Find(&bean)
	return bean, err
}

func (bc *BeanConnect)GetGameRecordByRoomId(roomId string)(bool,*modelGameInfo.GameRecord,error){
	bean:=modelGameInfo.GameRecord{}
	bean.RoomId = roomId
	exist,err:= bc.Engine.Table(table.TbGameRecord).Get(&bean)
	return exist,&bean,err
}

func (bc *BeanConnect)FindGameEvent(condi map[string]interface{})(int64,[]*modelGameInfo.GameEvent,error){
	session:=bc.Engine.NewSession()
	defer session.Close()
	for key,val:=range condi{
		session = session.In(table.TbGameEvent+"."+key,val)
	}
	bean:=make([]*modelGameInfo.GameEvent,0)
	total,err:=session.Table(table.TbGameEvent).FindAndCount(&bean)
	return total,bean,err
}

func (bc *BeanConnect)FindGameWho(condi map[string]interface{})(int64,[]*modelGameInfo.GameWho,error){
	session:=bc.Engine.NewSession()
	defer session.Close()
	for key,val:=range condi{
		session = session.In(table.TbGameWho+"."+key,val)
	}
	bean:=make([]*modelGameInfo.GameWho,0)
	total,err:=session.Table(table.TbGameWho).FindAndCount(&bean)
	return total,bean,err
}

func (bc *BeanConnect)FindRoomByCondi(condi map[string]interface{},page,perpage int)(int64,[]*modelRoom.Room,error){
	session:=bc.Engine.NewSession()
	defer session.Close()
	for key,val := range condi{
		session = session.In(table.TbRoom+"."+key,val)
	}
	bean:=make([]*modelRoom.Room,0)
	if perpage == -1{
		total,err:=session.Table(table.TbRoom).FindAndCount(&bean)
		return total,bean,err
	}
	total,err:=session.Table(table.TbRoom).Limit(perpage,(page-1)*perpage).FindAndCount(&bean)
	return total,bean,err
}