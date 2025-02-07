package response

import (
	"gin_work/message"
	"net/http"
)

type response struct {
	Code    int         `json:"code" yaml:"code" xml:"code" bson:"code"`
	Message string      `json:"message" yaml:"message" xml:"message" bson:"message"`
	Data    interface{} `json:"data" yaml:"data" xml:"data" bson:"data"`
}

func newResponse(code int, message string, data interface{}) *response {
	return &response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func setSuccessResponse(r *response, args ...interface{}) {
	if args != nil {
		if len(args) > 2 {
			r.Data = args[0]
			r.Code = int(args[1].(message.Code))
			r.Message = args[2].(string)
		} else if len(args) > 1 {
			r.Data = args[0]
			r.Code = int(args[1].(message.Code))
		} else {
			r.Data = args[0]
		}
	}
}

func setFailResponse(r *response, args ...interface{}) {
	if args != nil {
		if len(args) > 2 {
			r.Code = int(args[0].(message.Code))
			r.Message = args[1].(string)
			r.Data = args[2]
		} else if len(args) > 1 {
			r.Code = int(args[0].(message.Code))
			r.Message = args[1].(string)
		} else {
			r.Code = args[0].(int)
		}
	}
}

func failResponse(code int, args ...interface{}) (int, *response) {
	r := newResponse(code, "Fail", nil)

	var params = make([]interface{}, 3)
	if len(args) > 0 {
		params[0] = args[0]
		mes := message.GetMessage(args[0].(message.Code))
		if len(args) == 1 && mes != "" {
			params[1] = mes
		} else {
			params[1] = args[1]
		}
		if len(args) == 3 {
			params[2] = args[2]
		}
	}
	setFailResponse(r, params...)

	return code, r
}

func Success(args ...interface{}) (int, *response) {
	r := newResponse(0, "Success", nil)
	setSuccessResponse(r, args...)
	return http.StatusOK, r
}

func Fail(args ...interface{}) (int, *response) {
	return failResponse(http.StatusInternalServerError, args...)
}

func RequestFail(args ...interface{}) (int, *response) {
	return failResponse(http.StatusBadRequest, args...)
}
