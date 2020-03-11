package models

import (
	"DataApi.Go/lib/common"
)
type SumPV struct {
	Total int
}

type StatPagePV struct {
	ID        uint `gorm:"primary_key"`
	Datetime_intid int
	Page_id string `gorm:"type:varchar(64);"`
	Page_title string
	Page_author string
	Page_url string
	Page_hostname string
	Pv_algonum string
	Pv int
	Pv_valid int
	Pv_invalid int
	Ypa_age int
	Ypa_gender int
	Fb_id int
	Line_id int
	Highlightedtext int
	Openlink int
	Pv_count int
	Stay_0_count int
	Stay_1_count int
	SumPV int
}

// Serialize serializes user data
func (u *StatPagePV) Serialize() common.JSON {
	return common.JSON{
		"id":     u.ID,
		"pv":     u.Pv,
		"page_id":		u.Page_id,
	}
}

func (u *StatPagePV) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	u.Pv = m["pv"].(int)
	u.Page_id = m["page_id"].(string)
}

func (u *SumPV) Serialize() common.JSON {
	return common.JSON{
		"total":     u.Total,
	}
}

func (u *SumPV) Read(m common.JSON) {
	u.Total = m["total"].(int)
}
