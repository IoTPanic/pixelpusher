package api

type APIStatus struct {
	Status            string `json:"status"`
	DevicesRegistered int    `json:"registered"`
	DevicesOnline     int    `json:"deviceOnline"`
}
