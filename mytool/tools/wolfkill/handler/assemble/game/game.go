package game

import (
	"github.com/gin-gonic/gin"
	"wolfkill/wolfkill/common/http"
	controller "wolfkill/wolfkill/controller/assemble/game"
	moduleGame "wolfkill/wolfkill/module/assemble/game"
	moduleRoom "wolfkill/wolfkill/module/room"
	"wolfkill/wolfkill/result/errs"
	"wolfkill/wolfkill/result/ok"
)

// create game for room
func CreateRoomToGame(c *gin.Context){
	param := moduleGame.CreateGameIn{}
	err := c.BindJSON(&param)
	if err != nil{
		http.ParamErr(c,err)
		return
	}
	if param.RoomOwner == ""{
		http.ParamErr(c, errs.ErrFunc(errs.Game_NullUserId))
		return
	}
	if len(param.RoleIds)<=3{
		http.ParamErr(c, errs.ErrFunc(errs.Game_RoleNumErr))
		return
	}
	roomId,roomCode,err:=controller.CreateRoom(&param)
	if err != nil{
		http.Failed(c,err)
		return
	}
	http.Ok(c,moduleGame.CreateGameOut{RoomId:roomId,RoomCode:roomCode})
}

func JoinGame(c *gin.Context){
	var param = moduleGame.JoinGameIn{}
	err:= c.BindJSON(&param)
	if err != nil{
		http.ParamErr(c,err)
		return
	}
	if param.UserId == ""{
		http.ParamErr(c, errs.ErrFunc(errs.User_NullId))
		return
	}
	if param.RoomCode == ""{
		http.ParamErr(c, errs.ErrFunc(errs.Room_NullCode))
		return
	}
	result,err:=controller.JoinGame(param.UserId,
		param.RoomCode,param.RoomPassWord)
	if err != nil{
		http.Failed(c,err)
		return
	}
	http.Ok(c,result)
}

func ListPlayer(c *gin.Context){
	var roomId  = c.Query(moduleRoom.RoomId)
	if roomId == ""{
		http.ParamErr(c, errs.ErrFunc(errs.Room_NullId))
		return
	}
	result,err:=controller.ListPlayer(roomId)
	if err != nil{
		http.Failed(c,err)
		return
	}
	http.Ok(c,result)
}

func GameStatusOp(c *gin.Context){
	var param = moduleGame.GameStatus{}
	var err error
	err = c.BindJSON(&param)
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
	if param.RoomGameStatus == moduleRoom.GamingStr{
		err = controller.GameStart(param.RoomId)
	}
	if param.RoomGameStatus == moduleRoom.OverStr{
		err = controller.GameOver(param.RoomId)
	}
	if param.RoomGameStatus == moduleRoom.WaitingStr{
		err = controller.GameWait(param.RoomId)
	}

	if err != nil{
		http.Failed(c,err)
		return
	}
	http.Ok(c,http.SuccessStr)
}

func CloseRoom(c *gin.Context){
	var roomId  = c.Query(moduleRoom.RoomId)
	if roomId == ""{
		http.ParamErr(c, errs.ErrFunc(errs.Room_NullId))
		return
	}
	if err := controller.CloseRoom(roomId);err != nil{
		http.Failed(c,err)
		return
	}
	http.Ok(c,ok.Public_Ok)
}