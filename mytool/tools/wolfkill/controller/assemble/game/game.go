package game

import (
	"fmt"
	"wolfkill/wolfkill/controller/assemble/game/gameinfo"
	daoAsse "wolfkill/wolfkill/dao/assemble/game"
	moduleGame "wolfkill/wolfkill/module/assemble/game"
	moduleGameinfo "wolfkill/wolfkill/module/assemble/game/gameinfo"
	moduleRole "wolfkill/wolfkill/module/role"
	moduleRoom "wolfkill/wolfkill/module/room"
	moduleUser "wolfkill/wolfkill/module/user"
	"wolfkill/wolfkill/result/errs"
	"wolfkill/wolfkill/util/utilstring"
	"time"
)

// 一台服务器的房间数量上限
var roomNum int = 50
var roomChan = make(chan int, roomNum)

// 记录每个房间线程的内置的chan
var roomChanMap = make(map[string]*roomThread)

// 初始进程
func InitThread() error {
	bc := daoAsse.BeanConnect{}
	bc.Engine = daoAsse.GetEngine()
	// 查询数据库当前存在的房间
	_, result, err := bc.FindRoomByCondi(map[string]interface{}{}, 1, -1)
	if err != nil {
		return err
	}
	// 重新建立守护线程
	for _, val := range result {
		var err error
		defer func() {
			if err != nil {
				<-roomChan
			}
		}()
		if len(roomChan) >= roomNum {
			return errs.ErrFunc(errs.Thread_FullThread)
		}
		roomChan <- 1
		total, games, err2 := bc.FindGameByRoomId(val.RoomId, 1, -1)
		if err2 != nil {
			err = err2
			return err
		}
		rt := NewRoomThread(int(total), 10, 5)
		rt.roomId = val.RoomId
		sortBeanGame := make([]moduleGame.Game, int(total))
		for _, item := range games {
			sortBeanGame[item.UserNumber-1] = *item
		}
		rt.game = sortBeanGame
		err = rt.setRoomThread()
		if err != nil {
			return err
		}
		go rt.goRoomThread()
	}
	return nil
}

// 初始化taskChan对象（防止空指针异常）
func NewTaskChan() *taskChan {
	gameData := moduleGame.Game{}
	result := taskResult{}
	result.gameData = &gameData
	task := taskChan{}
	task.result = &result
	return &task
}

// 初始化roomThread对象（防止空指针异常）
func NewRoomThread(gameLen, taskOnlyNum, taskManyNum int) *roomThread {
	if gameLen < 0 {
		gameLen = 0
	}
	if taskOnlyNum <= 0 {
		taskOnlyNum = 10
	}
	if taskManyNum <= 0 {
		taskManyNum = 10
	}
	rt := roomThread{}
	rt.taskOnly = make(chan *taskChan, taskOnlyNum)
	rt.taskMany = make(chan *taskChan, taskManyNum)
	rt.game = make([]moduleGame.Game, gameLen)
	return &rt
}

// 房间守护线程struct
type roomThread struct {
	roomId   string
	game     []moduleGame.Game
	taskOnly chan *taskChan
	taskMany chan *taskChan
	err      error
}

// 线程任务strcut
type taskChan struct {
	taskType string
	userId   string
	roomId   string
	session  *daoAsse.BeanConnect
	result   *taskResult
}

type taskResult struct {
	err      error
	ok       bool
	gameData *moduleGame.Game
}

func (rt *roomThread) goRoomThread() {
	// 占用和释放房间chan
	defer func() {
		<-roomChan
		fmt.Println("end")
	}()
	fmt.Println("start")

	roomChanMap[rt.roomId] = rt

	// thread doing
	defer func() {
		if p := recover(); p != nil {
			<-roomChan
			rt.err = p.(error)
			return
		}
	}()
	timeOut := time.NewTimer(time.Hour * 24)
	for {
		select {
		case task := <-roomChanMap[rt.roomId].taskOnly:
			if task.taskType == moduleGame.FirstGetRole {
				err := task.fistGetRoleFunc()
				if err != nil {
					task.result.err = err
					break
				}
			}
			if task.taskType == moduleGame.CloseRoom {
				err := closeRoom(task.roomId)
				if err != nil {
					task.result.err = err
				}else{
					task.result.ok = true
					return
				}
			}

		case taskMany := <-roomChanMap[rt.roomId].taskMany:
			func(task *taskChan) {

			}(taskMany)
		case <-timeOut.C:
			return
		}
	}
}

