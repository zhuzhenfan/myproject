// user
package controller

import (
	"bytes"
	dao "wolfkill/wolfkill/dao/user"
	module "wolfkill/wolfkill/module/user"
	"wolfkill/wolfkill/result/errs"
	"wolfkill/wolfkill/util/utilstring"
	"strconv"
	"time"
)

func UserAdd(bean *module.User) error {
	bc := dao.BeanConnect{}
	bc.Engine = dao.GetEngine()
	bc.Session = dao.GetSession()
	defer bc.Session.Close()

	if bean.UserId == ""{
		bean.UserId = utilstring.SimpleUUID()
	}

	// 判断是否存在
	exist,_,err:=bc.GetByName(bean.UserName)
	if err != nil{
		return err
	}
	if exist == true{
		return errs.ErrFunc(errs.User_Exist)
	}
	bean.UserOnline = true
	bean.CreateTime = time.Now().Unix()
	count,err:= bc.Insert(bean)
	if err != nil{
		return err
	}
	if count != 1{
		return errs.ErrFunc(errs.User_OPNumber)
	}
	return nil
}

func UserDeleteById(userId string)error{
	bc := dao.BeanConnect{}
	bc.Session = dao.GetSession()
	defer bc.Session.Close()
	bean:=module.User{}
	bean.UserId  = userId
	count,err:= bc.Delete(&bean)
	if err != nil{
		return err
	}
	if count != 1{
		return errs.ErrFunc(errs.User_OPNumber)
	}
	return nil
}

func UserDelete(bean *module.User)error{
	bc := dao.BeanConnect{}
	bc.Session = dao.GetSession()
	defer bc.Session.Close()
	_,err:= bc.Delete(bean)
	return err
}

func UserUpdateById(bean *module.User)error{
	bc := dao.BeanConnect{}
	bc.Session = dao.GetSession()
	defer bc.Session.Close()
	count,err:= bc.UpdateById(bean)
	if err != nil{
		return err
	}
	if count != 1{
		return errs.ErrFunc(errs.User_OPNumber)
	}
	return nil
}

func UserGetById(userId string)(*module.User,error){
	bc := dao.BeanConnect{}
	bc.Engine = dao.GetEngine()
	exist,result,err:=bc.GetById(userId)
	if err != nil{
		return nil,err
	}
	if exist == false{
		return nil, errs.ErrFunc(errs.User_NotExist)
	}
	return result,nil
}

func UserFindByName(userName string,page,perpage int)(int64,[]*module.User,error){
	bc := dao.BeanConnect{}
	bc.Engine = dao.GetEngine()
	// 条件拼接
	condition:=make(map[string]interface{})
	if userName != ""{
		condition[module.UserName]=userName
	}
	if page <=0{
		page =1
	}
	if perpage<=0 && perpage != -1{
		perpage=10
	}
	return bc.Find(condition,page,perpage)
}

// user login
func UserLogin(userName,passWord string)(*module.LoginOut,error){
	bc := dao.BeanConnect{}
	bc.Engine = dao.GetEngine()
	bc.Session = dao.GetSession()
	defer bc.Session.Close()

	exist,userInfo,err:=bc.GetByName(userName)
	if err != nil{
		return nil,err
	}
	if exist == false{
		return nil, errs.ErrFunc(errs.User_NotExist)
	}
	if passWord != userInfo.PassWord{
		return nil, errs.ErrFunc(errs.User_PassWordErr)
	}
	userInfo.LoginTime = time.Now().Unix()
	userInfo.UserOnline = true
	_,err=bc.UpdateById(userInfo)
	if err != nil{
		return nil,err
	}

	var buffer bytes.Buffer
	timeStamp := time.Now().Unix()
	buffer.WriteString(strconv.FormatInt(timeStamp,10))
	buffer.WriteString("--")
	buffer.WriteString(userInfo.UserId)

	token,err:=utilstring.SimpleTokenCreate(buffer.String())
	result:=module.LoginOut{}
	result.User = userInfo
	result.Token = token
	result.PassWord = ""
	return &result,err
}