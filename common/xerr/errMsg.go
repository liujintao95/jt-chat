package xerr

var message map[uint32]string

const (
	OK = 0
)

const (
	DataLocked = iota + 10001
	RequestParamError
	DbError
	DbUpdateAffectedZeroError
	PermissionError
)

const (
	TokenExpireError = iota + 11001
	TokenGenerateError
)

const (
	UserMobilePwdError = iota + 20001
	UserAlreadyExists
	UserNotExists
	GroupNotExists
	GroupAlreadyExists
	ContactNotExists
	ContactAlreadyExists
	ApplicationNotExists
	ApplicationAlreadyExists
)

const (
	MsgNotExists = iota + 30001
	FileOpenErr
	FileSaveErr
)

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[RequestParamError] = "参数错误"
	message[TokenExpireError] = "token失效，请重新登陆"
	message[TokenGenerateError] = "生成token失败"
	message[DbError] = "数据库繁忙,请稍后再试"
	message[DbUpdateAffectedZeroError] = "更新数据影响行数为0"
	message[UserMobilePwdError] = "账号或密码不正确"
	message[UserAlreadyExists] = "用户已存在"
	message[UserNotExists] = "用户不存在"
	message[GroupNotExists] = "群聊组不存在"
	message[GroupAlreadyExists] = "群聊组已存在"
	message[DataLocked] = "数据被锁定"
	message[PermissionError] = "权限不足"
	message[ContactNotExists] = "联系人不存在"
	message[ContactAlreadyExists] = "联系人已存在"
	message[ApplicationNotExists] = "联系人申请不存在"
	message[ApplicationAlreadyExists] = "联系人申请已存在"
	message[MsgNotExists] = "消息不存在"
	message[FileOpenErr] = "打开文件错误"
	message[FileSaveErr] = "存入文件错误"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
