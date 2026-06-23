package dxgi

import (
	"testing"
)

func TestDXGISurface_GetDesc(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	// 表面对象通常从设备创建，需要 D3D 上下文
	t.Log("Surface tests require D3D device context")
}

func TestDXGISurface_Map(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	// 需要实际的表面对象
	t.Log("Surface Map tests require actual surface object")
}
