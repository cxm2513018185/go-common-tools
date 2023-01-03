package timer

import (
	"fmt"
	"testing"
)

func TestGetNowTime(t *testing.T) {
	nowTime := GetNowTime()
	fmt.Println(nowTime)
}
