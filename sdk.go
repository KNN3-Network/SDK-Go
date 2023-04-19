package knn3

var baseURL string = "https://knn3-gateway.knn3.xyz/data-api"

type KNN3SDK struct {
	apiKey string
}

func NewKNN3SDK(apiKey string) *KNN3SDK {
	if apiKey == "" {
		apiKey = "knn3-common-AswT-mcYf"
	}

	return &KNN3SDK{apiKey: apiKey}
}

func (sdk *KNN3SDK) GetAddr(address string) (IAddress, error) {
	return getAddr(address, sdk.apiKey)
}

func (sdk *KNN3SDK) GetAddrList(addressIn []string, limit int, cursor string) (*IAddrList, error) {
	return getAddrList(addressIn, limit, cursor, sdk.apiKey)
}

func (sdk *KNN3SDK) AddEvents(address string, eventName string, limit int, cursor string) (*IAddrAttendEventsList, error) {
	return addEvents(address, eventName, limit, cursor, sdk.apiKey)
}

func (sdk *KNN3SDK) GetSpaceIdList(spaceId string, addr string, limit int, cursor string) (*ISpaceIdList, error) {
	return getSpaceIdList(spaceId, addr, limit, cursor, sdk.apiKey)
}

func (sdk *KNN3SDK) GetBitList(account string, addr string, limit int, cursor string) (*IBitList, error) {
	return getBitList(account, addr, limit, cursor, sdk.apiKey)
}

func (sdk *KNN3SDK) GetLensList(handle string, profileId int, limit int, cursor string) (*ILensList, error) {
	return getLensList(handle, profileId, limit, cursor, sdk.apiKey)
}

func (sdk *KNN3SDK) GetLensFollowers(profileId int, limit int, cursor string) (*IAddrList, error) {
	return getLensFollowers(profileId, limit, cursor, sdk.apiKey)
}

func (sdk *KNN3SDK) GetLensPublications(profileId int, pubId int, pubType string, limit int, cursor string) (*IPubList, error) {
	return getLensPublications(profileId, pubId, pubType, limit, cursor, sdk.apiKey)
}

func (sdk *KNN3SDK) GetLensRate(profileId int) (*ILensRate, error) {
	return getLensRate(profileId, sdk.apiKey)
}
