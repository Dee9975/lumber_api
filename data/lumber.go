package data

type Lumber struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Length int `json:"len"`
	Amount int `json:"amount"`
	TeamID int `json:"team_id"`
}

type TeamLumberResponse struct {
	ID     int `json:"id"`
	Width  int `json:"width"`
	Height int `json:"height"`
	Length int `json:"len"`
	Amount int `json:"amount"`
}

type LumberResponse struct {
	Lumber Lumber `json:"lumber"`
	Volume int    `json:"volume"`
}
