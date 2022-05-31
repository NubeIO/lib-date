package datelib

import (
	"strings"
)

func clean(s string) string {
	if idx := strings.Index(s, ":"); idx != -1 {
		i := strings.Trim(s[idx:], ":")
		i = strings.Join(strings.Fields(strings.TrimSpace(i)), " ")
		return i
	}
	return s
}

type HardwareClock struct {
	Localtime               string `json:"localtime"`
	UniversalTime           string `json:"utc_time"`
	RTCtime                 string `json:"rtc_time"`
	Timezone                string `json:"timezone"`
	SystemClockSynchronized string `json:"system_clock_synchronized"`
	NTPService              string `json:"ntp_service"`
	RTCInLocalTZ            string `json:"rtc_in_local_tz"`
}

func (inst *Admin) GetHardwareTZ() (string, error) {

	inst.CMD.Commands = Builder("cat", "/etc/timezone")
	res := inst.CMD.RunCommand()
	out := strings.Split(string(res.Out), "\n")
	if len(out) >= 0 {
		return out[0], res.Err
	} else {
		return "", res.Err
	}
}

func (inst *Admin) GetHardwareClock() (HardwareClock, error) {
	inst.CMD.Commands = Builder("timedatectl", "status")
	res := inst.CMD.RunCommand()
	var hc HardwareClock
	if res.Err != nil {
		return hc, res.Err
	}
	var items []string
	list := strings.Split(string(res.Out), "\n")
	for _, s := range list {
		items = append(items, clean(s))
	}
	if len(items) >= 6 {
		hc.Localtime = items[0]
		hc.UniversalTime = items[1]
		hc.RTCtime = items[2]
		hc.Timezone = items[3]
		hc.SystemClockSynchronized = items[4]
		hc.NTPService = items[5]
		hc.RTCInLocalTZ = items[6]
	}
	return hc, nil
}