func (rt *roomThread) setRoomThread() error {
	if rt.roomId == "" {
		return errs.ErrFunc(errs.Room_NullId)
	}
	if len(rt.game) <= 3 {
		return errs.ErrFunc(errs.Game_RoleNumErr)
	}
	roomChanMap[rt.roomId] = rt
	return nil
}

func (tc *taskChan) taskCreate() error {
	if tc.roomId == "" {
		return errs.ErrFunc(errs.Room_NullId)
	}
	if tc.taskType == "" {
		return errs.ErrFunc(errs.Game_TaskTypeNull)
	}

	roomChanMap[tc.roomId].taskOnly <- tc
	return nil
}

func (tc *taskChan) fistGetRoleFunc() error {
	if tc.roomId == "" {
		return errs.ErrFunc(errs.Room_NullId)
	}
	if tc.userId == "" {
		return errs.ErrFunc(errs.User_NullId)
	}
	// 按顺序挑选角色
	var index int
	var bean moduleGame.Game
	for key, val := range roomChanMap[tc.roomId].game {
		if val.Select == false {
			index = key
			bean = val
			break
		}
	}
	if bean.GameId == "" {
		tc.result.err = errs.ErrFunc(errs.Room_FullPlayer)
	}

	// 加入游戏
	bean.UserId = tc.userId
	bean.Select = true
	bean.Status = true
	bc := daoAsse.BeanConnect{}
	bc.Session = daoAsse.GetSession()
	defer bc.Session.Close()

	num, err := bc.UpdateGameByGameId(&bean, []string{
		moduleGame.Select, moduleGame.UserId, moduleGame.Status})
	bc.Session.Close()
	if err != nil {
		return err
	}
	if num != 1 {
		return err
	}

	roomChanMap[tc.roomId].game[index].UserId = tc.userId
	roomChanMap[tc.roomId].game[index].Select = true
	roomChanMap[tc.roomId].game[index].Status = true
	tc.result.gameData = &bean
	return nil
}

func CreateRoom(bean *moduleGame.CreateGameIn) (string, string, error) {
	var err error
	if len(roomChan) >= roomNum {
		return "", "", errs.ErrFunc(errs.Thread_FullThread)
	}
	defer func(err *error) {
		if err != nil {
			<-roomChan
		}
	}(&err)
	roomChan <- 1

	roomId, roomCode, beanGame, err2 := createRoomToGame(bean)
	if err != nil {
		err = err2
		return "", "", err
	}
	if roomId == "" {
		return "", "", err
	}
	// 数据赋予守护线程
	rt := NewRoomThread(0, 10, 5)
	rt.roomId = roomId
	sortBeanGame := make([]moduleGame.Game, len(bean.RoleIds))
	for _, val := range *beanGame {
		sortBeanGame[val.UserNumber-1] = val
	}
	rt.game = sortBeanGame
	err = rt.setRoomThread()
	if err != nil {
		return "", "", err
	}
	go rt.goRoomThread()

	return roomId, roomCode, err
}

func createRoomToGame(bean *moduleGame.CreateGameIn) (string, string, *[]moduleGame.Game, error) {
	roomId := utilstring.SimpleUUID()

	bc := daoAsse.BeanConnect{}
	bc.Engine = daoAsse.GetEngine()

	// 检查创建者是否已经包含法官或者玩家的身份
	exist, _, err := bc.GetUserForRoom(bean.RoomOwner)
	if err != nil {
		return "", "", nil, err
	}
	if exist == true {
		return "", "", nil, errs.ErrFunc(errs.Game_UserExist)
	}
	exist, _, err = bc.GetUserForGame(bean.RoomOwner)
	if err != nil {
		return "", "", nil, err
	}
	if exist == true {
		return "", "", nil, errs.ErrFunc(errs.Game_UserExist)
	}
	// 检测该随机房间号码是否已经存在，若存在则采取自增号码
	roomCode := utilstring.RandNumber(4)
	exist, err = bc.ExistRoomByCode(roomCode)
	if err != nil {
		return "", "", nil, err
	}
	// 创建房间
	bc.Session = daoAsse.GetSession()
	defer bc.Session.Close()
	err = bc.Session.Begin()
	if err != nil {
		return "", "", nil, err
	}
	bean.RoomId = roomId
	bean.RoomPlayerNum = len(bean.RoleIds)
	bean.RoomGameStatus = moduleRoom.Waiting
	if exist == false {
		bean.RoomCode = roomCode
	}
	num, err := bc.InsertRoom(&bean.Room)
	if err != nil {
		return "", "", nil, err
	}
	if num != 1 {
		return "", "", nil, errs.ErrFunc(errs.SqlErr_Insert)
	}

	// 创建绑定房间角色 (这里不用指针是一个对象)
	beanGame := make([]moduleGame.Game, len(bean.RoleIds))
	temp := moduleGame.Game{}
	temp.RoomId = bean.RoomId
	for key, val := range bean.RoleIds {
		temp.GameId = utilstring.SimpleUUID()
		temp.RoleId = val
		beanGame[key] = temp
	}
	// 赋予编号
	numberMap := make(map[int]int)
	for i := 1; i <= len(bean.RoleIds); i++ {
		numberMap[i] = i
	}
	for key, _ := range numberMap {
		beanGame[key-1].UserNumber = key
	}
	num, err = bc.InsertGame(&beanGame)
	if err != nil {
		return "", "", nil, err
	}
	if num != int64(len(bean.RoleIds)) {
		return "", "", nil, errs.ErrFunc(errs.SqlErr_Insert)
	}

	return bean.RoomId, bean.RoomCode, &beanGame, bc.Session.Commit()
}

