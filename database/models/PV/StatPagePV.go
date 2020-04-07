package PV

import (
	"DataApi.Go/lib/common"
)
type SumPV struct {
	Total int
}

type StatPagePV struct {
	ID        uint `gorm:"primary_key"`
	DatetimeIntid int `gorm:"type:int(11);column:datetime_intid;"`
	PageId string `gorm:"type:varchar(64);column:page_id;"`
	PageTitle string `gorm:"type:varchar(2048);column:page_title;"`
	PageAuthor string `gorm:"type:varchar(45);column:page_author;"`
	PageUrl string `gorm:"type:varchar(8192);column:page_url;"`
	PageHostname string `gorm:"type:varchar(512);column:page_hostname;"`
	PvAlgonum string `gorm:"type:varchar(45);column:pv_algonum;"`
	Pv int `gorm:"type:int(11);column:pv;"`
	PvValid int `gorm:"type:int(11);column:pv_valid;"`
	PvInvalid int `gorm:"type:int(11);column:pv_invalid;"`
	YpaAge int `gorm:"type:int(11);column:ypa_age;"`
	YpaGender int `gorm:"type:int(11);column:ypa_gender;"`
	FbId int `gorm:"type:int(11);column:fb_id;"`
	LineId int `gorm:"type:int(11);column:line_id;"`
	Highlightedtext int `gorm:"type:int(11);column:highlightedtext;"`
	Openlink int `gorm:"type:int(11);column:openlink;"`
	PvCount int `gorm:"type:int(11);column:pv_count;"`
	Stay0Count int `gorm:"type:int(11);column:stay_0_count;"`
	Stay1Count int `gorm:"type:int(11);column:stay_1_count;"`
}

// Serialize serializes user data
func (u *StatPagePV) Serialize() common.JSON {
	return common.JSON{
		"id":     u.ID,
		"pv":     u.Pv,
		"page_id":		u.PageId,
	}
}

func (u *StatPagePV) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	u.Pv = m["pv"].(int)
	u.PageId = m["page_id"].(string)
}

func (u *SumPV) Serialize() common.JSON {
	return common.JSON{
		"total":     u.Total,
	}
}

func (u *SumPV) Read(m common.JSON) {
	u.Total = m["total"].(int)
}
