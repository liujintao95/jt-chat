package constant

import "time"

const (
	UidKey = "uid"
)

const (
	AdminName = "admin"
	AdminUid  = "0000001"
)

const (
	TransportTypeWebRTC = iota
	TransportTypeNormal
)

const (
	MsgPong    = "pong"
	MsgWelcome = "欢迎登录!"
)

const (
	ContentTypeText = iota
	ContentTypeFile
	ContentTypePicture
	ContentTypeAudio
	ContentTypeVideo
)

const (
	ChatSingle = UserContactType
	ChatGroup  = GroupContactType
)

const (
	WriteWait  = 10 * time.Second
	PongWait   = 60 * time.Second
	PingPeriod = (PongWait * 9) / 10
)
