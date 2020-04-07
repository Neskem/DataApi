package PV

import "DataApi.Go/lib/common"

type PVMonthly struct {
	ID        uint `gorm:"primary_key"`
	PageId string `gorm:"type:varchar(128);column:page_id;"`
	Pv int `gorm:"type:int(11);column:pv;"`
	PvValid int `gorm:"type:int(11);column:pv_valid;"`
	PvInvalid int `gorm:"type:int(11);column:pv_invalid;"`
}

func (u *PVMonthly) Serialize() common.JSON {
	return common.JSON{
		"id":     u.ID,
		"pv":     u.Pv,
		"page_id":		u.PageId,
	}
}

func (u *PVMonthly) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	u.Pv = m["pv"].(int)
	u.PageId = m["page_id"].(string)
}
