package user

import (
	"wolfkill/wolfkill/common/table"
	"wolfkill/wolfkill/module/user"
)
func (bc *BeanConnect) GetById(userId string) (bool, *user.User, error) {
	bean := user.User{}
	bean.UserId = userId
	exist, err := bc.Engine.Table(table.TbUser).Get(&bean)
	return exist, &bean, err
}

func (bc *BeanConnect) GetByName(userName string) (bool, *user.User, error) {
	bean := user.User{}
	bean.UserName = userName
	exist, err := bc.Engine.Table(table.TbUser).Get(&bean)
	return exist, &bean, err
}

func (bc *BeanConnect) Find(condition map[string]interface{},page,perpage int) (int64, []*user.User, error) {
	session := bc.Engine.NewSession()
	defer session.Close()

	for key, val := range condition {
		session = session.In(table.TbUser+"."+key, val)
	}

	bean := make([]*user.User, 0)
	
	if perpage==-1{
		count, err := session.FindAndCount(&bean)
		return count, bean, err
	}

	count, err := session.Limit(perpage,(page-1)*perpage).FindAndCount(&bean)
	return count, bean, err
}
