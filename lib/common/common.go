package common

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type JSON = map[string]interface{}

func GetBetweenDays(startDate int, endDate int) []string {
	var betweenDays []string
	startYear := startDate / 1e4
	startMonth := (startDate % 1e4)/ 1e2
	startDay := startDate % 1e2
	currentDay := time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)

	endYear := endDate / 1e4
	endMonth := (endDate % 1e4)/ 1e2
	endDay := endDate % 1e2
	finalDay := time.Date(endYear, time.Month(endMonth), endDay, 0, 0, 0, 0, time.UTC)

	between := int(finalDay.Sub(currentDay).Hours() / 24)
	betweenDays = append(betweenDays, strconv.Itoa(startDate))
	for i := 0; i < between; i++ {
		var month string
		var day string
		currentDay = currentDay.AddDate(0, 0, 1)
		m := currentDay.Month()
		d := currentDay.Day()

		if m < 10 {
			month = "0" + strconv.Itoa(int(m))
		} else {
			month = strconv.Itoa(int(m))
		}

		if d < 10 {
			day = "0" + strconv.Itoa(d)
		} else {
			day = strconv.Itoa(d)
		}

		currentDayString := strconv.Itoa(currentDay.Year()) + month + day
		betweenDays = append(betweenDays, currentDayString)
	}

	return betweenDays
}

func GetSparkPVTableName(date string) string{
	return "stat_page_pv_" + date
}

func GetPVMonthlyTableName(date string) string{
	return "pv_monthly_" + date
}

func GetPageID(unParsedUrl string) string {
	unParsedUrl = strings.ReplaceAll(unParsedUrl, "#.*", "")
	parsedUrl, err := url.Parse(unParsedUrl)
	if err != nil {
		panic(err)
	}
	netLoc := parsedUrl.Hostname()
	if len(parsedUrl.Port()) > 0 {
		netLoc = netLoc + parsedUrl.Port()
	}
	var text string
	if len(parsedUrl.Query()) > 0 {
		text = netLoc + parsedUrl.Path + parsedUrl.RawQuery
	} else {
		text = netLoc + parsedUrl.Path
	}
	hashCode := sha1.New()
	hashCode.Write([]byte(text))
	fmt.Println("text: ", text)
	return hex.EncodeToString(hashCode.Sum(nil))
}
