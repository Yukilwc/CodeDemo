package timetool

import "time"

func CompareTime(t1 time.Time, t2 time.Time) int {
	now := time.Now()
	dt1 := GetDetailTime(now, t1)
	dt2 := GetDetailTime(now, t2)
	if dt1.After(dt2) {
		return 1
	} else if dt1.Equal(dt2) {
		return 0
	} else {
		return -1
	}
}

func GetDetailTime(dateInfo time.Time, timeInfo time.Time) time.Time {
	t := time.Date(
		dateInfo.Year(),
		dateInfo.Month(),
		dateInfo.Day(),
		timeInfo.Hour(),
		timeInfo.Minute(),
		timeInfo.Second(),
		0,
		dateInfo.Location(),
	)
	return t

}
