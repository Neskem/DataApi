package YPA

import "DataApi.Go/lib/common"

type YpaSourceTagIgnore struct {
	ID        uint `gorm:"primary_key"`
	Source_tag string `gorm:"type:varchar(64);"`
}

func (y *YpaSourceTagIgnore) Serialize() common.JSON {
	return common.JSON{
		"id":     y.ID,
		"source_tag": y.Source_tag,
	}
}

func (y *YpaSourceTagIgnore) Read(m common.JSON) {
	y.ID = uint(m["id"].(float64))
	y.Source_tag = m["source_tag"].(string)
}
