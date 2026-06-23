package dxgi

import (
	"testing"
)

func TestDXGIResource_GetUsage(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	// 资源对象通常从设备创建
	t.Log("Resource tests require D3D device context")
}

func TestDXGIResource_GetEvictionPriority(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	// 需要实际的资源对象
	t.Log("Resource tests require actual resource object")
}
