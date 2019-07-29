// op
package user

import (
	"wolfkill/wolfkill/common/table"
	"wolfkill/wolfkill/module/user"
)

func (bc *BeanConnect) Insert(bean ...*user.User) (int64,error){
	return bc.Session.Table(table.TbUser).Insert(&bean)
}

func (bc *BeanConnect) Delete(bean *user.User) (int64,error) {
	return bc.Session.Table(table.TbUser).Delete(bean)
}

func (bc *BeanConnect)Deletes(condi map[string]interface{})(int64,error){
	for key,val:=range condi{
		bc.Session = bc.Session.In(table.TbUser+"."+key,val)
	}
	bean:=user.User{}
	return bc.Session.Delete(&bean)
}

func (bc *BeanConnect) UpdateById(bean *user.User) (int64,error) {
	return bc.Session.Table(table.TbUser).ID(bean.UserId).Update(bean)
}

func (bc *BeanConnect)UpdataOnlineById(bean *user.User)(int64,error){
	return bc.Session.Table(table.TbUser).ID(bean.UserId).Update(bean)
}