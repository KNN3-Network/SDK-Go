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
