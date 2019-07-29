package gameinfo

const (
	RoomId = "room_id"
	Winner = "winner"
	Away = "away"
	WinnerLive = "winner_live"
	WinRound = "win_round"
	GameEventId = "game_event_id"
	EventTime = "event_time"
	Number = "number"
	Round = "round"
	CreateTime = "create_time"
	WhoBeforeId = "who_before_id"
	Skill = "skill"
	Doing = "doing"
	WhoAfterId = "who_after_id"
)

const (
	DayTime = "dayTime"
	NightTime = "nightTime"
)

// 游戏记录表
type GameRecord struct {
	// 游戏房间id
	RoomId string `json:"room_id" xorm:"pk unique notnull"`
	// 本局游戏胜方式
	Winner string `json:"winner"`
	// 胜利方式
	Away string `json:"away"`
	// 胜方生存人数
	WinnerLive string `json:"winner_live"`
	// 胜利的轮次
	WinRound int `json:"win_round"`
	// 创建时间
	CreateTime int64 `json:"create_time" xorm:"created"`
}

// 游戏事件表
type GameEvent struct {
	// 游戏事件id
	GameEventId string `json:"game_event_id" xorm:"pk unique notnull"`
	// 游戏房间id
	RoomId string `json:"room_id" xorm:"notnull"`
	// 事件时间
	EventTime string `json:"event_time" xorm:"notnull"`
	// 编号
	Number int `json:"number" xorm:"notnull"`
	// 创建时间
	CreateTime int64 `json:"create_time" xorm:"created"`
}

// 游戏玩家交互表
type GameWho struct {
	// 游戏事件id
	GameEventId string `json:"game_event_id" xorm:"notnull"`
	// 主动方
	WhoBeforeId string `json:"who_before_id" xorm:"notnull"`
	// 使用技能
	Skill string `json:"skill" xorm:"notnull"`
	// 事件动作
	Doing string `json:"doing" xorm:"notnull"`
	// 被动方
	WhoAfterId string `json:"who_after_id" xorm:"notnull"`
}

// 游戏op
type GameOpIn struct {
	RoomId      string   `json:"room_id"`
	Number      int      `json:"number"`
	RoleId      string   `json:"role_id"`
	RoleName    string   `json:"role_name"`
	WhoBeforeId string   `json:"who_before_id"`
	WhoAfterIds []string `json:"who_after_ids"`
}