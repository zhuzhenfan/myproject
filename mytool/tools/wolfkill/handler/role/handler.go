package role

import (
	"github.com/gin-gonic/gin"
	"wolfkill/wolfkill/common/http"
	"wolfkill/wolfkill/controller"
	"wolfkill/wolfkill/module/role"
	"wolfkill/wolfkill/result/errs"
)

const (
	role_id="role_id"
	role_type="role_type"
)
func UpdateById(c *gin.Context){
	var param role.Role
	err:=c.BindJSON(&param)
	if err != nil{
		http.ParamErr(c,err)
		return
	}
	if param.RoleId == ""{
		http.ParamErr(c, errs.ErrFunc(errs.Role_NullId))
		return
	}
	err = controller.RoleUpdateById(&param)
	if err != nil{
		http.Failed(c,err)
		return
	}
	http.Ok(c,http.SuccessStr)
}

func GetById(c *gin.Context){
	roleId:=c.PostForm(role_id)
	if roleId == ""{
		http.ParamErr(c, errs.ErrFunc(errs.Role_NullId))
		return
	}
	result,err:=controller.RoleGet(roleId)
	if err != nil{
		http.Failed(c,err)
		return
	}
	http.Ok(c,result)
}

func FindByType(c *gin.Context){
	var param role.FindTypeIn
	err:=c.BindJSON(&param)
	if err != nil{
		http.ParamErr(c,err)
	}
	total,result,err:=controller.RoleFindByType(param.RoleType,param.Page,param.Perpage)
	if err != nil{
		http.Failed(c,err)
	}
	http.Ok(c,map[string]interface{}{"total":total,"result":result})
}
