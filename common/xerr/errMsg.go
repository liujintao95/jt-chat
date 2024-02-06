package xerr

var message map[uint32]string

const OK uint32 = 0
const DataLocked uint32 = 100001
const RequestParamError uint32 = 100002
const TokenExpireError uint32 = 100003
const TokenGenerateError uint32 = 100004
const DbError uint32 = 100005
const DbUpdateAffectedZeroError uint32 = 100006
const UserMobilePwdError uint32 = 100007
const UserAlreadyExists uint32 = 100008
const UserNotExists uint32 = 100009
const GroupNotExists uint32 = 100010
const GroupAlreadyExists uint32 = 100011
const PermissionError uint32 = 100012
const ContactNotExists uint32 = 100013

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
