package errs

import "errors"

func ErrFunc(str string) error {
	if str == "" {
		return errors.New(public_nullErrorObj)
	}
	return errors.New(str)
}


// public error
const (
	Public_UnKnowError  = "未知错误"
	Public_NullObj = "空对象异常"

	public_nullErrorObj = "错误实体为空"

)

// auth error
const (
	Auth_PublicKeyErr      = "公钥(yue)不正确"
	Auth_PrivateKeyErr     = "私钥（yue）不正确"
	Auth_NullAuthorization = "auth认证为空"
	Auth_TokenErr          = "token错误"
	Auth_TokenExpiredTime  = "token过期，请重新登录"
	Auth_UserIdErr         = "用户id解析异常"
)

// string error
const (
	Str_NullString = "空字符串"
)

// sql error
const (
	SqlErr_Session = "sql: 数据库连接异常"
	SqlErr_Insert  = "sql: 插入数据异常"
	SqlErr_Delete  = "sql: 删除数据异常"
	SqlErr_Update  = "sql: 更新数据异常"
	SqlErr_Query   = "sql: 查询数据异常"
)

// user error
const (
	User_Exist       = "用户已存在"
	User_NotExist    = "用户不存在"
	User_NullId      = "用户id为空"
	User_NullPasswd  = "用户密码为空"
	User_OPNumber    = "用户数量操作异常"
	User_NullName    = "用户名称为空"
	User_PassWordErr = "用户密码错误"
)

// role error
const (
	Role_NullId      = "角色id为空"
	Role_NullName    = "角色名字为空"
	Role_NullSkill   = "角色技能为空"
	Role_NullExplain = "角色说明为空"
	Role_NullType    = "角色类型为空"
	Role_OPNumber    = "用户数量操作异常"
	Role_NotExist    = "角色不存在"
)

// room error
const (
	Room_NullId               = "房间id为空"
	Room_NullCode             = "房间号码为空"
	Room_NotExist             = "房间不存在"
	Room_NullOwner            = "房间创建者为空"
	Room_FullPlayer           = "房间人数已满"
	Room_PassWordErr          = "房间密码错误"
	Room_NullPassWord         = "房间密码不能为空"
	Room_NullGameStatus       = "房间游戏状态不能为空"
	Room_GameStatusOverErr    = "当前游戏状态无法开始"
	Room_GameStatusWaitingErr = "当前游戏状态无法结束"
	Room_GameStatusGamingErr  = "当前游戏状态无法回到房间"
)

// game error
const (
	Game_UserExist       = "该用户已经在一个房间中"
	Game_NullUserId      = "创建者不能为空"
	Game_RoleNumErr      = "游戏人数不能少于3人"
	Game_StatusErr       = "游戏状态不正确"
	Game_PlayerNotEnough = "房间尚未满员"
	Game_TaskTypeNull = "空任务类型"
)

// gameinfo error
const (
	GameInfo_RecordNotExist = "游戏记录不存在"
)

// thread error
const (
	Thread_FullThread = "服务器繁忙"
	Thread_TimeOut = "请求超时"
)