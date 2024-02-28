package constant

const (
	TransportTypeWebRTC = iota
	TransportTypeNormal
)

const (
	MsgPong    = "pong"
	MsgWelcome = "欢迎登录!"
)

const (
	AdminName = "admin"
	AdminUid  = "0000001"
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
