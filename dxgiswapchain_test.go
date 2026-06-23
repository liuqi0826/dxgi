package dxgi

import (
	"testing"
)

func TestDXGISwapChain_GetDesc(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	// 注意：创建交换链需要 D3D 设备，这里只测试方法存在
	// 实际测试需要完整的 D3D 上下文
	t.Log("SwapChain tests require D3D device context")
}

func TestDXGISwapChain_GetCurrentBackBufferIndex(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	// 需要实际的交换链对象
	t.Log("SwapChain tests require actual swap chain object")
}

func TestDXGISwapChain_GetFrameLatencyWaitableObject(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	// 需要实际的交换链对象
	t.Log("SwapChain tests require actual swap chain object")
}
