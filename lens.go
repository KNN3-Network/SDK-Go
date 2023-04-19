package knn3

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ILens struct {
	ProfileId    int    `json:"profileId"`
	Creator      string `json:"creator"`
	To           string `json:"to"`
	Handle       string `json:"handle"`
	ImageURI     string `json:"imageURI"`
	FollowModule string `json:"followModule"`
	Metadata     string `json:"metadata"`
	Address      string `json:"address"`
}

type ILensList struct {
	List   []ILens `json:"list"`
	Cursor int     `json:"cursor"`
}

type IPub struct {
	Id                        int    `json:"id"`
	ProfileId                 int    `json:"profileId"`
	PubId                     int    `json:"pubId"`
	ContentURI                string `json:"contentURI"`
	ProfileIdPointed          int    `json:"profileIdPointed"`
	PubIdPointed              int    `json:"pubIdPointed"`
	ReferenceModuleData       string `json:"referenceModuleData"`
	CollectModule             string `json:"collectModule"`
	CollectModuleReturnData   string `json:"collectModuleReturnData"`
	ReferenceModule           string `json:"referenceModule"`
	ReferenceModuleReturnData string `json:"referenceModuleReturnData"`
	Type                      string `json:"type"`
	Timestamp                 int    `json:"timestamp"`
	TransactionHash           string `json:"transactionHash"`
	BlockNumber               int    `json:"blockNumber"`
}

type IPubList struct {
	List   []IPub `json:"list"`
	Cursor int    `json:"cursor"`
}

type ILensRate struct {
	ProfileId           string `json:"profileId"`
	Address             string `json:"address"`
	Influ_level         int    `json:"influ_level"`
	Influ_level_str     string `json:"influ_level_str"`
	Campaign_level      int    `json:"campaign_level"`
	Campaign_level_str  string `json:"campaign_level_str"`
	Engager_level       int    `json:"engager_level"`
	Engager_level_str   string `json:"engager_level_str"`
	Creator_level       int    `json:"creator_level"`
	Creator_level_str   string `json:"creator_level_str"`
	Collector_level     int    `json:"collector_level"`
	Collector_level_str string `json:"collector_level_str"`
	Curator_level       int    `json:"curator_level"`
	Curator_level_str   string `json:"curator_level_str"`
	Overall_level_score int    `json:"overall_level_score"`
	Overall_level_rank  int    `json:"overall_level_rank"`
	Overall_level       int    `json:"overall_level"`
	Overall_level_str   string `json:"overall_level_str"`
}

// getLensList
//
//	@param handle
//	@param profileId
//	@param limit
//	@param cursor
//	@param apiKey
//	@return *ILensList
//	@return error
func getLensList(handle string, profileId int, limit int, cursor string, apiKey string) (*ILensList, error) {
	params := make(map[string]interface{})
	if handle != "" {
		params["handle"] = handle
	}
	if profileId > 0 {
		params["profileId"] = profileId
	}
	if limit > 50 {
		params["limit"] = 50
	} else if limit > 0 {
		params["limit"] = limit
	}
	if cursor != "" {
		params["cursor"] = cursor
	}

	resp, err := doRequest(http.MethodGet, "lens", apiKey, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var lensList ILensList
	err = json.NewDecoder(resp.Body).Decode(&lensList)
	if err != nil {
		return nil, err
	}
	return &lensList, nil
}

// getLensFollowers
//
//	@param profileId
//	@param limit
//	@param cursor
//	@param apiKey
//	@return *IAddrList
//	@return error
func getLensFollowers(profileId int, limit int, cursor string, apiKey string) (*IAddrList, error) {
	params := make(map[string]interface{})
	if profileId <= 0 {
		return new(IAddrList), fmt.Errorf("missing required argument: profileId")
	}

	if limit > 50 {
		params["limit"] = 50
	} else if limit > 0 {
		params["limit"] = limit
	}
	if cursor != "" {
		params["cursor"] = cursor
	}

	resp, err := doRequest(http.MethodGet, fmt.Sprintf("lens/followers/%d", profileId), apiKey, params)
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

// getLensPublications
//
//	@param profileId
//	@param pubId
//	@param pubType
//	@param limit
//	@param cursor
//	@param apiKey
//	@return *IPubList
//	@return error
func getLensPublications(profileId int, pubId int, pubType string, limit int, cursor string, apiKey string) (*IPubList, error) {
	params := make(map[string]interface{})
	if profileId > 0 {
		params["profileId"] = profileId
	}
	if pubId > 0 {
		params["pubId"] = pubId
	}
	if pubType == "Comment" || pubType == "Post" || pubType == "Mirror" {
		params["type"] = pubType
	}
	if limit > 50 {
		params["limit"] = 50
	} else if limit > 0 {
		params["limit"] = limit
	}
	if cursor != "" {
		params["cursor"] = cursor
	}

	resp, err := doRequest(http.MethodGet, "lens/publications", apiKey, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pubList IPubList
	err = json.NewDecoder(resp.Body).Decode(&pubList)
	if err != nil {
		return nil, err
	}
	return &pubList, nil
}

// getLensRate
//
//	@param profileId
//	@param apiKey
//	@return *ILensRate
//	@return error
func getLensRate(profileId int, apiKey string) (*ILensRate, error) {
	if profileId <= 0 {
		return new(ILensRate), fmt.Errorf("missing required argument: profileId")
	}

	resp, err := doRequest(http.MethodGet, fmt.Sprintf("lens/level/%d", profileId), apiKey, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var lensRate ILensRate
	err = json.NewDecoder(resp.Body).Decode(&lensRate)
	if err != nil {
		return nil, err
	}

	return &lensRate, nil
}
