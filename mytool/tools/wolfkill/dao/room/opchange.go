package room

import (
	"wolfkill/wolfkill/common/table"
	"wolfkill/wolfkill/module/room"
)

func (bc *BeanConnect)Insert(bean ...*room.Room)(int64,error){
	return bc.Session.Table(table.TbRoom).Insert(&bean)
}

func (bc *BeanConnect)DeleteById(bean *room.Room)(int64,error){
	return bc.Session.Table(table.TbRoom).ID(bean.RoomId).Delete(bean)
}

func (bc *BeanConnect)Deletes(condi map[string]interface{})(int64,error){
	for key,val:=range condi{
		bc.Session = bc.Session.In(table.TbRoom+"."+key,val)
	}
	bean:=room.Room{}
	return bc.Session.Table(table.TbRoom).Delete(&bean)
}

func (bc *BeanConnect)UpdateById(bean *room.Room)(int64,error){
	return bc.Session.Table(table.TbRoom).ID(bean.RoomId).Update(bean)
}
