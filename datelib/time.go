package datelib

import (
	"github.com/NubeIO/lib-date/datectl"
	"time"
)

type Time struct {
	DateStamp       time.Time `json:"date_stamp"`
	TimeLocal       string    `json:"time_local"`
	TimeUTC         string    `json:"time_utc"`
	CurrentDay      string    `json:"current_day"`
	CurrentDayUTC   string    `json:"current_day_utc"`
	DateFormatLocal string    `json:"date_format_local"`
	DateFormatUTC   string    `json:"date_format_utc"`
	SystemTimezone  string    `json:"system_timezone"`
}

func (inst *Date) SystemTime() *Time {
	dateCTL := datectl.DateCTL{}
	t := new(Time)
	t.DateStamp = time.Now()
	timeUTC := t.DateStamp.UTC()
	t.TimeLocal = t.DateStamp.Format("15:04:05")
	t.TimeUTC = timeUTC.Format("15:04:05")
	t.CurrentDay = t.DateStamp.Format("Monday")
	t.CurrentDayUTC = timeUTC.Format("Monday")
	t.DateFormatLocal = t.DateStamp.Format("2006-01-02 15:04:05")
	t.DateFormatUTC = timeUTC.Format("2006-01-02 15:04:05")
	zone, _ := dateCTL.GetHardwareTZ()
	t.SystemTimezone = zone
	return t
}
