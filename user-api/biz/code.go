package biz

const Ok = 200

var (
	DBError         = NewError(10000, "数据库错误")
	AlreadyRegister = NewError(10100, "用户已注册")
	NoRegister      = NewError(10100, "改用户还未注册")
	TokenError      = NewError(10200, "token 有问题")
	InvalidToken    = NewError(10300, "token 失效")
	UserNotExist    = NewError(10400, "用户不存在")
)
