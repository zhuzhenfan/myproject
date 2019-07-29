// user
package user

import (
	"github.com/gin-gonic/gin"
	"wolfkill/wolfkill/common/http"
	control "wolfkill/wolfkill/controller"
	module "wolfkill/wolfkill/module/user"
	"wolfkill/wolfkill/result/errs"
)

const (
	user_id = "user_id"
)
// Add add a user
func Add(c *gin.Context) {
	var param module.User
	if err := c.BindJSON(&param); err != nil {
		http.ParamErr(c, err)
		return
	}
	if param.UserName == ""{
		http.ParamErr(c, errs.ErrFunc(errs.User_NullName))
	}
	if param.PassWord == ""{
		http.ParamErr(c, errs.ErrFunc(errs.User_NullPasswd))
		return
	}

	err := control.UserAdd(&param)
	if err != nil{
		http.Failed(c,err)
		return
	}
	http.Ok(c,http.SuccessStr)
}

// DeleteById delete user by id
func DeleteById(c *gin.Context) {
	var userId = c.PostForm(user_id)
	if userId == ""{
		http.ParamErr(c, errs.ErrFunc(errs.User_NullId))
		return
	}
	err:=control.UserDeleteById(userId)
	if err != nil{
		http.Failed(c,err)
		return
	}
	http.Ok(c,http.SuccessStr)
}

// Delete delete user by condition
func Delete(c *gin.Context) {
	var param *module.User
	if err := c.BindJSON(&param); err != nil {
		http.ParamErr(c, err)
		return
	}
	err:=control.UserDelete(param)
	if err != nil{
		http.Failed(c,err)
		return
	}
	http.Ok(c,http.SuccessStr)
}

// UpdateById update user by id
func UpdateById(c *gin.Context) {
	var param *module.User
	if err := c.BindJSON(&param); err != nil {
		http.ParamErr(c, err)
		return
	}
	err:=control.UserUpdateById(param)
	if err != nil{
		http.Failed(c,err)
		return
	}
	http.Ok(c,http.SuccessStr)
}

// GetById get user by id
func GetById(c *gin.Context) {
	var userId = c.Query(user_id)
	if userId == ""{
		http.ParamErr(c, errs.ErrFunc(errs.User_NullId))
		return
	}
	result,err:=control.UserGetById(userId)
	if err != nil{
		http.Failed(c,err)
		return
	}
	http.Ok(c,result)
}

func Find(c *gin.Context) {
	var param module.FindIn
	if err:=c.BindJSON(&param);err != nil{
		http.ParamErr(c, err)
		return
	}
	total,result,err:=control.UserFindByName(param.UserName, param.Page,param.Perpage)
	if err != nil{
		http.Failed(c,err)
		return
	}
	http.Ok(c,map[string]interface{}{"total":total,"result":result})
}

// Login user need to user_name and pass_word to login
func Login(c *gin.Context){
	var param module.LoginIn
	if err:=c.BindJSON(&param);err != nil{
		http.ParamErr(c,err)
	}
	if param.UserName == ""{
		http.ParamErr(c, errs.ErrFunc(errs.User_NullName))
	}
	if param.PassWord == ""{
		http.ParamErr(c, errs.ErrFunc(errs.User_NullPasswd))
	}
	result,err:=control.UserLogin(param.UserName,param.PassWord)
	if err != nil{
		http.Failed(c,err)
		return
	}
	http.Ok(c,result)
}