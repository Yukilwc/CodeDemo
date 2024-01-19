package timetool

import (
	"fmt"
	"testing"
	"time"
)

func TestCompareTime(t *testing.T) {
	now := time.Now()
	now2 := now.Add(-time.Second * 60)
	res := CompareTime(now, now2)
	fmt.Printf("\n res: %+v\n", res)
}
