package room

import (
	"github.com/gin-gonic/gin"
	"wolfkill/wolfkill/common/http"
	"wolfkill/wolfkill/controller"
	"wolfkill/wolfkill/module/room"
	"wolfkill/wolfkill/result/errs"
)

func Insert(c *gin.Context){
	param:=room.InsertIn{}
	err:=c.BindJSON(&param)
	if err != nil{
		http.ParamErr(c,err)
		return
	}
	if param.RoomOwner == ""{
		http.ParamErr(c, errs.ErrFunc(errs.Room_NullOwner))
		return
	}
	if param.RoomPlayerNum >=3{
		http.ParamErr(c, errs.ErrFunc(errs.Game_RoleNumErr))
		return
	}

	result,err := controller.RoomAdd(param.RoomOwner,param.RoomPassWord,param.RoomPlayerNum)
	if err != nil{
		http.Failed(c,err)
		return
	}
	http.Ok(c,result)
}

func GetById(c *gin.Context){
	var roomId = c.Query(room.RoomId)
	if roomId == ""{
		http.ParamErr(c, errs.ErrFunc(errs.Room_NullId))
		return
	}
	result,err:=controller.RoomGetById(roomId)
	if err != nil{
		http.Failed(c,err)
		return
	}
	http.Ok(c,result)
}

func Find(c *gin.Context){
	var param room.FindIn
	err:=c.BindJSON(&param)
	if err != nil{
		http.ParamErr(c,err)
	}
	total,result,err:=controller.RoomFindByCode(param.RoomCode,param.Page,param.Perpage)
	if err != nil{
		http.Failed(c,err)
	}
	http.Ok(c,map[string]interface{}{"total":total,"result":result})
}

func UpdateById(c *gin.Context){
	var param = room.UpdateGameStatusIn{}
	err:=c.BindJSON(&param)
	if err != nil{
		http.ParamErr(c,err)
		return
	}
	if param.RoomId == ""{
		http.ParamErr(c, errs.ErrFunc(errs.Room_NullId))
		return
	}
	if param.RoomGameStatus == ""{
		http.ParamErr(c, errs.ErrFunc(errs.Room_NullGameStatus))
		return
	}
	err = controller.UpdateGameStatus(param.RoomId,param.RoomGameStatus)
	if err != nil{
		http.Failed(c,err)
		return
	}
	http.Ok(c,http.SuccessStr)
}