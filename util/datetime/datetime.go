package datetime

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Date and Time now
var DateTimeNow = time.Now()

const (
	// DefaultTimeZone Default Time Zone Thailand
	DefaultTimeZone = "Asia/Bangkok"
	// set default DateTime format
	Default_DateTime_Formate = "2006-01-02 15:04:05"
	// YYYYmmddhisFormatDatetime DataTime format YearMonthDayHourMinutesSecond
	YYYYmmddHHmmSSFormatDatetime = "20060102150405"
	// YYYY-MM-DD: 2022-03-23
	YYYYMMDD = "2006-01-02"
	// 24h hh:mm:ss: 14:23:20
	HHMMSS24h = "15:04:05"
	// 12h hh:mm:ss: 2:23:20 PM
	HHMMSS12h = "3:04:05 PM"
	// text date: March 23, 2022
	TextDate = "January 2, 2006"
	// text date with weekday: Wednesday, March 23, 2022
	TextDateWithWeekday = "Monday, January 2, 2006"
	// abbreviated text date: Mar 23 Wed
	AbbrTextDate = "Jan 2 Mon"
)

// GetDateTimeCountryZone function to get date time from country zone format 2018-01-01 00:00:00
func GetDateTimeCountryZone(strDateTime string) time.Time {
	t, _ := time.Parse(Default_DateTime_Formate, strDateTime)
	year, month, day := t.Date()
	hour, min, second := t.Clock()
	nsec := t.Nanosecond()
	loc, _ := time.LoadLocation(DefaultTimeZone)

	//set timezone,
	retTime := time.Date(year, month, day, hour, min, second, nsec, loc)
	return retTime
}

func GetCurrentDateTimeNano() string {
	return GetCurrentDateTimeNow_to_String("2006-01-02 15:04:05.000000000")
}

// GetCurrentTime get current DateTime
func GetCurrentDateTimeNow_to_String(format string) string {
	if strings.TrimSpace(format) == "" {
		format = Default_DateTime_Formate
	}
	return time.Now().Format(format)
}

func GetTimeZone() string {
	loc, _ := time.LoadLocation(DefaultTimeZone)
	localDateTime, err := time.ParseInLocation(Default_DateTime_Formate, DateTimeNow.Format(Default_DateTime_Formate), loc)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	fmt.Println(localDateTime)
	return localDateTime.Format(Default_DateTime_Formate)
}
