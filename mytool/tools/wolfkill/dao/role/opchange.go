package role

import (
	"wolfkill/wolfkill/common/table"
	"wolfkill/wolfkill/module/role"
)

func (bc *BeanConnect)Insert(bean ...*role.Role)(int64,error){
	return bc.Session.Table(table.TbRole).Insert(&bean)
}

func (bc *BeanConnect)Delete(bean *role.Role)(int64,error){
	return bc.Session.Table(table.TbRole).Delete(bean)
}

func (bc *BeanConnect)Deletes(condi map[string]interface{}){}

func (bc *BeanConnect)UpdateById(bean *role.Role)(int64,error){
	return bc.Session.Table(table.TbRole).ID(bean.RoleId).Update(bean)
}
