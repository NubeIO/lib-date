package datectl

import (
	"os/exec"
)

func (inst *DateCTL) ntpEnable(enable bool) (*Message, error) {
	cmd := exec.Command("timedatectl", "set-ntp", "false")
	if enable {
		cmd = exec.Command("timedatectl", "set-ntp", "true")
	}
	output, err := cmd.Output()
	cleanCommand(string(output), cmd, err, debug)
	if err != nil {
		return nil, err
	}
	return &Message{Message: "ok"}, err
}

func (inst *DateCTL) NTPDisable() (*Message, error) {
	return inst.ntpEnable(false)
}

func (inst *DateCTL) NTPEnable() (*Message, error) {
	return inst.ntpEnable(true)
}

const (
	timeSyncConfPath = "/etc/systemd/timesyncd.conf"
)

// TimeSyncConfig Json request
type TimeSyncConfig struct {
	NTP                []string `json:"NTP"`
	FallbackNTP        []string `json:"fallback_ntp"`
	RootDistanceMaxSec string   `json:"root_distance_max_sec"`
	PollIntervalMinSec string   `json:"poll_interval_min_sec"`
	PollIntervalMaxSec string   `json:"poll_interval_max_sec"`
}

func (inst *DateCTL) GenerateTimeSyncConfig(body *TimeSyncConfig) string {
	if body == nil {
		return ""
	}
	conf := "[Time]\n"
	ntpConf := "NTP="
	for _, s := range body.NTP {
		ntpConf += s + " "
	}
	conf += ntpConf + "\n"
	if len(body.FallbackNTP) > 0 {
		FallbackNTP := "FallbackNTP="
		for _, s := range body.FallbackNTP {
			FallbackNTP += s + " "
		}
		conf += FallbackNTP + "\n"
	}

	if body.RootDistanceMaxSec != "" {
		conf += "RootDistanceMaxSec=" + body.RootDistanceMaxSec + "\n"
	}

	if body.PollIntervalMinSec != "" {
		conf += "PollIntervalMinSec=" + body.PollIntervalMinSec + "\n"
	}

	if body.PollIntervalMaxSec != "" {
		conf += "PollIntervalMaxSec=" + body.PollIntervalMaxSec + "\n"
	}
	return conf
}
