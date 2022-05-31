package datelib

import (
	"strings"
)

func (inst *Admin) GetTimeZoneList() ([]string, error) {
	inst.CMD.Commands = Builder("timedatectl", "list-timezones")
	res := inst.CMD.RunCommand()
	if res.Err != nil {
		return nil, res.Err
	}
	list := strings.Split(string(res.OutByte), "\n")
	return list, nil
}
