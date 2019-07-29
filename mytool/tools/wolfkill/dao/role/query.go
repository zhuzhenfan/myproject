package role

import (
	"wolfkill/wolfkill/common/table"
	"wolfkill/wolfkill/module/role"
)

const (
	DESC = "DESC"
	ASC = "ASC"
)

func (bc *BeanConnect) GetById(roleId string) (bool, *role.Role, error) {
	bean := role.Role{}
	bean.RoleId = roleId
	exist, err := bc.Engine.Table(table.TbRole).Get(&bean)
	return exist, &bean, err
}

func (bc *BeanConnect)GetByName(roleName string)(bool, *role.Role, error){
	bean := role.Role{}
	bean.RoleName = roleName
	exist, err := bc.Engine.Table(table.TbRole).Get(&bean)
	return exist, &bean, err
}

func (bc *BeanConnect) Find(condition map[string]interface{},
	page, perpage int) (int64, []*role.Role, error) {
	session := bc.Engine.NewSession()
	defer session.Close()
	for key, val := range condition {
		session = session.In(table.TbRole+"."+key, val)
	}
	bean := make([]*role.Role, 0)
	if perpage == -1 {
		total, err := session.Table(table.TbRole).FindAndCount(&bean)
		return total, bean, err
	}
	total, err := session.Table(table.TbRole).
		Limit(perpage, (page-1)*perpage).
		OrderBy(table.TbRole+"."+role.RoleOrderNumber+ASC).FindAndCount(&bean)
	return total, bean, err
}
