package controller

import (
	dao "wolfkill/wolfkill/dao/room"
	module "wolfkill/wolfkill/module/room"
	"wolfkill/wolfkill/result/errs"
	"wolfkill/wolfkill/util/utilstring"
)

func RoomAdd(roomOwner, RoomPassWord string,roomPlayerNum int) (string,error) {
	bc := dao.BeanConnect{}
	bc.Engine = dao.GetEngine()
	exist,_,err:=bc.GetByOwner(roomOwner)
	if err != nil{
		return "",err
	}
	if exist == true{
		return "", errs.ErrFunc(errs.Game_UserExist)
	}
	roomCode:=utilstring.RandNumber(4)
	exist,err =bc.ExistRoomByCode(roomCode)
	if err != nil{
		return "",err
	}

	bc.Session = dao.GetSession()
	defer bc.Session.Close()

	bean := module.Room{}
	bean.RoomPassWord = RoomPassWord
	bean.RoomOwner = roomOwner
	bean.RoomPlayerNum = roomPlayerNum
	if exist == false{
		bean.RoomCode = roomCode
	}
	bean.RoomId = utilstring.SimpleUUID()
	_, err = bc.Insert(&bean)
	return bean.RoomCode,err
}

func RoomGetById(roomId string) (*module.Room, error) {
	bc := dao.BeanConnect{}
	bc.Engine = dao.GetEngine()
	exist, result, err := bc.GetById(roomId)
	if err != nil {
		return nil, err
	}
	if exist == false {
		return nil, errs.ErrFunc(errs.Room_NotExist)
	}
	return result, nil
}

func RoomFindByCode(roomCode string, page, perpage int) (int64, []*module.Room, error) {
	bc := dao.BeanConnect{}
	bc.Engine = dao.GetEngine()
	return bc.Find(
		map[string]interface{}{module.RoomCode: roomCode}, page, perpage)
}

func UpdateGameStatus(roomId, gameStatus string)error{
	bc := dao.BeanConnect{}
	bc.Session = dao.GetSession()
	defer bc.Session.Close()

	status := module.GetStatus(gameStatus)
	if status == ""{
		return errs.ErrFunc(errs.Game_StatusErr)
	}

	bean:=module.Room{}
	bean.RoomId = roomId
	bean.RoomGameStatus = status
	total,err := bc.UpdateById(&bean)
	if err != nil{
		return err
	}
	if total != 1{
		return errs.ErrFunc(errs.SqlErr_Update)
	}
	return nil
}