package models

type ForecastWeek struct {
	Week    int     `json:"week"`
	Opening float64 `json:"opening"`
	Inflow  float64 `json:"inflow"`
	Outflow float64 `json:"outflow"`
	Closing float64 `json:"closing"`
	Warning bool    `json:"warning"`
}

type Forecast struct {
	StartingCash float64        `json:"starting_cash"`
	Weeks        []ForecastWeek `json:"weeks"`
}
