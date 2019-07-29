// common
package common

var (
	wkConfPath        = "/etc/tools/conf/wolfkill/wolfkill.conf"
	wkConfLinuxPath   = "../conf/wolfkill.conf"
	wkConfWindowsPath = "wolfkill/conf/wolfkill.conf"

	roleJsonPath        = "/etc/tools/conf/wolfkill/roleinfo.json"
	roleJsonLinuxPath   = "../conf/wolfkill/roleinfo.json"
	roleJsonWindowsPath = "wolfkill/conf/roleinfo.json"
)
var (
	WolfKillConf        = wkConfPath
	WolfKillLinuxConf   = wkConfLinuxPath
	WolfKillWindowsConf = wkConfWindowsPath

	RoleJsonPath        = roleJsonPath
	RoleJsonLinuxPath   = roleJsonLinuxPath
	RoleJsonWindowsPath = roleJsonWindowsPath
)

//log path
var (
	LogPath = "/var/tools/log/wolfkill.log"
	LogDay  = 7
)

var (
	ServerAddr string
	ServerPort string

	dbType     string
	dbAddr     string
	dbPort     string
	dbName     string
	dbUserName string
	dbPassWord string
)
