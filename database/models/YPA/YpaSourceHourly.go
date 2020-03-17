package YPA

type YpaSourceHourly struct {
	ID        uint `gorm:"primary_key"`
	Data_hour string `gorm:"type:datetime;"`
	Product string `gorm:"type:varchar(45);"`
	Market string `gorm:"type:varchar(45);"`
	Source_tag string `gorm:"type:varchar(256);"`
	Device_type string `gorm:"type:varchar(45);"`
	Searches int `gorm:"type:int(11);"`
	Bidded_searches int `gorm:"type:int(11);"`
	Bidded_results int `gorm:"type:int(11);"`
	Bidded_clicks int `gorm:"type:int(11);"`
	Coverage float64 `gorm:"type:double;"`
	Estimated_gross_revenue float64 `gorm:"type:double;"`
	Rps float64 `gorm:"type:double;"`
	Ctr float64 `gorm:"type:double;"`
	Ppc float64 `gorm:"type:double;"`
	Tq_score float64 `gorm:"type:double;"`
}
