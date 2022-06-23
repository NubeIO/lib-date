package datelib

import (
	"fmt"
	"testing"
)

func TestAdmin_GetHardwareClock(t *testing.T) {
	host := &Date{}
	run := New(host)
	run.GetTimeZoneList()

	//out, err := run.SystemTimeHardware()
	//fmt.Println(err, out)

	out2 := run.SystemTime()
	fmt.Println(out2)
}
