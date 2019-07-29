package assemble

type UserRole struct {
	UserId string `json:"user_id"`
	RoleId string `json:"role_id"`
	Status bool `json:"status"`
	DeadWay string `json:"dead_way"`
}
