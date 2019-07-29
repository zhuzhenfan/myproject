package room

import (
	"wolfkill/wolfkill/common/table"
	"wolfkill/wolfkill/module/room"
)

func (bc *BeanConnect)GetById(roomId string)(bool,*room.Room,error){
	bean:=room.Room{}
	bean.RoomId = roomId
	exist,err:= bc.Engine.Table(table.TbRoom).ID(bean.RoomId).Get(&bean)
	return exist,&bean,err
}

func (bc *BeanConnect)GetByOwner(userId string)(bool,*room.Room,error){
	bean:=room.Room{}
	bean.RoomOwner = userId
	exist,err:= bc.Engine.Table(table.TbRoom).
		In(table.TbRoom+"."+room.RoomOwner).Get(&bean)
	return exist,&bean,err
}


func (bc *BeanConnect)Find(condi map[string]interface{},
	page, perpage int)(int64,[]*room.Room,error){
	session:=bc.Engine.NewSession()
	defer  session.Close()
	for key,val:=range condi{
		session = session.In(table.TbRoom+"."+key,val)
	}
	bean:=make([]*room.Room,0)
	if perpage != -1{
		total,err:=session.Table(table.TbRoom).FindAndCount(&bean)
		return total,bean,err
	}
	total,err:=session.Table(table.TbRoom).
		Limit(perpage,(page-1)*perpage).FindAndCount(&bean)
	return total,bean,err
}

func (bc *BeanConnect)FindColString(colName string)(*[]string,error){
	var colValue []string
	// 取某一列
	// 返回切片指针 （for range *colValue）
	return &colValue ,bc.Engine.Table(table.TbRoom).
		Cols(table.TbRoom+"."+colName).Find(&colValue)
}

func (bc *BeanConnect)ExistRoomByCode(roomCode string)(bool,error){
	return bc.Engine.Table(table.TbRoom).
		In(table.TbRoom+"."+room.RoomCode,roomCode).Exist()
}