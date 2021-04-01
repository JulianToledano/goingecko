package events

type Country struct {
	Data []CountryData `json:"data"`
}

type CountryData struct {
	Country string `json:"country"`
	Code    string `json:"code"`
}
