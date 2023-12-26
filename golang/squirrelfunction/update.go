package squirrelfunction

import (
	"reflect"

	"github.com/Masterminds/squirrel"
)

func GetUpdateBuilderWithEntity(entity any, updateBuilder squirrel.UpdateBuilder, removeKeys *[]string) (bool, squirrel.UpdateBuilder) {
	keyColumns, _ := GetColumns(entity, removeKeys)
	reflectVal := reflect.ValueOf(entity)
	reflectVal = reflect.Indirect(reflectVal)
	m := make(map[string]any)
	for _, key := range *keyColumns {
		fieldVal := reflectVal.FieldByName(key)
		m[key] = fieldVal.Interface()
	}
	updateBuilder = updateBuilder.SetMap(m)
	return true, updateBuilder
}
