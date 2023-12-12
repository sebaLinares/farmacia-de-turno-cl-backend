package models

type Pharmacy struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Phone      string  `json:"phone"`
	Address    string  `json:"address"`
	RegionID   int     `json:"regionId"`
	RegionName string  `json:"regionName"`
	CommuneID  int     `json:"communeId"`
	Commune    string  `json:"communeName"`
	OpenTime   string  `json:"openTime"`
	CloseTime  string  `json:"closeTime"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}
