package squirrelfunction

import (
	"reflect"

	"github.com/Masterminds/squirrel"
)

func GetUpdateBuilderWithEntity(entity any, updateBuilder squirrel.UpdateBuilder, removeKeys *[]string) squirrel.UpdateBuilder {
	keyColumns, dbColumns := GetColumns(entity, removeKeys)
	reflectVal := reflect.ValueOf(entity)
	reflectVal = reflect.Indirect(reflectVal)
	m := make(map[string]any)
	for i, key := range *keyColumns {
		fieldVal := reflectVal.FieldByName(key)
		m[(*dbColumns)[i]] = fieldVal.Interface()
	}
	updateBuilder = updateBuilder.SetMap(m)
	return updateBuilder
}
