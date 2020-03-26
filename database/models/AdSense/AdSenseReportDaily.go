package AdSense


type AdSenseReportDaily struct {
	ID        uint `gorm:"primary_key"`
	Date string `gorm:"type:date;"`
	AdClientId string `gorm:"type:varchar(45);column:ad_client_id;"`
	DomainName string `gorm:"type:varchar(256);column:domain_name;"`
	PageViews int `gorm:"type:int(11);column:page_views;"`
	MatchedAdRequests int `gorm:"type:int(11);column:matched_ad_requests;"`
	Clicks int `gorm:"type:int(11);column:clicks;"`
	Earnings float64 `gorm:"type:double;column:earnings;"`
	PageRpm float64 `gorm:"type:double;column:page_rpm;"`
	ImpressionRpm float64 `gorm:"type:double;column:impression_rpm;"`
	AccountId string `gorm:"type:varchar(45);column:account_id;"`
	NetworkCode uint64 `gorm:"type:bigint(20);column:network_code;"`
	AdExchangeImpressions int `gorm:"type:int(11);column:ad_exchange_impressions;"`
	AdExchangeClicks int `gorm:"type:int(11);column:ad_exchange_clicks;"`
	AdExchangeEstimatedRevenue float64 `gorm:"type:double;column:ad_exchange_estimated_revenue;"`
	AdExchangeImpressionRpm float64 `gorm:"type:double;column:ad_exchange_impression_rpm;"`
	CustomerAdExchangeEstimatedRevenuePercentage float64 `gorm:"type:double;column:customer_ad_exchange_estimated_revenue_percentage;"`
	CustomerAdExchangeEstimatedRevenue float64 `gorm:"type:double;column:customer_ad_exchange_estimated_revenue;"`
}

type AdSenseRevenue struct {
	AccountId string `gorm:"type:varchar(45);column:account_id;"`
	CustomerAdExchangeEstimatedRevenue float64 `gorm:"type:double;column:customer_ad_exchange_estimated_revenue;"`
}

type AdSenseDomain struct {
	AccountId string `gorm:"type:varchar(45);column:account_id;"`
	DomainName string `gorm:"type:varchar(256);column:domain_name;"`
}