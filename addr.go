package knn3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type IAddress struct {
	Address string   `json:"addr"`
	Ens     []string `json:"ens"`
	Name    string   `json:"name"`
}

type IAddrList struct {
	List   []IAddress `json:"list"`
	Cursor string     `json:"cursor"`
}

func getAddr(address string, apiKey string) (IAddress, error) {
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
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}
	if cursor != "" {
		params["cursor"] = cursor
	}
	resp, err := doRequest(http.MethodGet, "addresses", apiKey, params)
	if err != nil {
		return new(IAddrList), err
	}
	defer resp.Body.Close()
	var addrList IAddrList
	err = json.NewDecoder(resp.Body).Decode(&addrList)
	if err != nil {
		return new(IAddrList), err
	}
	return &addrList, nil
}
