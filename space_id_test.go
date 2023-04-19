package knn3

import (
	"testing"
)

func TestGetSpaceIdList(t *testing.T) {
	knn3sdk := NewKNN3SDK("")
	// Call the GetSpaceIdList method and check the results
	spaceIdList, err := knn3sdk.GetSpaceIdList("", "", 2, "")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	t.Logf("spaceIdList: %+v", spaceIdList)
	expectedListLen := 2
	if len(spaceIdList.List) != expectedListLen {
		t.Errorf("Unexpected address list length: expected %d, got %d", expectedListLen, len(spaceIdList.List))
	}
}
