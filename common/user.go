package common

import (
	"gin_work/message"
	"strings"
)

func GetPwd(password, salt string) string {
	encrypt := Sha256Encode([]byte(password))
	return Sha256Encode([]byte(encrypt + salt))
}

func CheckPwd(password, enPassword, salt string) message.Code {
	if !strings.EqualFold(GetPwd(password, salt), enPassword) {
		return message.UsernameOrPasswordError
	}
	return message.Success
}
