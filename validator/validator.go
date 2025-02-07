package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type CustomValidator struct{}

func (*CustomValidator) Domain() validator.Func {
	return func(fl validator.FieldLevel) bool {
		data := fl.Field().Interface().(string)
		b, err := regexp.Match("^((http)s?://)?[A-z\\d][-A-z\\d]{0,62}(/.[A-z\\d][-A-z\\d]{0,62})+/.?", []byte(data))
		if err == nil && b {
			return true
		} else {
			return false
		}
	}
}

func (*CustomValidator) Path() validator.Func {
	return func(fl validator.FieldLevel) bool {
		data := fl.Field().Interface().(string)
		b, err := regexp.Match("^(?:(?:[a-zA-Z]:|\\.{1,2})?[\\\\/](?:[^\\\\?/*|<>:\"]+[\\\\/])*)(?:(?:[^\\\\?/*|<>:\"]+?)(?:\\.[^.\\\\?/*|<>:\"]+)?)?$", []byte(data))
		if err == nil && b {
			return true
		} else {
			return false
		}
	}
}

func (*CustomValidator) NormalName() validator.Func {
	return func(fl validator.FieldLevel) bool {
		data := fl.Field().Interface().(string)
		b, err := regexp.Match("^[\u4E00-\u9FA5A-z\\d_]{2,20}$", []byte(data))
		if err == nil && b {
			return true
		} else {
			return false
		}
	}
}

func (*CustomValidator) ChineseName() validator.Func {
	return func(fl validator.FieldLevel) bool {
		data := fl.Field().Interface().(string)
		b, err := regexp.Match("^[\u9fa6-\u9fcb\u3400-\u4db5\u4e00-\u9fa5]{2,5}([\u25cf\u00b7][\u9fa6-\u9fcb\u3400-\u4db5\u4e00-\u9fa5]{2,5})*$", []byte(data))
		if err == nil && b {
			return true
		} else {
			return false
		}
	}
}

func (*CustomValidator) Phone() validator.Func {
	return func(fl validator.FieldLevel) bool {
		data := fl.Field().Interface().(string)
		b, err := regexp.Match("^(13\\d|14[5|7]|15[012356789]|18[012356789])\\d{8}$", []byte(data))
		if err == nil && b {
			return true
		} else {
			return false
		}
	}
}

func (*CustomValidator) IdentityCard() validator.Func {
	return func(fl validator.FieldLevel) bool {
		data := fl.Field().Interface().(string)
		b, err := regexp.Match("^(\\d{14}|\\d{17})(\\d[xX])$", []byte(data))
		if err == nil && b {
			return true
		} else {
			return false
		}
	}
}

func (*CustomValidator) CarNo() validator.Func {
	return func(fl validator.FieldLevel) bool {
		data := fl.Field().Interface().(string)
		b, err := regexp.Match("^([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领A-Z][a-zA-Z](([DF]((?![IO])[a-zA-Z\\d](?![IO]))\\d{4})|(\\d{5}[DF]))|[京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领A-Z][A-Z][A-Z\\d]{4}[A-Z\\d挂学警港澳])$", []byte(data))
		if err == nil && b {
			return true
		} else {
			return false
		}
	}
}
