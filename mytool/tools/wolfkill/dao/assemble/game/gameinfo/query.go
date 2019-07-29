package gameinfo

import (
	"wolfkill/wolfkill/common/table"
	"wolfkill/wolfkill/module/assemble/game"
	"wolfkill/wolfkill/module/role"
)

func (bc *BeanConnect)GetUserByCode(){}

func (bc *BeanConnect)GetRoleById(roleId string)(bool,*role.Role,error){
	bean:=role.Role{}
	bean.RoleId = roleId
	exist,err:=bc.Engine.Table(table.TbRole).ID(bean.RoleId).Get(&bean)
	return exist,&bean,err
}

func (bc *BeanConnect)GetGameByUserId(userId string)(bool,*game.Game,error){
	bean:=game.Game{}
	exist,err:=bc.Engine.Table(table.TbGame).In(game.UserId,userId).Get(&bean)
	return exist,&bean,err
}