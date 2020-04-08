package common

import (
	"crypto/sha1"
	"encoding/hex"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type JSON = map[string]interface{}

func GetBetweenDays(startDate int, endDate int, convertDate bool) []string {
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
	if convertDate == true {
		betweenDays = append(betweenDays, ConvertTime(startDate))
	} else {
		betweenDays = append(betweenDays, strconv.Itoa(startDate))
	}
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
		var currentDayString string
		if convertDate == true {
			currentDayString = strconv.Itoa(currentDay.Year()) + "-" + month + "-" + day
		} else {
			currentDayString = strconv.Itoa(currentDay.Year()) + month + day
		}
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
	return hex.EncodeToString(hashCode.Sum(nil))
}

func ConvertTime(date int) string {
	year := date / 1e4
	month := (date % 1e4) / 1e2
	day := date % 1e2

	var m, d string
	if month < 10 {
		m = "0" + strconv.Itoa(int(month))
	} else {
		m = strconv.Itoa(int(month))
	}

	if day < 10 {
		d = "0" + strconv.Itoa(day)
	} else {
		d = strconv.Itoa(day)
	}
	return strconv.Itoa(year) + "-" + m + "-" + d
}

func Unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func GetCaAccountIds(accountIds []string) []string{
	var caAccountIds []string
	for _, id := range accountIds {
		caAccountIds = append(caAccountIds, "ca-" + id)
	}
	return caAccountIds
}

func GetZiUrl(author string, articleId string) string{
	return "https://zi.media/" + "@" + author + "/" + "post" + "/" + articleId
}

func Difference(slice1 []string, slice2 []string) []string {
	var diff []string
	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}
	return diff
}
