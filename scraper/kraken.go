package scraper

import (
	"fmt"
	"strings"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// APIURL is the official Kraken API Endpoint
	APIURL = "https://api.kraken.com"
	// APIVersion is the official Kraken API Version Number
	APIVersion = "0"
	// APIUserAgent identifies this library with the Kraken API
	APIUserAgent = "Kraken GO API Agent (https://github.com)"
)

var publicMethods = []string{
	"Time",
}

var privateMethods = []string{}

// a kraken api client connection
type KrakenApi struct {
	key    string
	secret string
	client *http.Client
}

func NewKraken(key, secret string) *KrakenApi {
	client := &http.Client{}

	return &KrakenApi{key, secret, client}
}

/*
	Public methods
*/

func (api *KrakenApi) Time() (*TimeResponse, error) {
	resp, err := api.queryPublic("Time", nil, &TimeResponse{})

	if err != nil {
		return nil, err
	}

	return resp.(*TimeResponse), nil
}

/*
	Query helpers
*/

func (api *KrakenApi) queryPublic(method string, values url.Values, returned interface{}) (interface{}, error) {
	url := fmt.Sprintf("%s/%s/public/%s", APIURL, APIVersion, method)

	resp, err := api.doRequest(url, values, nil, returned)

	return resp, err
}

// doRequest executes a HTTP Request to the Kraken API and returns the result
// this holds shared functionality across queryPublic and queryPrivate, hence options like `headers`
func (api *KrakenApi) doRequest(reqUrl string, values url.Values, headers map[string]string, returned interface{}) (interface{}, error) {
	req, err := http.NewRequest("POST", reqUrl, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, requestError(err.Error())
	}

	req.Header.Add("User-Agent", APIUserAgent)
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// execute the request using KrakenApi's nested http client
	resp, err := api.client.Do(req)
	if err != nil {
		return nil, requestError(err.Error())
	}
	defer resp.Body.Close()

	// read the request
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, requestError(err.Error())
	}

	// parse the request
	var jsonData KrakenResponse

	// unmarshal into given type instead of generic interface{}
	if returned != nil {
		jsonData.Result = returned
	}

	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, requestError(err.Error())
	}

	// check for a Kraken API error
	if len(jsonData.Error) > 0 {
		return nil, fmt.Errorf("Could not execute request due to Kraken error! (%s)", jsonData.Error)
	}

	return jsonData.Result, nil

}

/*
	Basic helpers
*/

func requestError(err string) error {
	return fmt.Errorf("Could not execute request! (%s)", err)
}