// 玩家加入游戏
func JoinGame(userId, roomCode, roomPassWord string) (*moduleRole.Role, error) {
	bc := daoAsse.BeanConnect{}
	bc.Engine = daoAsse.GetEngine()

	// 检查创建者是否已经包含法官(房间是否存在)
	exist, _, err := bc.GetUserForRoom(userId)
	if err != nil {
		return nil, err
	}
	if exist == true {
		return nil, err
	}
	// 检查玩家的身份是否已在游戏中
	exist, gameInfo, err := bc.GetUserForGame(userId)
	if err != nil {
		return nil, err
	}
	if exist == true {
		exist, roleInfo, err := bc.GetRoleById(gameInfo.RoleId)
		if err != nil {
			return nil, err
		}
		if exist == false {
			return nil, errs.ErrFunc(errs.Role_NotExist)
		}
		return roleInfo, nil
	}

	// 获取房间信息
	exist, roomInfo, err := bc.GetRoomByCode(roomCode)
	if err != nil {
		return nil, err
	}
	if exist == false {
		return nil, errs.ErrFunc(errs.Room_NotExist)
	}

	//验证密码
	if roomPassWord == "" && roomInfo.RoomPassWord != "" {
		return nil, errs.ErrFunc(errs.Room_NullPassWord)
	}
	if roomPassWord != roomInfo.RoomPassWord {
		return nil, errs.ErrFunc(errs.Room_PassWordErr)
	}
	roleId := ""
	task := NewTaskChan()
	task.roomId = roomInfo.RoomId
	task.userId = userId
	task.taskType = moduleGame.FirstGetRole
	err = task.taskCreate()
	if err != nil {
		return nil, err
	}
	startTime := time.Now().Unix()
	for {
		if task.result.err != nil {
			return nil, err
		}
		if task.result.gameData.RoleId != "" {
			roleId = task.result.gameData.RoleId
			break
		}
		if time.Now().Unix()-startTime >= 10 {
			return nil, errs.ErrFunc(errs.Thread_TimeOut)
		}
	}
	// 查询角色的信息
	exist, roleInfo, err := bc.GetRoleById(roleId)
	if exist == false {
		return nil, errs.ErrFunc(errs.Role_NotExist)
	}
	return roleInfo, err
}

