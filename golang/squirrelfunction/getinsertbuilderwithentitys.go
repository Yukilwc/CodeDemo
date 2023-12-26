package squirrelfunction

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Masterminds/squirrel"
)

func GetInsertBuilderWithEntitys(entity *[]any) (ok:bool,insertBuilder squirrel.InsertBuilder) {
	if len(entity) == 0 {
		return false,nil
	}
	val := reflect.ValueOf(entity)
	val = reflect.Indirect(val)
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		valField := val.Field(i)
		typeField := typ.Field(i)
		tagValue := typeField.Tag.Get("db")
		switch tagValue {
		case "-":
			continue
		case "":
			continue
		default:
			if strings.Contains(tagValue, ",") {
				tagValue = strings.Split(tagValue, ",")[0]
			}
			tagValue = strings.TrimSpace(tagValue)
			tagValue = fmt.Sprintf("`%s`", tagValue)
			insertMap[tagValue] = valField.Interface()
		}
	}
	insertBuilder = insertBuilder.SetMap(insertMap)
	return insertBuilder

}
