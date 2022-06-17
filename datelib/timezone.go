package datelib

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func (inst *Date) GetTimeZoneList() ([]string, error) {
	inst.CMD.Commands = Builder("timedatectl", "list-timezones")
	res := inst.CMD.RunCommand()
	if res.Err != nil {
		return nil, res.Err
	}
	list := strings.Split(string(res.OutByte), "\n")
	return list, nil
}

// UpdateTimezone sets the current machine's timezone to the given timezone
func (inst *Date) UpdateTimezone(newZone string) error {
	list, err := inst.GetTimeZoneList()
	if err != nil {
		return err
	}
	var matchZone bool
	for _, zone := range list {
		if zone == newZone {
			matchZone = true
		}
	}
	if !matchZone {
		return errors.New("incorrect zone passed in try, Australia/Sydney")
	}
	inst.CMD.Commands = Builder("timedatectl", "set-timezone", strings.TrimSpace(newZone))
	res := inst.CMD.RunCommand()
	if res.Err != nil {
		return res.Err
	}
	return nil
}

func (inst *Date) SetSystemTime(date string) error {
	layout := "Mon, 02 Jan 2006 15:04:05 MST"
	// parse time
	t, err := time.Parse(layout, date)

	if err != nil {
		return fmt.Errorf("could not parse date  %s", err)
	}
	// set system time
	inst.CMD.Commands = Builder("date", "-s", fmt.Sprintf("@%d", t.Unix()))
	res := inst.CMD.RunCommand()
	if res.Err != nil {
		return res.Err
	}

	return nil
}
