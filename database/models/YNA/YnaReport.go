package YNA

import (
	"DataApi.Go/lib/common"
	"time"
)
type YnaReprot struct {
	ID        uint `gorm:"primary_key"`
	Date int `gorm:"type:int(11);column:date;"`
	AdUnitId int `gorm:"type:int(11);column:adunit_id;"`
	Impressions int `gorm:"type:int(11);"`
	Clicks int `gorm:"type:int(11);"`
	Revenuennusd float64 `gorm:"type:double;column:revenueInUSD;"`
	Revenueintwd float64 `gorm:"type:double;column:revenueInTWD;"`
	Customerrevenuepercentage float64 `gorm:"type:double;column:customerRevenuePercentage;"`
	Customerrevenueintwd float64 `gorm:"type:double;column:customerRevenueInTWD;"`
	UpdateTime time.Time `gorm:"type:timestamp;"`
	CreateTime time.Time `gorm:"type:timestamp;"`
}

func (y *YnaReprot) Serialize() common.JSON {
	return common.JSON{
		"id":     y.ID,
		"adunit_id": y.AdUnitId,
		"impressions": y.Impressions,
		"clicks": y.Clicks,
		"revenueintwd": y.Revenueintwd,
		"customerrevenueintwd": y.Customerrevenueintwd,
	}
}

func (y *YnaReprot) Read(m common.JSON) {
	y.ID = uint(m["id"].(float64))
	y.AdUnitId = m["adunit_id"].(int)
	y.Impressions = m["impressions"].(int)
	y.Clicks = m["clicks"].(int)
	y.Revenueintwd = m["revenueintwd"].(float64)
	y.Customerrevenueintwd = m["customerrevenueintwd"].(float64)
}
