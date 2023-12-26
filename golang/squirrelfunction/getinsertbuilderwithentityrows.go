package squirrelfunction

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Masterminds/squirrel"
)

func GetInsertBuilderWithEntityRows(entitys *[]any, insertBuilder squirrel.InsertBuilder, removeKeys *[]string) (bool, squirrel.InsertBuilder) {
	if len(*entitys) == 0 {
		return false, insertBuilder
	}
	insertBuilder.columns(columns)
	columns := GetColumns((*entitys)[0], removeKeys)

	for _, entity := range *entitys {
		val := reflect.ValueOf(entity)
		val = reflect.Indirect(val)
		typ := val.Type()
		for i := 0; i < val.NumField(); i++ {
			// valField := val.Field(i)
			typeField := typ.Field(i)
			dbTagValue := typeField.Tag.Get("db")
			// keyName := typeField.Name
			switch dbTagValue {
			case "-":
				continue
			case "":
				continue
			default:
				if strings.Contains(dbTagValue, ",") {
					dbTagValue = strings.Split(dbTagValue, ",")[0]
				}
				dbTagValue = strings.TrimSpace(dbTagValue)
				dbTagValue = fmt.Sprintf("`%s`", dbTagValue)
				fmt.Printf("\n dbTagValue : %+v\n", dbTagValue)
				// insertMap[dbTagValue] = valField.Interface()
			}
		}
	}

	return true, insertBuilder

}
func GetDbTagFromType() {

}
func GetColumns(obj any, removeKeys *[]string) []string {
	columns := []string{}
	reflectVal := reflect.ValueOf(obj)
	reflectVal = reflect.Indirect(reflectVal)
	reflectType := reflectVal.Type()
	for i := 0; i < reflectVal.NumField(); i++ {
		fieldType := reflectType.Field(i)
		if IsRemoveKey(removeKeys, fieldType.Name) {
		} else {
			columns = append(columns, fieldType.Name)
		}
	}
	return columns
}

func IsRemoveKey(removeKeys *[]string, key string) bool {
	for _, v := range *removeKeys {
		if v == key {
			return true
		}
	}
	return false
}
