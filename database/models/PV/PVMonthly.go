package PV

import "DataApi.Go/lib/common"

type PVMonthly struct {
	ID        uint `gorm:"primary_key"`
	Page_id string `gorm:"type:varchar(64);"`
	Pv int
	Pv_valid int
	Pv_invalid int
}

func (u *PVMonthly) Serialize() common.JSON {
	return common.JSON{
		"id":     u.ID,
		"pv":     u.Pv,
		"page_id":		u.Page_id,
	}
}

func (u *PVMonthly) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	u.Pv = m["pv"].(int)
	u.Page_id = m["page_id"].(string)
}
