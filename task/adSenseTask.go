package task

import (
	"DataApi.Go/database/orm"
	"DataApi.Go/lib/common"
	"github.com/jinzhu/gorm"
)

func QueryAdSenseReportList(db *gorm.DB, accountId []string, startDate int, endDate int) []common.JSON{
	caAccountId := common.GetCaAccountIds(accountId)
	reportList := orm.SelectAdSenseReport(db, caAccountId, common.ConvertTime(startDate), common.ConvertTime(endDate))
	return reportList
}

func QueryAdSenseRevenueList(db *gorm.DB, accountId []string, startDate int, endDate int) []common.JSON{
	reportList := orm.SelectAdSenseRevenue(db, accountId, common.ConvertTime(startDate), common.ConvertTime(endDate))
	return reportList
}

func QueryAdSenseDomainList(db *gorm.DB) map[string][]string {
	response := orm.SelectAdSenseDomainMapping(db)
	return response
}