func ListPlayer(roomId string) ([]*moduleGame.ListPlayerOut, error) {
	bc := daoAsse.BeanConnect{}
	bc.Engine = daoAsse.GetEngine()
	// 查询房间内游戏玩家的id
	_, games, err := bc.FindGameByRoomId(roomId, 1, -1)
	if err != nil {
		return nil, err
	}

	// 查询玩家的信息
	userIds := make([]string, 0)
	roleIds := make([]string, 0)
	for _, val := range games {
		if val.UserId != "" {
			userIds = append(userIds, val.UserId)
		}
		roleIds = append(roleIds, val.RoleId)
	}

	users := make([]*moduleUser.User, 0)
	roles := make([]*moduleRole.Role, 0)
	userChan := make(chan int, 1)
	roleChan := make(chan int, 1)
	if len(userIds) != 0 {
		go func(uc chan int) {
			users, err = bc.FindUserByCondi(
				map[string]interface{}{moduleUser.UserId: userIds})
			uc <- 1
		}(userChan)
	}

	// 游戏规则  role不可能没有
	go func(rc chan int) {
		roles, err = bc.FindRoleByCondi(
			map[string]interface{}{moduleRole.RoleId: roleIds})
		rc <- 1
	}(roleChan)

	if len(userIds) != 0 {
		<-userChan
	}
	<-roleChan
	if err != nil {
		return nil, err
	}
	userMap := make(map[string]*moduleUser.User)
	roleMap := make(map[string]*moduleRole.Role)
	go func(uc chan int) {
		for _, val := range users {
			userMap[val.UserId] = val
		}
		userChan <- 1
	}(userChan)
	go func(rc chan int) {
		for _, val := range roles {
			roleMap[val.RoleId] = val
		}
		roleChan <- 1
	}(roleChan)
	if len(userIds) != 0 {
		<-userChan
	}
	<-roleChan

	result := make([]*moduleGame.ListPlayerOut, len(games))
	for key, val := range games {
		temp := moduleGame.ListPlayerOut{}
		if val.UserId != "" {
			temp.User = *userMap[val.UserId]
		}
		temp.Role = *roleMap[val.RoleId]
		result[key] = &temp
	}
	return result, nil
}

func CloseRoom(roomId string) error {
	task := NewTaskChan()
	task.roomId = roomId
	task.taskType = moduleGame.CloseRoom
	err := task.taskCreate()
	if err != nil {
		return err
	}
	for {
		if task.result.ok == true {
			break
		}
		if task.result.err != nil {
			return task.result.err
		}
	}
	// 释放chan
	delete(roomChanMap, roomId)
	return nil
}

func closeRoom(roomId string) error {
	bc := daoAsse.BeanConnect{}
	bc.Engine = daoAsse.GetEngine()

	// 查询相关数据
	total, events, err := bc.FindGameEvent(map[string]interface{}{moduleGameinfo.RoomId: roomId})
	if err != nil {
		return err
	}
	eventIds := make([]string, total)
	for index, item := range events {
		eventIds[index] = item.GameEventId
	}

	// 清楚相关数据
	bc.Session = daoAsse.GetSession()
	defer bc.Session.Close()
	err = bc.Session.Begin()
	if err != nil {
		return err
	}
	_, err = bc.DeleteRoom(map[string]interface{}{moduleRoom.RoomId: roomId})
	if err != nil {
		return err
	}
	_, err = bc.DeleteGame(map[string]interface{}{moduleRoom.RoomId: roomId})
	if err != nil {
		return err
	}
	_, err = bc.DeleteGameRecord(map[string]interface{}{moduleRoom.RoomId: roomId})
	if err != nil {
		return err
	}
	_, err = bc.DeleteGameEvent(map[string]interface{}{moduleRoom.RoomId: roomId})
	if err != nil {
		return err
	}
	_, err = bc.DeleteGameWho(map[string]interface{}{moduleGameinfo.GameEventId: eventIds})
	if err != nil {
		return err
	}
	return bc.Session.Commit()
}

