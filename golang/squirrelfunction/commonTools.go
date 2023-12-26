package squirrelfunction

import (
	"fmt"
	"reflect"
	"strings"
)

func GetDBTagFromStructField(field reflect.StructField) string {
	dbTagValue := field.Tag.Get("db")
	switch dbTagValue {
	case "-":
		return ""
	case "":
		return ""
	default:
		if strings.Contains(dbTagValue, ",") {
			dbTagValue = strings.Split(dbTagValue, ",")[0]
		}
		dbTagValue = strings.TrimSpace(dbTagValue)
		dbTagValue = fmt.Sprintf("`%s`", dbTagValue)
		return dbTagValue
	}
}

func GetColumns(obj any, removeKeys *[]string) (keyColumns *[]string, dbColumns *[]string) {
	keyColumns = &[]string{}
	dbColumns = &[]string{}
	reflectVal := reflect.ValueOf(obj)
	reflectVal = reflect.Indirect(reflectVal)
	reflectType := reflectVal.Type()
	for i := 0; i < reflectVal.NumField(); i++ {
		typeField := reflectType.Field(i)
		if IsRemoveKey(removeKeys, typeField.Name) {
		} else {
			dbTag := GetDBTagFromStructField(typeField)
			if dbTag != "" {
				*dbColumns = append(*dbColumns, dbTag)
				*keyColumns = append(*keyColumns, typeField.Name)
			}
		}
	}
	return keyColumns, dbColumns
}

func IsRemoveKey(removeKeys *[]string, key string) bool {
	for _, v := range *removeKeys {
		if v == key {
			return true
		}
	}
	return false
}
