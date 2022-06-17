package datelib

import (
	"fmt"
	"testing"
)

func TestDate_UpdateTimezone(t *testing.T) {
	host := &Date{}
	run := New(host)

	err := run.UpdateTimezone("Australia/Sydney")
	fmt.Println(err)
	if err != nil {
		return
	}

	err = run.SetSystemTime("Mon, 02 Jan 2006 15:04:05 AEST")
	fmt.Println(err)
	if err != nil {
		return
	}

}
