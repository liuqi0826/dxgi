package dxgi

import (
	"testing"
)

func TestDXGIDisplayControl_IsStereoEnabled(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	// 显示控制需要从适配器获取
	t.Log("DisplayControl tests require adapter object")
}

func TestDXGIDisplayControl_SetStereoEnabled(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	// 需要实际的显示控制对象
	t.Log("DisplayControl tests require actual display control object")
}
