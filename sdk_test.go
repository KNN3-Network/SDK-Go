package knn3

import (
	"strings"
	"testing"
)

func TestGetAddress(t *testing.T) {
	// 创建 KNN3SDK 实例并设置 API 密钥
	knn3sdk := NewKNN3SDK("")

	// 调用 GetAddr 函数获取地址信息
	address := "0x88520C10ad3d35aD2D3220CdE446CcB33f09331B" // 请替换为您要查询的地址
	addr, err := knn3sdk.GetAddr(address)
	if err != nil {
		t.Fatal(err)
	}

	// 检查返回的地址是否与期望相符
	expectedAddress := strings.ToLower(address)
	if addr.Address != expectedAddress {
		t.Errorf("Expected address '%s', but got '%s'", expectedAddress, addr.Address)
	}
}
