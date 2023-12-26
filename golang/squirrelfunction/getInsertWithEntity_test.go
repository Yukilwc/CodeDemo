package squirrelfunction

import (
	"fmt"
	"testing"
	"time"

	"github.com/Masterminds/squirrel"
)

func TestGetInsertWithEntity(t *testing.T) {
	entity := &EventNotification{}
	entity.UpdateAt = time.Now()
	// do something to init entity
	insertBuilder := squirrel.Insert("mytablename")
	insertBuilder = GetInsertWithEntity(insertBuilder, entity)
	sqlStr, values, err := insertBuilder.ToSql()
	fmt.Printf("\n sqlStr: %+v\n", sqlStr)
	fmt.Printf("\n values: %+v\n", values)
	fmt.Printf("\n to sql err: %+v\n", err)

}
