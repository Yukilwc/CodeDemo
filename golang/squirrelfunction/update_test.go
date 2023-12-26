package squirrelfunction

import (
	"fmt"
	"testing"
	"time"

	"github.com/Masterminds/squirrel"
)

func TestGetUpdateBuilderWithEntity(t *testing.T) {
	entity := &EventNotification{}
	entity2 := &EventNotification{}
	// do something to init entity
	entity.UpdateAt = time.Now()
	insertBuilder := squirrel.Insert("mytablename")
	ok, insertBuilder := GetInsertBuilderWithEntitys(
		&[]any{entity, entity2},
		insertBuilder,
		&[]string{"Id", "CreateAt", "UpdateAt"})
	if !ok {
		fmt.Println("数组为空")
		return
	}
	sqlStr, values, err := insertBuilder.ToSql()
	fmt.Printf("\n sqlStr: %+v\n", sqlStr)
	fmt.Printf("\n values: %+v\n", values)
	fmt.Printf("\n to sql err: %+v\n", err)

}
