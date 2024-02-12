package data

type RateConfig struct {
	Name         string `json:"name"`
	HeatingRate  int    `json:"heating_rate"`
	ForkliftRate int    `json:"forklift_rate"`
	DefaultRate  int    `json:"default_rate"`
	HolidayRate  int    `json:"holiday_rate"`
}
