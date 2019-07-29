package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"wolfkill/wolfkill/common/http"
	"wolfkill/wolfkill/handler/assemble/game"
	"wolfkill/wolfkill/handler/role"
	"wolfkill/wolfkill/handler/room"
	"wolfkill/wolfkill/handler/user"
	"wolfkill/wolfkill/result/errs"
	"wolfkill/wolfkill/util/utilstring"
	"strconv"
	"strings"
	"time"
)

func UserMain(router *gin.Engine) {
	router.POST("/wolfkill/user", user.Add)
	router.DELETE("/wolfkill/user/", user.DeleteById)
	router.DELETE("/wolfkill/users", user.Delete)
	router.PUT("/wolfkill/user/", user.UpdateById)
	router.GET("/wolfkill/user/", user.GetById)
	router.GET("/wolfkill/users", user.Find)
	router.POST("/wolfkill/user/token", user.Login)
}

func RoleMain(router *gin.Engine) {
	router.PUT("/wolfkill/role/", role.UpdateById)
	router.GET("/wolfkill/role/", role.GetById)
	router.POST("/wolfkill/roles", role.FindByType)
}

func RoomMain(router *gin.Engine) {
	router.POST("/wolfkill/room", room.Insert)
	router.PUT("/wolfkill/room", room.UpdateById)
	router.GET("/wolfkill/room", room.GetById)
}

func GameMain(router *gin.Engine) {
	router.POST("/wolfkill/game", game.CreateRoomToGame)
	router.PUT("/wolfkill/game", game.JoinGame)
	router.GET("/wolfkill/game", game.ListPlayer)
	router.PUT("/wolfkill/game/status",game.GameStatusOp)
	router.DELETE("/wolfkill/game",game.CloseRoom)
}

// allow request for Cross region and option request
func XOptions(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}

// token func
func MiddleWareFunc() gin.HandlerFunc {
	return func(c *gin.Context) { middlewareImpl(c) }
}

const (
	authorization = "Authorization"
	bearer        = "Bearer"
	expiredTime   = 24 * 60 * 60 * 2
)

type userIdJson struct {
	UserId string `json:"user_id"`
}

// token function impl
func middlewareImpl(c *gin.Context) {
	if c.Request.URL.String() == "/wolfkill/user/token" ||
		c.Request.URL.String() == "/wolfkill/user" {
		c.Next()
		return
	}
	auth := c.GetHeader(authorization)
	if auth == "" {
		c.AbortWithStatusJSON(401, gin.H{"msg": errs.Auth_NullAuthorization})
		c.Abort()
		return
	}
	tokenStr := strings.Trim(auth, bearer+" ")
	tokenParas, err := utilstring.SimpleTokenParase(tokenStr)
	if err != nil {
		c.AbortWithStatusJSON(http.CodeAuthError, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}
	stringArr := strings.Split(tokenParas, "--")
	if len(stringArr) != 2 {
		c.AbortWithStatusJSON(http.CodeAuthError, gin.H{"msg": errs.Auth_TokenErr})
		c.Abort()
		return
	}
	timeStamp, err := strconv.Atoi(stringArr[0])
	if err != nil {
		c.AbortWithStatusJSON(http.CodeAuthError, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}
	if time.Now().Unix()-int64(timeStamp) > expiredTime {
		c.AbortWithStatusJSON(http.CodeAuthError, gin.H{"msg": errs.Auth_TokenExpiredTime})
		c.Abort()
		return
	}

	// token内置的用户id和参数用户id对比
	userId := c.PostForm("user_id")
	if userId == "" {
		userId = c.Query("user_id")
		if userId == "" {
			// body里面的数据只能用一次，所以要取出来备份，再重新赋值回去
			var idJson = userIdJson{}
			var bodyBytes []byte
			if c.Request.Body != nil {
				bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
			}
			err = json.Unmarshal(bodyBytes, &idJson)
			if idJson.UserId == "" {
				c.AbortWithStatusJSON(http.CodeAuthError, gin.H{"msg": errs.User_NullId})
				c.Abort()
				return
			}
			userId = idJson.UserId
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}
	}
	if userId != stringArr[1] || userId == "" {
		c.AbortWithStatusJSON(http.CodeAuthError, gin.H{"msg": errs.Auth_TokenErr})
		c.Abort()
		return
	}
	c.Next()
}
