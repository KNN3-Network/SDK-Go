package knn3

import (
	"encoding/json"
	"net/http"
)

type ISpaceId struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type ISpaceIdList struct {
	List   []ISpaceId `json:"list"`
	Cursor string     `json:"cursor"`
}

// getSpaceIdList
//
//	@param spaceId
//	@param addr
//	@param limit
//	@param cursor
//	@param apiKey
//	@return *ISpaceIdList
//	@return error
func getSpaceIdList(spaceId string, addr string, limit int, cursor string, apiKey string) (*ISpaceIdList, error) {
	params := make(map[string]interface{})
	if spaceId != "" {
		params["spaceId"] = spaceId
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

	resp, err := doRequest(http.MethodGet, "spaceId", apiKey, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var spaceIdList ISpaceIdList
	err = json.NewDecoder(resp.Body).Decode(&spaceIdList)
	if err != nil {
		return nil, err
	}
	return &spaceIdList, nil
}
