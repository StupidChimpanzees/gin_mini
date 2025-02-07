package model

import (
	"errors"
	"fmt"
	"gin_work/wrap/utils"
	"gorm.io/gorm"
	"reflect"
	"strings"
	"time"
)

type SoftDelete struct {
	DeleteField string
}

func AND(where [][]interface{}) (interface{}, []interface{}) {
	var condition []string
	var params []interface{}
	for _, s := range where {
		condition = append(condition, fmt.Sprintf("%s %s ?", s[0], s[1]))
		params = append(params, s[2])
	}
	if len(condition) <= 0 {
		return "", params
	} else {
		return "( " + strings.Join(condition, " AND ") + " )", params
	}
}

func Or(where [][]interface{}) (interface{}, []interface{}) {
	var condition []string
	var params []interface{}
	for _, s := range where {
		condition = append(condition, fmt.Sprintf("%s %s ?", s[0], s[1]))
		params = append(params, s[2])
	}
	if len(condition) <= 0 {
		return "", params
	} else {
		return "( " + strings.Join(condition, " OR ") + " )", params
	}
}

func SoftWhere(model interface{}, where interface{}) interface{} {
	fieldName := utils.TypeDefaultValue(model, "SoftDelete")
	if fieldName == "" {
		return where
	}
	if condition := where.(string); condition == "" {
		return fieldName + " IS NULL"
	} else {
		return "( " + condition + " ) AND " + fieldName + " IS NULL"
	}
}

func SoftDel(tx *gorm.DB, model interface{}, where interface{}, args ...interface{}) (int64, error) {
	tagValue := utils.TypeDefaultValue(model, "SoftDelete")
	if tagValue == "" {
		return 0, errors.New("delete field is not exist")
	}
	field := utils.FieldByTagValue(model, "gorm", tagValue)
	switch utils.FieldType(model, field) {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result := tx.Model(&model).Where(SoftWhere(model, where), args...).Updates(map[string]interface{}{tagValue: time.Now().Unix()})
		return result.RowsAffected, result.Error
	default:
		result := tx.Model(&model).Where(SoftWhere(model, where), args).Updates(map[string]string{tagValue: time.Now().Format("2006-01-02 15:04:05")})
		return result.RowsAffected, result.Error
	}
}
