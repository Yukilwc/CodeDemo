package crontool

import (
	"fmt"
	"time"

	"github.com/liujiawm/gocalendar"
)

// 指定的日期，是否是指定的农历日期
func DayIsLunarDate(day time.Time, dateMatchingLunarDate time.Time) bool {
	// 把这个日期转回农历
	c := gocalendar.DefaultCalendar()
	lunarDate := c.GregorianToLunar(dateMatchingLunarDate.Year(), int(dateMatchingLunarDate.Month()), dateMatchingLunarDate.Day())
	fmt.Printf("\n lunarDate: %+v\n", lunarDate)
	fmt.Printf("\n lunar date month,day: %+v,%+v,%+v\n", lunarDate.Month, lunarDate.Day, lunarDate.LeapStr)
	dayLunarDate := c.GregorianToLunar(day.Year(), int(day.Month()), day.Day())
	fmt.Printf("\n specific day lunar date: %+v\n", dayLunarDate)
	fmt.Printf("\n specific day lunar month day: %+v,%+v,%+v\n", dayLunarDate.Month, dayLunarDate.Day, dayLunarDate.LeapStr)
	if lunarDate.Month == dayLunarDate.Day && lunarDate.Day == dayLunarDate.Day && lunarDate.LeapStr == dayLunarDate.LeapStr {
		// 月份，日，闰都一致，才是同一个农历
		return true
	} else {
		return false
	}
}

// 检查时间是否在指定日期的前三天
func CheckTimeWithinTimeRange(checkDate time.Time, setDate time.Time) (bool, int64) {
	leftLimit := setDate.AddDate(0, 0, -3)
	leftLimit = GetDate(leftLimit)
	rightLimit := GetDate(setDate)
	checkDate = GetDate(checkDate)
	fmt.Printf("\n 限定的左边界: %+v\n", leftLimit)
	fmt.Printf("\n 限定的右边界: %+v\n", rightLimit)
	fmt.Printf("\n 需要检查的日期: %+v\n", checkDate)
	if checkDate.Before(leftLimit) {
		return false, 0
	} else if checkDate.After(rightLimit) {
		return false, 0
	} else {
		// 计算下差距
		offset := rightLimit.Sub(checkDate).Hours() / 24
		fmt.Printf("\n 距离最后时间的日期: %+v\n", offset)
		return true, int64(offset)
	}
}

func GetDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func GetMonthDay() {

}
