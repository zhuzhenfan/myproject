package controller

import (
	"encoding/json"
	"io/ioutil"
	"wolfkill/wolfkill/common"
	dao "wolfkill/wolfkill/dao/role"
	module "wolfkill/wolfkill/module/role"
	"wolfkill/wolfkill/result/errs"
	"wolfkill/wolfkill/util/utilstring"
	"time"
)

// 天真的想法遇上现实的不方便。。。
var chineseName = make(map[string]string)

// 浅拷贝
func GetChineseName() map[string]string {
	// 即使meke了新的map，但是把其他map赋值给他时，二者共用内存中同一份map
	// 最好初始话的时候指定长度，go会优先根据大小分配到栈上，速度快，减少gc的压力
	// 若make 0的话，go是分配到堆上，频繁申请会造成内存碎片
	return chineseName
}

func RoleInit() error {
	// read content for json file
	content, err := readRoleJson()
	if err != nil {
		return err
	}
	bean := module.RoleJson{}
	// 遇到复杂的json或者想拿key 用map[string]interface{}
	err = json.Unmarshal(content, &bean)
	if err != nil {
		return err
	}

	//for _, name := range bean.RoleChineseNames {
	//	for key, val := range name {
	//		chineseName[key] = val
	//	}
	//}
	// check role exits
	bc := dao.BeanConnect{}
	bc.Session = dao.GetSession()
	bc.Engine = dao.GetEngine()
	defer bc.Session.Close()

	_, result, err := bc.Find(map[string]interface{}{}, 1, -1)
	if err != nil {
		return err
	}
	existMap := make(map[string]string)
	for _,val:=range result{
		existMap[val.RoleName]=""
	}
	roleArr:=make([]*module.Role,0)
	for _,val:=range bean.RoleInfos{
		if _,ok:=existMap[val.Name];ok==false{
			roleInfo:=module.Role{}
			roleInfo.RoleName = val.Name
			roleInfo.RoleChinaName = val.ChinaName
			roleInfo.RoleOrderNumber = val.OrderNumber
			roleInfo.RoleId = utilstring.SimpleUUID()
			roleInfo.RoleType = val.Types
			roleInfo.RoleSkill = val.Skill
			roleInfo.RoleExplain = val.Explain
			roleInfo.RoleOpenNumber = val.OpenNumber
			roleArr = append(roleArr,&roleInfo)
		}
	}
	err = bc.Session.Begin()
	if err != nil{
		return err
	}
	// insert role to database
	total,err:=bc.Insert(roleArr...)
	if err != nil{
		return err
	}
	if total != int64(len(roleArr)){
		return errs.ErrFunc(errs.SqlErr_Insert)
	}

	return bc.Session.Commit()
}

func readRoleJson() ([]byte, error) {
	jsonFile, err := ioutil.ReadFile(common.RoleJsonPath)
	if err != nil {
		jsonFile, err = ioutil.ReadFile(common.RoleJsonLinuxPath)
		if err != nil {
			jsonFile, err = ioutil.ReadFile(common.RoleJsonWindowsPath)
			if err != nil {
				return jsonFile, err
			}
		}
	}
	return jsonFile, nil
}

func RoleAdd(bean ...*module.Role) error {
	bc := dao.BeanConnect{}
	bc.Session = dao.GetSession()
	defer bc.Session.Close()

	// 判断空值 赋值两个number
	for key, _ := range bean {
		bean[key].RoleId = utilstring.Md5String(bean[key].RoleName + time.Now().String())
		if bean[key].RoleName == "" {
			return errs.ErrFunc(errs.Role_NullName)
		}
		if bean[key].RoleId == "" {
			return errs.ErrFunc(errs.Role_NullId)
		}
		if bean[key].RoleExplain == "" {
			return errs.ErrFunc(errs.Role_NullExplain)
		}
		if bean[key].RoleSkill == "" {
			return errs.ErrFunc(errs.Role_NullSkill)
		}
		if bean[key].RoleType == "" {
			return errs.ErrFunc(errs.Role_NullType)
		}
	}

	total, err := bc.Insert(bean...)
	if err != nil {
		return err
	}
	if total == 0 {
		return errs.ErrFunc(errs.Role_OPNumber)
	}
	return err
}

func RoleDelete(bean *module.Role) error {
	bc := dao.BeanConnect{}
	common.GetEngine()
	bc.Session = dao.GetSession()
	defer bc.Session.Close()

	count, err := bc.Delete(bean)
	if err != nil {
		return err
	}
	if count != 1 {
		return errs.ErrFunc(errs.Role_OPNumber)
	}
	return nil
}

func RoleUpdateById(bean *module.Role) error {
	bc := dao.BeanConnect{}
	bc.Session = dao.GetSession()
	defer bc.Session.Close()

	count, err := bc.UpdateById(bean)
	if err != nil {
		return err
	}
	if count != 1 {
		return errs.ErrFunc(errs.Role_OPNumber)
	}
	return nil
}

func RoleGet(roleId string) (*module.Role, error) {
	bc := dao.BeanConnect{}
	bc.Engine = dao.GetEngine()

	exist, result, err := bc.GetById(roleId)
	if err != nil {
		return nil, err
	}
	if exist == false {
		return nil, errs.ErrFunc(errs.Role_NotExist)
	}
	return result, nil
}

func RoleFindByType(roleType string, page, perpage int) (int64, *[]*module.Role, error) {
	bc := dao.BeanConnect{}
	bc.Engine = dao.GetEngine()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 && perpage != -1 {
		perpage = 10
	}
	condi := make(map[string]interface{})
	if roleType != "" {
		condi[module.RoleType] = roleType
	}
	total,result,err := bc.Find(condi, page, perpage)
	if err != nil{
		return 0,nil,err
	}

	return total,&result,nil
}
