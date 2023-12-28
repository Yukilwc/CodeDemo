package squirrelfunction

import (
	"fmt"
	"reflect"

	"github.com/Masterminds/squirrel"
)

func GetInsertBuilderWithEntitys(entitys *[]any, insertBuilder squirrel.InsertBuilder, removeKeys *[]string) (squirrel.InsertBuilder, bool) {

	if len(*entitys) == 0 {
		return insertBuilder, false
	}
	keyColumns, dbColumns := GetColumns((*entitys)[0], removeKeys)
	insertBuilder = insertBuilder.Columns(*dbColumns...)
	for _, entity := range *entitys {
		reflectVal := reflect.ValueOf(entity)
		reflectVal = reflect.Indirect(reflectVal)
		values := []any{}
		fmt.Printf("\n range entity: %+v\n", entity)
		for _, key := range *keyColumns {
			fieldVal := reflectVal.FieldByName(key)
			// if fieldVal.IsValid() {
			values = append(values, fieldVal.Interface())
			// }
		}
		fmt.Printf("\n 一个entity的values: %+v\n", values)
		// 插入values
		insertBuilder = insertBuilder.Values(values...)

	}
	return insertBuilder, true

}
