package address

import (
	"fmt"
)

type Codes struct {
	Fias string `json:"fias"`
	Ga   string `json:"ga"`
	Osm  string `json:"osm"`
}

type Part struct {
	Codes    *Codes `json:"codes"`
	FullName string `json:"fullName"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

type HouseDetails struct {
	FullName string `json:"fullName"`
	House    string `json:"house"`
	Case     string `json:"case"`
	Build    string `json:"build"`
	Liter    string `json:"liter"`
	Lend     string `json:"lend"`
	Block    string `json:"block"`
	Pav      string `json:"pav"`
	Floor    string `json:"floor"`
	Flat     string `json:"flat"`
	Office   string `json:"office"`
	Kab      string `json:"kab"`
	Abon     string `json:"abon"`
	Plot     string `json:"plot"`
	Sek      string `json:"sek"`
	Entr     string `json:"entr"`
	Room     string `json:"room"`
	Hostel   string `json:"hostel"`
	Munit    string `json:"munit"`
}

type Coordinates struct {
	// Широта в градусах. Значение должно быть в диапазоне [-90.0, +90.0].
	Latitude float64 `json:"latitude"`
	// Долгота в градусах. Значение должно быть в диапазоне[-180.0, +180.0].
	Longitude float64 `json:"longitude"`
}

func (c Coordinates) String() string {
	return fmt.Sprintf("%f,%f", c.Latitude, c.Longitude)
}

type Country struct {
	Name    string `json:"name"`
	Alpha2  string `json:"alpha2"`
	Alpha3  string `json:"alpha3"`
	Numeric int    `json:"numeric"`
}

type CleanRequest struct {
	Query       string `json:"query"`
	CountryCode string `json:"countryCode"`
}

type CleanResponse struct {
	Address
	Quality    *Quality `json:"quality"`
	Original   string   `json:"original"`
	PostcodeIn string   `json:"postcodeIn"`
	Valid      bool     `json:"valid"`
}

type SuggestRequest struct {
	Query       string `json:"query"`
	CountryCode string `json:"countryCode"`
	Count       int    `json:"count"`
}

type SuggestResponse struct {
	Suggestions []*Address `json:"suggestions"`
}

type Address struct {
	Region        *Part         `json:"region"`
	Area          *Part         `json:"area"`
	City          *Part         `json:"city"`
	CityArea      *Part         `json:"cityArea"`
	Settlement    *Part         `json:"settlement"`
	PlanStructure *Part         `json:"planStructure"`
	Street        *Part         `json:"street"`
	HouseDetails  *HouseDetails `json:"houseDetails"`
	Coordinates   *Coordinates  `json:"coordinates"`
	Country       *Country      `json:"country"`
	Address       string        `json:"address"`
	Postcode      string        `json:"postcode"`
}

type CleanIqdqResponse struct {
	CJsonKvant  *HouseDetails `json:"c_json_kvant"`
	CCoordinate *struct {
		CLon   float64 `json:"c_lon"`
		CLat   float64 `json:"c_lat"`
		CLevel int     `json:"c_level"`
	} `json:"c_coordinate"`
	CIscheck         string `json:"c_ischeck"`
	CIndexIn         string `json:"c_index_in"`
	CZipcode         string `json:"c_zipcode"`
	CAddressOriginal string `json:"c_address_original"`
	CAddressFull     string `json:"c_address_full"`
	CKladr           string `json:"c_kladr"`
	CGaCode          string `json:"c_gaCode"`
	CCountry         string `json:"c_country"`
	CCountryIsoCode  string `json:"c_country_iso_code"`
	CRegionName      string `json:"c_region_name"`
	CRegionAbbr      string `json:"c_region_abbr"`
	CRegionFias      string `json:"c_region_fias"`
	CDistrictName    string `json:"c_district_name"`
	CDistrictAbbr    string `json:"c_district_abbr"`
	CDistrictFias    string `json:"c_district_fias"`
	CCityName        string `json:"c_city_name"`
	CCityAbbr        string `json:"c_city_abbr"`
	CCityFias        string `json:"c_city_fias"`
	CCommunityName   string `json:"c_community_name"`
	CCommunityAbbr   string `json:"c_community_abbr"`
	CCommunityFias   string `json:"c_community_fias"`
	CStreetName      string `json:"c_street_name"`
	CStreetAbbr      string `json:"c_street_abbr"`
	CStreetFias      string `json:"c_street_fias"`
	CHouseStr        string `json:"c_house_str"`
	CAddrLost        string `json:"c_addr_lost"`
	CStatusError     string `json:"c_status_error"`
	CHouseError      string `json:"c_house_error"`
	CHouseErrorDesc  string `json:"c_house_error_desc"`
	CKladr19         string `json:"c_kladr19"`
	CGninmb          string `json:"c_gninmb"`
	COkato           string `json:"c_okato"`
	COktmo           string `json:"c_oktmo"`
	CAoguid          string `json:"c_aoguid"`
	CAolevel         string `json:"c_aolevel"`
	CHouseguid       string `json:"c_houseguid"`
	CTimezone        string `json:"c_timezone"`
}
