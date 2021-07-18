package global

type Player struct {
	UUID       int64  // 角色id
	OpenID     string // 通行wx返回openid
	SessionKey string // 通行vx返回SessionKey
	Token      string // 校验本次登陆信息
	State      int    // 登录状态

	Auth int // 玩家权限等级

	Lv int //
	Exp int64 //

	Bag       map[int]interface{}    // 背包类型 => 背包数据
	Money     map[int]int64          // 货币类型 => 货币数量
	Playing   map[string]interface{} //玩法信息 => 玩法数据
	Task      map[int]interface{}    // 任务类型 =>任务数据
	Cultivate map[int]interface{}    // 养成类型 => 养成信息
	Mail      map[int64]interface{}  // 邮件创建时间戳 => 邮件信息
	Log       []interface{}          // 日志数据
	Goal      []interface{}          // 目标数据
	TimeClock map[string]interface{} // 时间相关数据

}
