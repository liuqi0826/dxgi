package dxgi

import (
	"testing"
)

func TestDXGIDevice_GetAdapter(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	// 设备对象通常从 D3D 创建，需要 D3D 上下文
	t.Log("Device tests require D3D device context")
}

func TestDXGIDevice_GetMaximumFrameLatency(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	// 需要实际的设备对象
	t.Log("Device tests require actual device object")
}
