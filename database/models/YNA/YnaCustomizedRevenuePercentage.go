package YNA

type YnaCustomizedRevenuePercentage struct {
	ID        uint `gorm:"primary_key"`
	AdUnit_id int `gorm:"type:int(11);"`
	CustomerRevenuePercentage float64 `gorm:"type:double;"`
	Update_time string `gorm:"type:timestamp;"`
	Create_time string `gorm:"type:timestamp;"`
}
