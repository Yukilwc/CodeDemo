package reflecttool

import (
	"database/sql"
	"reflect"
	"time"

	"github.com/jinzhu/copier"
)

// 遍历结构体实例的属性
func RangeStruct(value any, callback func(val reflect.Value, sf reflect.StructField)) {
	rVal := reflect.ValueOf(value)
	rVal = reflect.Indirect(rVal)
	rType := rVal.Type()
	for i := 0; i < rVal.NumField(); i++ {
		typeField := rType.Field(i)
		valueField := rVal.Field(i)
		callback(valueField, typeField)
	}
}

// 结构体实例是否包含某个key
func HasKey(value any, key string) bool {
	rVal := reflect.ValueOf(value)
	rVal = reflect.Indirect(rVal)
	rType := rVal.Type()
	_, ok := rType.FieldByName(key)
	return ok
}
func GetStructFieldByKey(value any, key string) (reflect.StructField, bool) {
	rVal := reflect.ValueOf(value)
	rVal = reflect.Indirect(rVal)
	rType := rVal.Type()
	f, ok := rType.FieldByName(key)
	return f, ok
}
func GetFieldByKey(value any, key string) (reflect.StructField, reflect.Value, bool) {
	rVal := reflect.ValueOf(value)
	rVal = reflect.Indirect(rVal)
	rType := rVal.Type()
	structField, ok := rType.FieldByName(key)
	valueField := rVal.FieldByName(key)
	return structField, valueField, ok
}

// 属性是否是Time或者是NullTime
func IsTimeType(sf reflect.StructField) bool {
	if (sf.Type == reflect.TypeOf(time.Time{})) {
		return true
	} else if (sf.Type == reflect.TypeOf(sql.NullTime{})) {
		return true
	} else {
		return false
	}
}
func IsInt64Ptr(sf reflect.StructField) bool {
	if sf.Type.Kind() == reflect.Pointer && sf.Type.Elem().Kind() == reflect.Int64 {
		return true
	} else {
		return false
	}
}
func IsInt64(sf reflect.StructField) bool {
	if sf.Type.Kind() == reflect.Int64 {
		return true
	} else {
		return false
	}
}

// 自动转换日期字段为int64
func ConvertTime2Int64(to any, from any) {
	// keyList := []string{}
	rv := reflect.ValueOf(from)
	rv = reflect.Indirect(rv)
	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		valueField := rv.Field(i)
		structField := rt.Field(i)
		if structField.Type == reflect.TypeOf(time.Time{}) {
			// 是个时间类型
			if t, ok := valueField.Interface().(time.Time); ok {
				// fmt.Printf("\n from中字段是时间类型: %+v\n", structField.Name)
				if !t.IsZero() {
					// fmt.Printf("\n from中时间值不是零值: %+v\n", valueField.Interface())
					// 时间不是零值
					reflectValueFromTo := reflect.ValueOf(to)
					reflectValueFromTo = reflect.Indirect(reflectValueFromTo)
					reflectTypeFromTo := reflectValueFromTo.Type()
					structFieldFromTo, ok := reflectTypeFromTo.FieldByName(structField.Name)
					keyValueFieldFromTo := reflectValueFromTo.FieldByName(structField.Name)
					if ok {
						// fmt.Printf("\n 从to中查找key的struct field: %+v\n", structFieldFromTo)
						// to中存在这个属性
						// 如果此属性是指针，且指向int64
						if structFieldFromTo.Type == reflect.TypeOf(int64(0)) {
							// fmt.Printf("\n to中的目标字段是int64类型:%+v\n", keyValueFieldFromTo)
							keyValueFieldFromTo.SetInt(t.UnixMilli())
						}
						if structFieldFromTo.Type == reflect.PointerTo(reflect.TypeOf(int64(0))) {
							// 从指针获取对应的值
							keyValueFieldFromTo = keyValueFieldFromTo.Elem()
							keyValueFieldFromTo.SetInt(t.UnixMilli())
						}

					}

				}
			}
		} else if structField.Type == reflect.TypeOf(sql.NullTime{}) {
			if t, ok := valueField.Interface().(sql.NullTime); ok {
				// fmt.Printf("\n from中字段是Null时间类型: %+v\n", structField.Name)
				if t.Valid && !t.Time.IsZero() {
					// fmt.Printf("\n from中时间值不是零值: %+v\n", valueField.Interface())
					// 时间不是零值
					reflectValueFromTo := reflect.ValueOf(to)
					reflectValueFromTo = reflect.Indirect(reflectValueFromTo)
					reflectTypeFromTo := reflectValueFromTo.Type()
					structFieldFromTo, ok := reflectTypeFromTo.FieldByName(structField.Name)
					keyValueFieldFromTo := reflectValueFromTo.FieldByName(structField.Name)
					if ok {
						// fmt.Printf("\n 从to中查找key的struct field: %+v\n", structFieldFromTo)
						// to中存在这个属性
						// 如果此属性是指针，且指向int64
						if structFieldFromTo.Type == reflect.TypeOf(int64(0)) {
							// fmt.Printf("\n to中的目标字段是int64类型:%+v\n", keyValueFieldFromTo)
							keyValueFieldFromTo.SetInt(t.Time.UnixMilli())
						}
						if structFieldFromTo.Type == reflect.PointerTo(reflect.TypeOf(int64(0))) {
							// fmt.Printf("\n to中的目标字段是int64 pointer类型:%+v\n", keyValueFieldFromTo)
							// 从指针获取对应的值
							stamp := t.Time.UnixMilli()
							// keyValueFieldFromTo = keyValueFieldFromTo.Elem()
							keyValueFieldFromTo.Set(reflect.ValueOf(&stamp))
						}

					}

				}
			}
		}
	}
}

