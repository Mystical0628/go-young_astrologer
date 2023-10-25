package entity

import "go-young_astrologer/package-core/scalar"

type Apod struct {
	Id             int32
	Title          string
	Explanation    string
	Url            string
	HdUrl          string
	Date           scalar.Date
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
}
