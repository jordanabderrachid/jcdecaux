package jcdecaux

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	// ErrUnsetAPIKey comment
	ErrUnsetAPIKey = errors.New("unset API key")

	// ErrJCDecauxInternal comment
	ErrJCDecauxInternal = errors.New("internal jcdecaux error (status code 500)")

	// ErrUnsupportedStatusCode comment
	ErrUnsupportedStatusCode = errors.New("unsupported status code")

	// ErrBadRequest comment
	ErrBadRequest = errors.New("bad request")

	// ErrNotFound comment
	ErrNotFound = errors.New("not found")

	// ErrUnauthorized comment
	ErrUnauthorized = errors.New("unauthorized, verify api token")
)

// Client comment
type Client struct {
	APIKey string
}

func (c *Client) addAPIKey(u *url.URL) error {
	if c.APIKey == "" {
		return ErrUnsetAPIKey
	}

	values := u.Query()
	values.Add("apiKey", c.APIKey)
	u.RawQuery = values.Encode()

	return nil
}

// GetContracts comment
func (c *Client) GetContracts() ([]Contract, error) {
	callURL := url.URL{Scheme: "https", Host: "api.jcdecaux.com", Path: "vls/v1/contracts"}

	err := c.addAPIKey(&callURL)
	if err != nil {
		return nil, err
	}

	body, err := performRequest(callURL)
	if err != nil {
		return nil, err
	}

	contracts := make([]Contract, 100)
	err = json.Unmarshal(body, &contracts)

	return contracts, err
}

// GetStations comment
func (c *Client) GetStations() ([]Station, error) {
	callURL := url.URL{Scheme: "https", Host: "api.jcdecaux.com", Path: "vls/v1/stations"}

	err := c.addAPIKey(&callURL)
	if err != nil {
		return nil, err
	}

	body, err := performRequest(callURL)
	if err != nil {
		return nil, err
	}

	stations := make([]Station, 100)
	err = json.Unmarshal(body, &stations)

	return stations, err
}

// GetStationsByContract comment
func (c *Client) GetStationsByContract(contractName string) ([]Station, error) {
	callURL := url.URL{Scheme: "https", Host: "api.jcdecaux.com", Path: "vls/v1/stations"}

	queryValues := callURL.Query()
	queryValues.Add("contract", contractName)
	callURL.RawQuery = queryValues.Encode()

	err := c.addAPIKey(&callURL)
	if err != nil {
		return nil, err
	}

	body, err := performRequest(callURL)
	if err != nil {
		return nil, err
	}

	stations := make([]Station, 100)
	err = json.Unmarshal(body, &stations)

	return stations, err
}

// GetStation comment
func (c *Client) GetStation(stationNumber int, contractName string) (Station, error) {
	path := fmt.Sprintf("vls/v1/stations/%d", stationNumber)
	callURL := url.URL{Scheme: "https", Host: "api.jcdecaux.com", Path: path}

	err := c.addAPIKey(&callURL)
	if err != nil {
		return Station{}, err
	}

	queryValues := callURL.Query()
	queryValues.Add("contract", contractName)
	callURL.RawQuery = queryValues.Encode()

	body, err := performRequest(callURL)
	if err != nil {
		return Station{}, err
	}

	station := Station{}
	err = json.Unmarshal(body, &station)

	return station, err
}

func performRequest(u url.URL) (body []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		return ioutil.ReadAll(resp.Body)
	case 400:
		return nil, ErrBadRequest
	case 403:
		return nil, ErrUnauthorized
	case 404:
		return nil, ErrNotFound
	case 500:
		return nil, ErrJCDecauxInternal
	default:
		return nil, ErrUnsupportedStatusCode
	}
}
