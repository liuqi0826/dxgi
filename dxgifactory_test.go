package dxgi

import (
	"testing"
)

func TestDXGIFactory_EnumAdapters(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	factory, err := CreateDXGIFactory1(IID_IDXGIFactory1)
	if err != nil {
		t.Fatalf("CreateDXGIFactory1 failed: %v", err)
	}
	defer factory.Release()

	// 测试 EnumAdapters (旧版本)
	adapter, err := factory.EnumAdapters(0)
	if err != nil {
		t.Logf("EnumAdapters(0) failed (expected if no adapter): %v", err)
	} else {
		defer adapter.Release()
		desc, err := adapter.GetDesc()
		if err == nil {
			t.Logf("Adapter 0 description retrieved successfully")
			_ = desc
		}
	}
}

func TestDXGIFactory_EnumAdapters1(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	factory, err := CreateDXGIFactory1(IID_IDXGIFactory1)
	if err != nil {
		t.Fatalf("CreateDXGIFactory1 failed: %v", err)
	}
	defer factory.Release()

	// 枚举所有适配器
	count := 0
	for i := uint32(0); i < 10; i++ {
		adapter, err := factory.EnumAdapters1(i)
		if err != nil {
			break
		}
		defer adapter.Release()
		count++
	}

	if count == 0 {
		t.Fatal("No adapters found")
	}
	t.Logf("Enumerated %d adapters", count)
}

func TestDXGIFactory_IsCurrent(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	factory, err := CreateDXGIFactory1(IID_IDXGIFactory1)
	if err != nil {
		t.Fatalf("CreateDXGIFactory1 failed: %v", err)
	}
	defer factory.Release()

	isCurrent := factory.IsCurrent()
	t.Logf("Factory IsCurrent: %v", isCurrent)
}

func TestDXGIFactory_IsWindowedStereoEnabled(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	factory, err := CreateDXGIFactory1(IID_IDXGIFactory1)
	if err != nil {
		t.Fatalf("CreateDXGIFactory1 failed: %v", err)
	}
	defer factory.Release()

	// 需要 QueryInterface 到 IDXGIFactory2
	enabled := factory.IsWindowedStereoEnabled()
	t.Logf("Windowed stereo enabled: %v", enabled)
}

func TestDXGIFactory_GetWindowAssociation(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	factory, err := CreateDXGIFactory1(IID_IDXGIFactory1)
	if err != nil {
		t.Fatalf("CreateDXGIFactory1 failed: %v", err)
	}
	defer factory.Release()

	hwnd, err := factory.GetWindowAssociation()
	if err != nil {
		t.Logf("GetWindowAssociation failed: %v", err)
	} else {
		t.Logf("Window association handle: %v", hwnd)
	}
}

func TestDXGIFactory_EnumWarpAdapter(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	factory, err := CreateDXGIFactory1(IID_IDXGIFactory1)
	if err != nil {
		t.Fatalf("CreateDXGIFactory1 failed: %v", err)
	}
	defer factory.Release()

	// 需要 QueryInterface 到 IDXGIFactory4
	warpAdapter, err := factory.EnumWarpAdapter(IID_IDXGIAdapter1)
	if err != nil {
		t.Logf("EnumWarpAdapter failed (may not be available): %v", err)
	} else {
		defer warpAdapter.Release()
		t.Log("WARP adapter enumerated successfully")
	}
}
