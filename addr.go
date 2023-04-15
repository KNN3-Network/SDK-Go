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

func getAddr(address string, apiKey string) (IAddress, error) {
	resp, err := doRequest(http.MethodGet, fmt.Sprintf("addresses/%s", address), apiKey)
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
