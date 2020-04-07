package AdSense

type AdmanagerAdexchangeAdunitReportDaily struct {
	ID        uint `gorm:"primary_key"`
	NetworkCode string `gorm:"type:bigint(20);"`
	Date string `gorm:"type:date;"`
	AdExchangeSiteName string `gorm:"type:varchar(256);column:ad_exchange_site_name;"`
	DeviceCategoryId string `gorm:"type:varchar(45);column:device_category_id;"`
	DeviceCategoryName string `gorm:"type:varchar(45);column:device_category_id;"`
	AdExchangeDfpAdUnit string `gorm:"type:varchar(256);column:ad_exchange_dfp_ad_unit;"`
	AdExchangeAdRequests int `gorm:"type:int(11);column:ad_exchange_ad_requests;"`
	AdExchangeMatchedRequests int `gorm:"type:int(11);column:ad_exchange_matched_requests;"`
	AdExchangeCoverage float64 `gorm:"type:double;column:ad_exchange_coverage;"`
	AdExchangeClicks int `gorm:"type:int(11);column:ad_exchange_clicks;"`
	AdExchangeAdRequestCtr float64 `gorm:"type:double;column:ad_exchange_ad_request_ctr;"`
	AdExchangeCtr float64 `gorm:"type:double;column:ad_exchange_ctr;"`
	AdExchangeAdCtr float64 `gorm:"type:double;column:ad_exchange_ad_ctr;"`
	AdExchangeCpc float64 `gorm:"type:double;column:ad_exchange_cpc;"`
	AdExchangeAdRequestEcpm float64 `gorm:"type:double;column:ad_exchange_ad_request_ecpm;"`
	AdExchangeMatchedEcpm float64 `gorm:"type:double;column:ad_exchange_matched_ecpm;"`
	AdExchangeLift float64 `gorm:"type:double;column:ad_exchange_lift;"`
	AdExchangeEstimatedRevenue float64 `gorm:"type:double;column:ad_exchange_estimated_revenue;"`
	AdExchangeImpressions int `gorm:"type:int(11);column:ad_exchange_impressions;"`
	AdExchangeAdEcpm float64 `gorm:"type:double;column:ad_exchange_ad_ecpm;"`
}