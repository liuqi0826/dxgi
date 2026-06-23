package dxgi

import (
	"testing"
)

func TestDXGIOutputDuplication_GetDesc(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	// 输出复制需要特定的权限和设备
	t.Log("OutputDuplication tests require specific permissions and device")
}

func TestDXGIOutputDuplication_AcquireNextFrame(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	// 需要实际的输出复制对象
	t.Log("OutputDuplication tests require actual duplication object")
}
