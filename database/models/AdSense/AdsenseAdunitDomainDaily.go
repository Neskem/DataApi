package AdSense


type AdsenseAdunitDomainDaily struct {
	ID        uint `gorm:"primary_key"`
	Date string `gorm:"type:date;"`
	Ad_client_id string `gorm:"type:varchar(45);"`
	Ad_unit_code uint64 `gorm:"type:bigint(20);"`

}