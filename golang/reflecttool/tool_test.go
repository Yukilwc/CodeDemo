package reflecttool

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestRangeStruct(t *testing.T) {
	data := &EventNotification{}
	RangeStruct(data, func(vField reflect.Value, sField reflect.StructField) {
		// fmt.Printf("\n 循环Range struct value: %+v\n", vField)
		if reflect.TypeOf(time.Time{}) == sField.Type {
			fmt.Printf("\n 循环Range struct field name: %+v\n", sField.Name)
			fmt.Printf("\n 循环Range struct field type name: %+v\n", sField.Type.Name())
			fmt.Printf("\n 循环Range struct field value: %+v\n", vField.Interface())
		} else if reflect.TypeOf(sql.NullTime{}) == sField.Type {
			fmt.Printf("\n 循环Range struct field name: %+v\n", sField.Name)
			fmt.Printf("\n 循环Range struct field type name: %+v\n", sField.Type.Name())
			fmt.Printf("\n 循环Range struct field value: %+v\n", vField.Interface())
		}

	})
}
func TestHasKey(t *testing.T) {
	data := &EventNotification{}
	fmt.Printf("\n Has SolarDate: %+v\n", HasKey(data, "SolarDate"))
	fmt.Printf("\n Has TestDate: %+v\n", HasKey(data, "TestDate"))
}

func TestConvertTime2Int64(t *testing.T) {
	from := &EventNotification{}
	from.SolarDate.Scan(time.Now())
	from.CreateAt = time.Now()
	from.Name = "测试"
	to := &EventNotificationResponse{}
	ConvertTime2Int64(to, from)
	fmt.Printf("\nto: %+v\n", to)
	fmt.Printf("\nto SolarDate: %+v\n", *to.SolarDate)
	fmt.Printf("\nto ChineseDate: %+v\n", to.ChineseDate)
	fmt.Printf("\nto CreateAt: %+v\n", to.CreateAt)
	fmt.Printf("\nto UpdateAt: %+v\n", to.UpdateAt)
	fmt.Printf("\nto DeleteAt: %+v\n", to.DeleteAt)

}

func TestCopierCopyTime(t *testing.T) {
	from := &EventNotification{}
	from.SolarDate.Scan(time.Now())
	from.CreateAt = time.Now()
	from.Name = "测试"
	to := &EventNotificationResponse{}
	CopierCopyTime(to, from)
	fmt.Printf("\nto: %+v\n", to)
	fmt.Printf("\nto SolarDate: %+v\n", to.SolarDate)
	fmt.Printf("\nto ChineseDate: %+v\n", to.ChineseDate)
	fmt.Printf("\nto CreateAt: %+v\n", to.CreateAt)
	fmt.Printf("\nto UpdateAt: %+v\n", to.UpdateAt)
	fmt.Printf("\nto DeleteAt: %+v\n", to.DeleteAt)
	fmt.Printf("\n reflect.TypeOf(int64(0): %+v\n", reflect.TypeOf(int64(0)))
	fmt.Printf("\n reflect.PointerTo(reflect.TypeOf(int64(0)): %+v\n", reflect.PointerTo(reflect.TypeOf(int64(0))))
}

type UserRow struct {
	Id    int64
	Name  string
	Email string
}

func TestUserName(t *testing.T) {
	resRows := &[]*EventNotificationResponse{
		{
			CreateBy: 1,
			UpdateBy: 1,
		},
		{
			CreateBy: 2,
			UpdateBy: 3,
		},
	}
	SetUserName(
		resRows,
		[]string{"CreateBy", "UpdateBy"},
		[]string{"CreateByName", "UpdateByName"},
		func(ids *[]int64) *[]UserRow {
			users := &[]UserRow{
				{
					Id:   1,
					Name: "名字一",
				},
				{
					Id:   2,
					Name: "名字二",
				},
				{
					Id:   3,
					Name: "名字三",
				},
			}
			return users
		},
	)

	for _, v := range *resRows {
		fmt.Printf("\n最终修改后的resRows : %+v\n", v)
	}
}
