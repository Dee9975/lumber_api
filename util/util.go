package util

import "lumber/data"

func CreateLumberResponse(l data.Lumber) *data.LumberResponse {
	volume := l.Width * l.Length * l.Height * l.Amount

	return &data.LumberResponse{
		Lumber: l,
		Volume: volume,
	}
}
