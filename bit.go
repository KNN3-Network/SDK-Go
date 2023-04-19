package knn3

import (
	"encoding/json"
	"net/http"
)

type IBit struct {
	Address     string `json:"addr"`
	Account     string `json:"account"`
	AlgorithmId int    `json:"algorithmId"`
	Chain       string `json:"chain"`
	Outpoint    string `json:"outpoint"`
}

type IBitList struct {
	List   []IBit `json:"list"`
	Cursor string `json:"cursor"`
}

// getSpaceIdList
//
//	@param account
//	@param addr
//	@param limit
//	@param cursor
//	@param apiKey
//	@return *ISpaceIdList
//	@return error
func getBitList(account string, addr string, limit int, cursor string, apiKey string) (*IBitList, error) {
	params := make(map[string]interface{})
	if account != "" {
		params["account"] = account
	}
	if addr != "" {
		params["addr"] = addr
	}
	if limit > 50 {
		params["limit"] = 50
	} else if limit > 0 {
		params["limit"] = limit
	}
	if cursor != "" {
		params["cursor"] = cursor
	}

	resp, err := doRequest(http.MethodGet, "bits", apiKey, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var bitList IBitList
	err = json.NewDecoder(resp.Body).Decode(&bitList)
	if err != nil {
		return nil, err
	}
	return &bitList, nil
}
