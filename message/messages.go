package message

type Code int

const (
	Success Code = 0

	RequestError Code = 40000
	ServerError  Code = 50000

	PermissionDenied  Code = 40001
	AuthFailure       Code = 40002
	AuthUnableUpdate  Code = 40003
	AuthHaveExpired   Code = 40004
	MoreThanFrequency Code = 40005

	UsernameOrPasswordError Code = 40101
	UsernameRepeatError     Code = 40102
	UserEmailRepeatError    Code = 40103
	UserCreateError         Code = 40104
	UserUpdateError         Code = 40105
	UserNotExist            Code = 40106
	UserDeleteError         Code = 40107
	UserPhoneUpdateError    Code = 40108
	UserPhoneCantUpdate     Code = 40109
	UserPhoneNotExist       Code = 40110
	UserPhoneAlreadyExist   Code = 40111
	UserEmailAlreadyExist   Code = 40112
)

var messages = map[Code]string{
	Success:                 "成功",
	UsernameOrPasswordError: "用户密码错误",
	UsernameRepeatError:     "账户名称重复",
	UserEmailRepeatError:    "账户邮箱已存在",
	UserCreateError:         "用户创建失败",
	UserUpdateError:         "用户更新失败",
	UserNotExist:            "用户不存在",
	UserDeleteError:         "用户删除失败",
	UserPhoneUpdateError:    "用户更换绑定手机号失败",
	UserPhoneCantUpdate:     "用户无法更换绑定手机号",
	UserPhoneNotExist:       "用户的手机号不存在",
	UserPhoneAlreadyExist:   "用户手机号已被使用",
	UserEmailAlreadyExist:   "用户邮箱已绑定其他用户",
	PermissionDenied:        "用户权限不足",
	AuthFailure:             "用户认证信息已失效",
	AuthUnableUpdate:        "用户认证信息无法更新",
	AuthHaveExpired:         "用户认证信息已过期",
}

func GetMessage(code Code) string {
	return messages[code]
}