// 自动转换日期的copier
func CopierCopyTime(to any, from any) {
	timeConverter := copier.TypeConverter{
		SrcType: time.Time{},
		DstType: int64(0),
		Fn: func(src interface{}) (interface{}, error) {
			// fmt.Printf("\n 开始进入time converter: \n")
			if v, ok := src.(time.Time); ok {
				if v.IsZero() {
					return int64(0), nil
				} else {
					return v.UnixMilli(), nil
				}
			} else {
				return int64(0), nil
			}
		},
	}
	// var intPtr int64 = 0
	// timePtrConverter := copier.TypeConverter{
	// 	SrcType: time.Time{},
	// 	DstType: &intPtr,
	// 	Fn: func(src interface{}) (interface{}, error) {
	// 		var zeroValue int64 = 0
	// 		if v, ok := src.(time.Time); ok {
	// 			if v.IsZero() {
	// 				return &zeroValue, nil
	// 			} else {
	// 				var timestamp = v.UnixMilli()
	// 				return &timestamp, nil
	// 			}
	// 		} else {
	// 			return &zeroValue, nil
	// 		}
	// 	},
	// }

	nullTimeConverter := copier.TypeConverter{
		SrcType: sql.NullTime{},
		DstType: int64(0),
		Fn: func(src interface{}) (interface{}, error) {
			// fmt.Printf("\n 开始进入null time converter: \n")
			if v, ok := src.(sql.NullTime); ok {
				if v.Valid {
					return int64(0), nil
				} else {
					return v.Time.UnixMilli(), nil
				}
			} else {
				return int64(0), nil
			}
		},
	}
	var dstTypeInt int64 = 0
	nullTimePtrConverter := copier.TypeConverter{
		SrcType: sql.NullTime{},
		DstType: &dstTypeInt,
		Fn: func(src interface{}) (interface{}, error) {
			// fmt.Printf("\n 开始进入null time ptr converter: \n")
			if v, ok := src.(sql.NullTime); ok {
				if v.Valid {
					r := v.Time.UnixMilli()
					return &r, nil
				} else {
					return nil, nil
				}
			} else {
				return nil, nil
			}
		},
	}

	copier.CopyWithOption(to, from, copier.Option{
		IgnoreEmpty:      false,
		CaseSensitive:    false,
		DeepCopy:         false,
		FieldNameMapping: []copier.FieldNameMapping{},
		Converters:       []copier.TypeConverter{timeConverter, nullTimeConverter, nullTimePtrConverter},
	})
}

// 自动填充用户名

func SetUserName[T any, U any](
	resRows *[]T,
	idKeys []string,
	nameKeys []string,
	getUsers func(ids *[]int64) *[]U,
) {
	// 从实体中，获取对应的id
	idList := []int64{}
	for _, resRow := range *resRows {
		for _, idKey := range idKeys {
			// 从中取出id值
			reflectVal := reflect.ValueOf(resRow)
			reflectVal = reflect.Indirect(reflectVal)
			valueField := reflectVal.FieldByName(idKey)
			// fmt.Printf("\n 查找的idKey: %+v\n", idKey)
			// fmt.Printf("\n 查找的idKey的结果: %+v\n", valueField)

			switch v := valueField.Interface().(type) {
			case int64:
				idList = append(idList, v)
			case sql.NullInt64:
				if v.Valid {
					idList = append(idList, v.Int64)
				}

			}
		}
	}

	// fmt.Printf("\n 从实体数组中生成的idList: %+v\n", idList)
	// 去重idList
	deduplicateList := []int64{}
	m := make(map[int64]struct{})
	for _, v := range idList {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			deduplicateList = append(deduplicateList, v)
		}
	}
	// fmt.Printf("\n 去重后的id数组: %+v\n", deduplicateList)
	// 通过这组id，查询到一组用户，这组用户，带有id和name的信息
	users := getUsers(&deduplicateList)
	// 转换下形态
	userIdMap := make(map[int64]any)
	for _, user := range *users {
		_, val, ok := GetFieldByKey(user, "Id")
		if ok {
			id, _ := val.Interface().(int64)
			userIdMap[id] = user
		}
	}
	// fmt.Printf("\n 生成id user的map: %+v\n", userIdMap)
	// 遍历返回rows，对齐内部的指定name字段，进行填充
	for _, resRow := range *resRows {
		for i, idKey := range idKeys {
			nameKey := nameKeys[i]
			_, idValue, _ := GetFieldByKey(resRow, idKey)
			_, nameValue, _ := GetFieldByKey(resRow, nameKey)
			// 拿到id
			id, _ := idValue.Interface().(int64)
			// 从id找到map的user
			user := userIdMap[id]
			// 从user中取出name
			_, userNameVal, _ := GetFieldByKey(user, "Name")
			name := userNameVal.Interface().(string)
			// 写入row中
			nameValue.SetString(name)
		}
	}
}
