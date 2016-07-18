package jcdecaux

import "fmt"

// Station comment
type Station struct {
	Number       int    `json:"number"`
	ContractName string `json:"contract_name"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Position     struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"position"`
	Banking             bool   `json:"banking"`
	Bonus               bool   `json:"bonus"`
	Status              string `json:"status"`
	BikeStands          int    `json:"bike_stands"`
	AvailableBikeStands int    `json:"available_bike_stands"`
	AvailableBikes      int    `json:"available_bikes"`
	LastUpdate          int    `json:"last_update"`
}

func (s Station) String() string {
	return fmt.Sprintf("%d - %s (%s)\n bike stands: %d, available bike stands: %d, availaible bikes: %d\n", s.Number, s.Name, s.ContractName, s.BikeStands, s.AvailableBikeStands, s.AvailableBikes)
}

// Contract comment
type Contract struct {
	Name           string `json:"name"`
	CommercialName string `json:"commercial_name"`
	CountryCode    string `json:"country_code"`
	Cities         []string
}

func (c Contract) String() string {
	return fmt.Sprintf("%s (%s)\n", c.Name, c.CountryCode)
}
