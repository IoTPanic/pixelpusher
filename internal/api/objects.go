package api

import "github.com/IoTPanic/pixels-go"

// APIStatus is for the status HTTP call
type APIStatus struct {
	Status            string `json:"status"`
	DevicesRegistered int    `json:"registered"`
	DevicesOnline     int    `json:"deviceOnline"`
}

// PixelChannelUpdate is submitted to update a strip
type PixelChannelUpdate struct {
	Parent int                  `json:"parent"` // Device ID or Universe ID depending on context
	Pixels map[int]pixels.Pixel `json:"pixels"` // Key is the pixel number
	RGBW   bool                 `json:"RGBW"`
}
