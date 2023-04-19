package knn3

import (
	"testing"
)

func TestGetBitList(t *testing.T) {
	knn3sdk := NewKNN3SDK("")
	// Call the GetBitList method and check the results
	bitList, err := knn3sdk.GetBitList("", "", 2, "")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	t.Logf("bitList: %+v", bitList)
	expectedListLen := 2
	if len(bitList.List) != expectedListLen {
		t.Errorf("Unexpected address list length: expected %d, got %d", expectedListLen, len(bitList.List))
	}
}
