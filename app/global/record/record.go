package record

/*
通用的结构数据都在这里存储
*/
type Player struct {
	UUID       int64
	OpenId     string
	SessionKey string
}
