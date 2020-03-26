package orm

import (
	"DataApi.Go/database/models/AdSense"
	"DataApi.Go/lib/common"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type AdSenseReportDaily = AdSense.AdSenseReportDaily
type AdSenseRevenue = AdSense.AdSenseRevenue
type AdSenseDomain = AdSense.AdSenseDomain

func QueryAdSenseReport(db *gorm.DB, accountId []string, startDate string, endDate string) []common.JSON {
	table := "adsense_report_daily"
	rows, err := db.Table(table).Model(&AdSenseReportDaily{}).Where("ad_client_id IN (?) AND date BETWEEN ? AND ?", accountId, startDate, endDate).Rows()
	var rowsList []common.JSON
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var adSenseReportDaily AdSenseReportDaily
		err := db.ScanRows(rows, &adSenseReportDaily)
		if err != nil {
			fmt.Println(err)
		}
		rowsList = append(rowsList, common.JSON{
			"account_id": strings.Replace(adSenseReportDaily.AdClientId, "ca-", "", -1),
			"ad_client_id": adSenseReportDaily.AdClientId,
			"ad_exchange_clicks": adSenseReportDaily.AdExchangeClicks,
			"ad_exchange_impression_rpm": adSenseReportDaily.AdExchangeImpressionRpm,
			"ad_exchange_impressions": adSenseReportDaily.AdExchangeImpressions,
			"clicks": adSenseReportDaily.Clicks,
			"customer_ad_exchange_estimated_revenue": adSenseReportDaily.CustomerAdExchangeEstimatedRevenue,
			"date": adSenseReportDaily.Date,
			"domain_name": adSenseReportDaily.DomainName,
			"earnings": adSenseReportDaily.Earnings,
			"impression_rpm": adSenseReportDaily.ImpressionRpm,
			"impressions": adSenseReportDaily.MatchedAdRequests,
			"network_code": adSenseReportDaily.NetworkCode,
			"page_rpm": adSenseReportDaily.PageRpm,
			"page_views": adSenseReportDaily.PageViews,
		})
	}
	fmt.Println(rowsList)
	return rowsList
}

func QueryAdSenseRevenue(db *gorm.DB, accountId []string, startDate string, endDate string) []common.JSON {
	table := "adsense_report_daily"
	rows, err := db.Table(table).Model(&AdSenseRevenue{}).Select("account_id, SUM(customer_ad_exchange_estimated_revenue) AS customer_ad_exchange_estimated_revenue").Where("account_id IN (?) AND date BETWEEN ? AND ?", accountId, startDate, endDate).Group("account_id").Rows()
	var rowsList []common.JSON
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var adSenseRevenue AdSenseRevenue
		err := db.ScanRows(rows, &adSenseRevenue)
		if err != nil {
			fmt.Println(err)
		}
		rowsList = append(rowsList, common.JSON{
			"account_id": adSenseRevenue.AccountId,
			"customer_ad_exchange_estimated_revenue": adSenseRevenue.CustomerAdExchangeEstimatedRevenue,
		})
	}
	fmt.Println(rowsList)
	return rowsList
}

func QueryAdSenseDomain(db *gorm.DB) map[string][]string {
	table := "adsense_report_daily"
	rows, err := db.Table(table).Model(&AdSenseDomain{}).Select("ad_client_id, GROUP_CONCAT(distinct domain_name) as domain_name").Where("domain_name is not null and domain_name != ?", "webcaches").Group("ad_client_id").Order("ad_client_id").Rows()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	rowsMap := make(map[string][]string)
	for rows.Next() {
		var adSenseDomain AdSenseDomain
		err := db.ScanRows(rows, &adSenseDomain)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		domainSlice := strings.Split(adSenseDomain.DomainName, ",")
		accountId := strings.Replace(adSenseDomain.AdClientId, "ca-", "", -1)
		rowsMap[accountId] = domainSlice
	}
	return rowsMap
}

func QueryAdSenseReportList(db *gorm.DB, accountId []string, startDate int, endDate int) []common.JSON{
	caAccountId := common.GetCaAccountIds(accountId)
	reportList := QueryAdSenseReport(db, caAccountId, common.ConvertTime(startDate), common.ConvertTime(endDate))
	return reportList
}

func QueryAdSenseRevenueList(db *gorm.DB, accountId []string, startDate int, endDate int) []common.JSON{
	reportList := QueryAdSenseRevenue(db, accountId, common.ConvertTime(startDate), common.ConvertTime(endDate))
	return reportList
}

func QueryAdSenseDomainList(db *gorm.DB) map[string][]string {
	response := QueryAdSenseDomain(db)
	return response
}