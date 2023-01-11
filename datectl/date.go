package datectl

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
	"time"
)

/*
https://gist.github.com/mutin-sa/eea1c396b1e610a2da1e5550d94b0453

How to change NTP server:
sudo nano /etc/systemd/timesyncd.conf
FallbackNTP=time.windows.com
sudo systemctl restart systemd-timesyncd
systemctl status systemd-timesyncd.service
timedatectl
*/

func (inst *DateCTL) GetHardwareTZ() (string, error) {
	cmd := exec.Command("cat", "/etc/timezone")
	output, err := cmd.Output()
	cleanCommand(string(output), cmd, err, debug)
	if err != nil {
		return "", err
	}
	out := strings.Split(string(output), "\n")
	if len(out) >= 0 {
		return out[0], nil
	} else {
		return "", errors.New("failed to find timezone")
	}
}

func (inst *DateCTL) GetTimeZoneList() ([]string, error) {
	cmd := exec.Command("timedatectl", "list-timezones")
	output, err := cmd.Output()
	cleanCommand(string(output), cmd, err, debug)
	if err != nil {
		return nil, err
	}
	var out []string
	list := strings.Split(string(output), "\n")
	for _, s := range list {
		if s != "" {
			out = append(out, s)
		}
	}
	return out, nil
}

// UpdateTimezone sets the current machine's timezone to the given timezone
func (inst *DateCTL) UpdateTimezone(newZone string) error {
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
	cmd := exec.Command("timedatectl", "set-timezone", strings.TrimSpace(newZone))
	output, err := cmd.Output()
	cleanCommand(string(output), cmd, err, debug)
	if err != nil {
		return err
	}
	return nil
}

//SetSystemTime sudo date -s '2015-11-23 08:10:40'
func (inst *DateCTL) SetSystemTime(dateTime string) error {
	layout := "2006-01-02 15:04:05"
	// parse time
	t, err := time.Parse(layout, dateTime)
	if err != nil {
		return fmt.Errorf("could not parse date  %s", err)
	}
	log.Infof("set time to %s", t.String())
	cmd := exec.Command("sudo", "date", "-s", dateTime)
	output, err := cmd.Output()
	cleanCommand(string(output), cmd, err, debug)
	if err != nil {
		return err
	}
	return nil
}

//SetSystemTimeCTL timedatectl set-time '2015-11-23 08:10:40'
func (inst *DateCTL) SetSystemTimeCTL(dateTime string) error {
	layout := "2006-01-02 15:04:05"
	// parse time
	t, err := time.Parse(layout, dateTime)
	if err != nil {
		return fmt.Errorf("could not parse date  %s", err)
	}
	log.Infof("set time to %s", t.String())
	cmd := exec.Command("timedatectl", "set-time", dateTime)
	output, err := cmd.Output()
	cleanCommand(string(output), cmd, err, debug)
	if err != nil {
		return err
	}
	return nil
}
