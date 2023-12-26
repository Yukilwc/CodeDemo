package squirrelfunction

import (
	"fmt"
	"testing"
)

func TestGetColumns(t *testing.T) {
	entity := &EventNotification{}
	keys, dbs := GetColumns(entity, &[]string{"Id", "CreateAt", "UpdateAt"})
	fmt.Printf("\n columns keys: %+v\n", keys)
	fmt.Printf("\n columns dbs: %+v\n", dbs)
}
