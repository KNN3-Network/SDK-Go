package knn3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IAddress struct {
	Address string   `json:"addr"`
	Ens     []string `json:"ens"`
	Name    string   `json:"name"`
}

type IEvent struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"imageUrl"`
}

type IAddrList struct {
	List   []IAddress `json:"list"`
	Cursor string     `json:"cursor"`
}

type IAddrAttendEventsList struct {
	List   []IEvent `json:"list"`
	Cursor *string  `json:"cursor,omitempty"`
}

func getAddr(address string, apiKey string) (IAddress, error) {
	if address == "" {
		return IAddress{}, fmt.Errorf("missing required argument: address")
	}
	resp, err := doRequest(http.MethodGet, fmt.Sprintf("addresses/%s", address), apiKey, nil)
	if err != nil {
		return IAddress{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return IAddress{}, err
	}
	var addr IAddress
	err = json.Unmarshal(body, &addr)
	if err != nil {
		return IAddress{}, err
	}
	return addr, nil
}

func getAddrList(addressIn []string, limit int, cursor string, apiKey string) (*IAddrList, error) {
	params := make(map[string]interface{})
	if len(addressIn) > 0 {
		params["addressIn"] = addressIn
	}
	if limit > 50 {
		params["limit"] = 50
	} else if limit > 0 {
		params["limit"] = limit
	}
	if cursor != "" {
		params["cursor"] = cursor
	}
	resp, err := doRequest(http.MethodGet, "addresses", apiKey, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var addrList IAddrList
	err = json.NewDecoder(resp.Body).Decode(&addrList)
	if err != nil {
		return nil, err
	}
	return &addrList, nil
}

func addEvents(address string, eventName string, limit int, cursor string, apiKey string) (*IAddrAttendEventsList, error) {
	if address == "" {
		return new(IAddrAttendEventsList), fmt.Errorf("missing required argument: address")
	}
	params := make(map[string]interface{})
	params["address"] = address
	if limit > 50 {
		params["limit"] = 50
	} else if limit > 0 {
		params["limit"] = limit
	}
	if eventName != "" {
		params["eventName"] = eventName
	}
	if cursor != "" {
		params["cursor"] = cursor
	}
	resp, err := doRequest(http.MethodGet, "addresses/attendEvents", apiKey, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var list IAddrAttendEventsList
	err = json.NewDecoder(resp.Body).Decode(&list)
	if err != nil {
		return nil, err
	}
	return &list, nil
}
