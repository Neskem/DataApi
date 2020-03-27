package YPA

import "DataApi.Go/lib/common"

type YpaReportDaily struct {
	ID        uint `gorm:"primary_key"`
	Date string `gorm:"type:date;"`
	EstimatedGrossRevenue float64 `gorm:"type:double;column:estimated_gross_revenue;"`
	RevenuePercentage1 float64 `gorm:"type:double;column:revenue_percentage_1;"`
	RevenuePercentage2 float64 `gorm:"type:double;column:revenue_percentage_2;"`
	Revenue float64 `gorm:"type:double;"`
	UpdateTime string `gorm:"type:timestamp;column:update_time;"`
	CreateTime string `gorm:"type:timestamp;column:create_time;"`
}

func (y *YpaReportDaily) Serialize() common.JSON {
	return common.JSON{
		"id":     y.ID,
		"revenue": y.Revenue,
		"revenue_percentage_1": y.RevenuePercentage1,
		"revenue_percentage_2": y.RevenuePercentage2,
		"estimated_gross_revenue": y.EstimatedGrossRevenue,
	}
}

func (y *YpaReportDaily) Read(m common.JSON) {
	y.ID = uint(m["id"].(float64))
	y.Revenue = m["revenue"].(float64)
	y.RevenuePercentage1 = m["revenue_percentage_1"].(float64)
	y.RevenuePercentage2 = m["revenue_percentage_2"].(float64)
	y.EstimatedGrossRevenue = m["estimated_gross_revenue"].(float64)
}
