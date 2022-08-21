package system

import (
	"os/exec"
	"strings"
)

type HardwareClock struct {
	Localtime               string `json:"localtime"`
	UniversalTime           string `json:"utc_time"`
	RTCtime                 string `json:"rtc_time"`
	Timezone                string `json:"timezone"`
	SystemClockSynchronized string `json:"system_clock_synchronized"`
	NTPService              string `json:"ntp_service"`
	RTCInLocalTZ            string `json:"rtc_in_local_tz"`
}

func (inst *DateCTL) GetHardwareClock() (*HardwareClock, error) {
	hc := &HardwareClock{}
	cmd := exec.Command("timedatectl", "status")
	output, err := cmd.Output()
	cleanCommand(string(output), cmd, err, debug)
	if err != nil {
		return nil, err
	}
	var items []string
	list := strings.Split(string(output), "\n")
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

func clean(s string) string {
	if idx := strings.Index(s, ":"); idx != -1 {
		i := strings.Trim(s[idx:], ":")
		i = strings.Join(strings.Fields(strings.TrimSpace(i)), " ")
		return i
	}
	return s
}