func GameStart(roomId string) error {
	bc := daoAsse.BeanConnect{}
	bc.Engine = daoAsse.GetEngine()

	exist, roomInfo, err := bc.GetRoomById(roomId)
	if err != nil {
		return err
	}
	if exist == false {
		return errs.ErrFunc(errs.Room_NotExist)
	}
	if roomInfo.RoomGameStatus == moduleRoom.Gaming {
		return nil
	}
	if roomInfo.RoomGameStatus == moduleRoom.Over {
		return errs.ErrFunc(errs.Room_GameStatusOverErr)
	}
	// 查询所有角色是否都有玩家存在
	_, games, err := bc.FindGameByRoomId(roomId, 1, -1)
	if err != nil {
		return err
	}
	for _, val := range games {
		if val.UserId == "" {
			return errs.ErrFunc(errs.Game_PlayerNotEnough)
		}
	}

	// 查询事件记录（用于删除who）
	total, events, err := bc.FindGameEvent(map[string]interface{}{moduleGameinfo.RoomId: roomId})
	if err != nil {
		return err
	}
	eventIds := make([]string, total)
	for key, val := range events {
		eventIds[key] = val.GameEventId
	}

	bc.Session = daoAsse.GetSession()
	defer bc.Session.Close()
	err = bc.Session.Begin()
	if err != nil {
		return err
	}
	// 清空旧的游戏记录
	_, err = bc.DeleteGameRecord(map[string]interface{}{moduleGameinfo.RoomId: roomId})
	if err != nil {
		return err
	}
	_, err = bc.DeleteGameEvent(map[string]interface{}{moduleGameinfo.RoomId: roomId})
	if err != nil {
		return err
	}
	_, err = bc.DeleteGameWho(map[string]interface{}{moduleGameinfo.GameEventId: eventIds})
	if err != nil {
		return err
	}

	// 更新房间状态
	beanRoom := moduleRoom.Room{}
	beanRoom.RoomId = roomId
	beanRoom.RoomGameStatus = moduleRoom.Gaming
	num, err := bc.UpdateRoomById(&beanRoom)
	if err != nil {
		return err
	}
	if num != 1 {
		return errs.ErrFunc(errs.SqlErr_Update)
	}
	// 创建游戏记录
	beanGameRecord := moduleGameinfo.GameRecord{}
	beanGameRecord.RoomId = roomId
	num, err = bc.InsertGameRecord(&beanGameRecord)
	if err != nil {
		return err
	}
	if num != 1 {
		return errs.ErrFunc(errs.SqlErr_Insert)
	}
	return bc.Session.Commit()
}

func GameOver(roomId string) error {
	bc := daoAsse.BeanConnect{}
	bc.Engine = daoAsse.GetEngine()

	// 查询房间信息
	exist, roomInfo, err := bc.GetRoomById(roomId)
	if err != nil {
		return err
	}
	if exist == false {
		return errs.ErrFunc(errs.Room_NotExist)
	}
	if roomInfo.RoomGameStatus == moduleRoom.Over {
		return nil
	}
	if roomInfo.RoomGameStatus == moduleRoom.Waiting {
		return errs.ErrFunc(errs.Room_GameStatusWaitingErr)
	}

	// 切断玩家和角色
	beanGame := moduleGame.Game{}
	beanGame.RoomId = roomId
	beanGame.UserId = ""
	beanGame.Status = false
	beanGame.DeadWay = ""
	beanGame.Select = false

	bc.Session = daoAsse.GetSession()
	defer bc.Session.Close()
	err = bc.Session.Begin()
	if err != nil {
		return err
	}
	num, err := bc.UpdateGameByRoomId(&beanGame, []string{moduleGame.UserId,
		moduleGame.Status, moduleGame.DeadWay, moduleGame.Select})
	if err != nil {
		return err
	}
	fmt.Println(num, int64(roomInfo.RoomPlayerNum))
	if num != int64(roomInfo.RoomPlayerNum) {
		return errs.ErrFunc(errs.SqlErr_Update)
	}

	// 修改房间状态
	beanRoom := moduleRoom.Room{}
	beanRoom.RoomId = roomId
	beanRoom.RoomGameStatus = moduleRoom.Over
	num, err = bc.UpdateRoomById(&beanRoom)
	if err != nil {
		return err
	}
	if num != 1 {
		return errs.ErrFunc(errs.SqlErr_Update)
	}

	return bc.Session.Commit()
}

func GameWait(roomId string) error {
	bc := daoAsse.BeanConnect{}
	bc.Engine = daoAsse.GetEngine()
	// 查询房间信息
	exist, roomInfo, err := bc.GetRoomById(roomId)
	if err != nil {
		return err
	}
	if exist == false {
		return errs.ErrFunc(errs.Room_NotExist)
	}
	if roomInfo.RoomGameStatus == moduleRoom.Waiting {
		return nil
	}
	if roomInfo.RoomGameStatus == moduleRoom.Gaming {
		return errs.ErrFunc(errs.Room_GameStatusGamingErr)
	}

	// 更新房间信息
	bc.Session = daoAsse.GetSession()
	defer bc.Session.Close()
	beanRoom := moduleRoom.Room{}
	beanRoom.RoomId = roomId
	beanRoom.RoomGameStatus = moduleRoom.Waiting
	num, err := bc.UpdateRoomById(&beanRoom)
	if err != nil {
		return err
	}
	if num != 1 {
		return errs.ErrFunc(errs.SqlErr_Update)
	}
	return nil
}

func GameOP(in moduleGameinfo.GameOpIn) {
	// 玩家操作
	gameinfo.Prophet(in)

	// 群体操作

	// 其他操作
}
