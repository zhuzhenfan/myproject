// main.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"wolfkill/wolfkill/common"
	"wolfkill/wolfkill/controller"
	"wolfkill/wolfkill/controller/assemble/game"
	"wolfkill/wolfkill/dao/factory"
	"wolfkill/wolfkill/handler"
)

func Seven(){
	for i:=1;i<500;i++{
		if i%7==0{
			fmt.Println(i)
			continue
		}
		if strings.Contains(strconv.Itoa(i),"7")==true{
			fmt.Println(i)
		}
	}
}
func main() {
	fmt.Println("Hello World!")
	Seven()
	if err := common.InitFunc(common.WolfKillConf, common.WolfKillLinuxConf,common.WolfKillWindowsConf); err != nil {
		fmt.Println(err)
		return
	}
	if err := common.CreatePGEngine(); err != nil {
		fmt.Println(err)
		return
	}
	if err := factory.TableFactory(); err != nil {
		fmt.Println(err)
		return
	}

	if err :=controller.RoleInit();err != nil{
		fmt.Println(err)
		return
	}
	if err:=game.InitThread();err != nil{
		fmt.Println(err)
		return
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	wolfkillMain(router)

	fmt.Println(common.ServerAddr+":"+common.ServerPort)
	fmt.Println("success")
	fmt.Println(router.Run(common.ServerAddr+":"+common.ServerPort))
}

func wolfkillMain(router *gin.Engine){
	//router.Use(handler.XOptions,handler.MiddleWareFunc())
	router.Use(handler.XOptions)
	handler.UserMain(router)
	handler.RoleMain(router)
	handler.RoomMain(router)
	handler.GameMain(router)
}