package datelib

import (
	"fmt"
	"testing"
)

func TestAdmin_GetHardwareClock(t *testing.T) {
	host := &Admin{}
	run := New(host)
	run.GetTimeZoneList()

	out, err := run.SystemTime()
	fmt.Println(err, out)
}
