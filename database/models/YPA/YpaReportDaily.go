package YPA

import "DataApi.Go/lib/common"

type YpaReportDaily struct {
	ID        uint `gorm:"primary_key"`
	Date string `gorm:"type:date;"`
	Estimated_gross_revenue float64 `gorm:"type:double;"`
	Revenue_percentage_1 float64 `gorm:"type:double;"`
	Revenue_percentage_2 float64 `gorm:"type:double;"`
	Revenue float64 `gorm:"type:double;"`
	Update_time string `gorm:"type:timestamp;"`
	Create_time string `gorm:"type:timestamp;"`
}

func (y *YpaReportDaily) Serialize() common.JSON {
	return common.JSON{
		"id":     y.ID,
		"revenue": y.Revenue,
		"revenue_percentage_1": y.Revenue_percentage_1,
		"revenue_percentage_2": y.Revenue_percentage_2,
		"estimated_gross_revenue": y.Estimated_gross_revenue,
	}
}

func (y *YpaReportDaily) Read(m common.JSON) {
	y.ID = uint(m["id"].(float64))
	y.Revenue = m["revenue"].(float64)
	y.Revenue_percentage_1 = m["revenue_percentage_1"].(float64)
	y.Revenue_percentage_2 = m["revenue_percentage_2"].(float64)
	y.Estimated_gross_revenue = m["estimated_gross_revenue"].(float64)
}
