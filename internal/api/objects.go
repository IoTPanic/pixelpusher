package api

// APIStatus is for the status HTTP call
type APIStatus struct {
	Status            string `json:"status"`
	DevicesRegistered int    `json:"registered"`
	DevicesOnline     int    `json:"deviceOnline"`
}
