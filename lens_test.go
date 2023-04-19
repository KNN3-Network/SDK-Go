package knn3

import (
	"testing"
)

func TestGetLensList(t *testing.T) {
	knn3sdk := NewKNN3SDK("")
	// Call the GetLensList method and check the results
	lensList, err := knn3sdk.GetLensList("", -1, 2, "")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	t.Logf("lensList: %+v", lensList)
	expectedListLen := 2
	if len(lensList.List) != expectedListLen {
		t.Errorf("Unexpected address list length: expected %d, got %d", expectedListLen, len(lensList.List))
	}
}

func TestGetLensFollowers(t *testing.T) {
	knn3sdk := NewKNN3SDK("")
	// Call the GetLensFollowers method and check the results
	addrList, err := knn3sdk.GetLensFollowers(5, 2, "")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	t.Logf("addrList: %+v", addrList)
	expectedListLen := 2
	if len(addrList.List) != expectedListLen {
		t.Errorf("Unexpected address list length: expected %d, got %d", expectedListLen, len(addrList.List))
	}
}

func TestGetLensPublications(t *testing.T) {
	knn3sdk := NewKNN3SDK("")
	// Call the GetLensPublications method and check the results
	pubList, err := knn3sdk.GetLensPublications(-1, -1, "", 2, "")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	t.Logf("pubList: %+v", pubList)
	expectedListLen := 2
	if len(pubList.List) != expectedListLen {
		t.Errorf("Unexpected address list length: expected %d, got %d", expectedListLen, len(pubList.List))
	}
}

func TestGetLensPublicationsQueryComment(t *testing.T) {
	knn3sdk := NewKNN3SDK("")
	// Call the GetLensPublications method and check the results
	pubList, err := knn3sdk.GetLensPublications(-1, -1, "Comment", 2, "")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	t.Logf("pubList: %+v", pubList)
	expectedListLen := 2
	if len(pubList.List) != expectedListLen {
		t.Errorf("Unexpected address list length: expected %d, got %d", expectedListLen, len(pubList.List))
	}
}

func TestGetLensRate(t *testing.T) {
	knn3sdk := NewKNN3SDK("")
	// Call the GetLensRate method and check the results
	lensRate, err := knn3sdk.GetLensRate(5)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	t.Logf("lensRate: %+v", lensRate)
	if lensRate.ProfileId != "5" {
		t.Errorf("Unexpected ProfileId expected %d, got %d", 5, 5)
	}
}
