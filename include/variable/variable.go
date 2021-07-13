package variable

var (
	// Player 相关
	PlayerBaseID = int64(10000)
	PlayerNumKey = "playerNumKey"

	// PlayerSvr 相关
	PlayerSvr interface{}

	// websocket 相关
	WSHub interface{}
)

// 全局前后端信息格式
type Message struct {
	Sender    	string
	Recver 		string
	Content   	string
	Rsp       	interface{}
}
