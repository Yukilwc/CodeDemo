package squirrelfunction

// 此版本弃用，改成差多多行的版本
import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Masterminds/squirrel"
)

func GetInsertWithEntity(insertBuilder squirrel.InsertBuilder, entity any) squirrel.InsertBuilder {
	insertMap := make(map[string]any)
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
