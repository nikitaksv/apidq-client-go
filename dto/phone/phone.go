package phone

type CleanRequest struct {
	Query       string `json:"query"`
	CountryCode string `json:"countryCode"`
}

type CleanResponse struct {
	Original         string   `json:"original"`
	International    string   `json:"international"`
	National         string   `json:"national"`
	E164             string   `json:"E164"`
	RFC3966          string   `json:"RFC3966"`
	Carrier          string   `json:"carrier"`
	Country          string   `json:"country"`
	AreaCode         string   `json:"areaCode"`
	Geocoding        string   `json:"geocoding"`
	SubscriberNumber string   `json:"subscriberNumber"`
	Type             string   `json:"type"`
	Timezones        []string `json:"timezones"`
	CountryCode      int      `json:"countryCode"`
	Possible         bool     `json:"possible"`
	Valid            bool     `json:"valid"`
}
