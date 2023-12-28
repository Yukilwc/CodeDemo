package squirrelfunction

import (
	"fmt"
	"testing"
	"time"

	"github.com/Masterminds/squirrel"
)

func TestGetUpdateBuilderWithEntity(t *testing.T) {
	entity := &EventNotification{}
	// do something to init entity
	entity.UpdateAt = time.Now()
	builder := squirrel.Update("mytablename")
	builder = GetUpdateBuilderWithEntity(
		entity,
		builder,
		&[]string{"Id", "CreateAt", "UpdateAt", "CreateBy"})

	sqlStr, values, err := builder.ToSql()
	fmt.Printf("\n sqlStr: %+v\n", sqlStr)
	fmt.Printf("\n values: %+v\n", values)
	fmt.Printf("\n to sql err: %+v\n", err)

}
