package YPA

type YpaTypeDaily struct {
	ID        uint `gorm:"primary_key"`
	Data_date string `gorm:"type:date;"`
	Product string `gorm:"type:varchar(45);"`
	Market string `gorm:"type:varchar(45);"`
	Source_tag string `gorm:"type:varchar(265);"`
	Device_type string `gorm:"type:varchar(45);"`
	Type_tag string `gorm:"type:varchar(256);"`
	Searches int `gorm:"type:int(11);"`
	Bidded_searches int `gorm:"type:int(11);"`
	Bidded_results int `gorm:"type:int(11);"`
	Bidded_clicks int `gorm:"type:int(11);"`
	Estimated_gross_revenue float64 `gorm:"type:double;"`
	Coverage float64 `gorm:"type:double;"`
	Ctr float64 `gorm:"type:double;"`
	Ppc float64 `gorm:"type:double;"`
	Tq_score float64 `gorm:"type:double;"`
	Rn float64 `gorm:"type:double;"`
}
