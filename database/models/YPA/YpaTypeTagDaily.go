package YPA

type YpaTypeTagDaily struct {
	ID        uint `gorm:"primary_key"`
	Date int `gorm:"type:int(11);"`
	Domain string `gorm:"type:varchar(200);"`
	Type_tag string `gorm:"type:varchar(256);"`
	Searches int `gorm:"type:int(11);"`
	Bidded_searches int `gorm:"type:int(11);"`
	Bidded_results int `gorm:"type:int(11);"`
	Bidded_clicks int `gorm:"type:int(11);"`
	Estimated_gross_revenue float64 `gorm:"type:double;"`
	Update_time string `gorm:"type:timestamp;"`
	Create_time string `gorm:"type:timestamp;"`
}
