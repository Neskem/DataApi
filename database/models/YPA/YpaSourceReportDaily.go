package YPA

import "DataApi.Go/lib/common"

type YpaSourceReportDaily struct {
	ID        uint `gorm:"primary_key"`
	Date string `gorm:"type:date;"`
	EstimatedGrossRevenue float64 `gorm:"type:double;column:estimated_gross_revenue;"`
	RevenuePercentage1 float64 `gorm:"type:double;column:revenue_percentage_1;"`
	RevenuePercentage2 float64 `gorm:"type:double;column:revenue_percentage_2;"`
	Revenue float64 `gorm:"type:double;"`
}

func (y *YpaSourceReportDaily) Serialize() common.JSON {
	return common.JSON{
		"id":     y.ID,
		"revenue": y.Revenue,
		"revenue_percentage_1": y.RevenuePercentage1,
		"revenue_percentage_2": y.RevenuePercentage2,
		"estimated_gross_revenue": y.EstimatedGrossRevenue,
	}
}

func (y *YpaSourceReportDaily) Read(m common.JSON) {
	y.ID = uint(m["id"].(float64))
	y.Revenue = m["revenue"].(float64)
	y.RevenuePercentage1 = m["revenue_percentage_1"].(float64)
	y.RevenuePercentage2 = m["revenue_percentage_2"].(float64)
	y.EstimatedGrossRevenue = m["estimated_gross_revenue"].(float64)
}
