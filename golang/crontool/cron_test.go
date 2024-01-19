package crontool

import (
	"fmt"
	"testing"
	"time"
)

func TestDayIsLunarDate(t *testing.T) {
	ti, err := time.Parse("2006-01-02", "2023-04-13")
	if err != nil {
		fmt.Printf("\n err: %+v\n", err)
		return
	}
	next := DayIsLunarDate(ti, time.Now())
	fmt.Printf("\n next:: %+v\n", next)
}

func TestCheckTimeWithinTimeRange(t *testing.T) {
	CheckTimeWithinTimeRange(time.Now().AddDate(0, 0, 3), time.Now().AddDate(0, 0, 3))

}

func TestGetDate(t *testing.T) {
	fmt.Printf("\n getDate now: %+v\n", GetDate(time.Now()))
}
