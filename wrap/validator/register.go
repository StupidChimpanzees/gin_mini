package validator

import (
	validator2 "gin_work/validator"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"reflect"
)

func Load() {
	var customValidator *validator2.CustomValidator
	validatorType := reflect.TypeOf(customValidator)
	validatorValue := reflect.ValueOf(customValidator)
	for i := 0; i < validatorValue.NumMethod(); i++ {
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			methodName := validatorType.Method(i).Name
			funcResult := validatorValue.Method(i).Call(nil)
			err := v.RegisterValidation(methodName, funcResult[0].Interface().(validator.Func))
			if err != nil {
				log.Fatalln(err.Error())
			}
		}
	}
}
