package goingecko

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/JulianToledano/goingecko/events"
)

func (c *Client) Events(countryCode, eventType, page, from, to string, onlyUpcoming bool) (*events.Events, error) {
	params := url.Values{}
	params.Add("country_code", countryCode)
	params.Add("type", eventType)
	params.Add("page", page)
	params.Add("upcoming_events_only", strconv.FormatBool(onlyUpcoming))
	params.Add("from_date", from)
	params.Add("to_date", to)

	rUrl := fmt.Sprintf("%s?%s", eventsURL, params.Encode())
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}

	var data *events.Events
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) EventsCountries() (*events.Country, error) {
	rUrl := fmt.Sprintf("%s/%s", eventsURL, "countries")
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}

	var data *events.Country
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) EventsTypes() (*events.Types, error) {
	rUrl := fmt.Sprintf("%s/%s", eventsURL, "types")
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}

	var data *events.Types
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
