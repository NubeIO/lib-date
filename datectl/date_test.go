package datectl

import (
	"fmt"
	pprint "github.com/NubeIO/lib-date/print"
	"testing"
	"time"
)

func TestDateCTL_GetHardwareClock(t *testing.T) {
	sys := New(&DateCTL{})
	time_, err := sys.GetHardwareClock()
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJOSN(time_)
}

func TestDateCTL_SetSystemTime(t *testing.T) {
	sys := New(&DateCTL{})
	fmt.Println(time.Now().String())
	err := sys.SetSystemTime("2015-11-24 08:10:40")
	fmt.Println(err)
	if err != nil {
		return
	}
}
func TestDateCTL_ntpEnable(t *testing.T) {
	sys := New(&DateCTL{})
	err, _ := sys.ntpEnable(false)
	fmt.Println(err)
	if err != nil {
		return
	}
}
