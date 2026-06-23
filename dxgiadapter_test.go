package dxgi

import (
	"testing"
)


func TestDXGIAdapter_GetDesc(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	adapter := getTestAdapter(t)
	defer adapter.Release()

	desc, err := adapter.GetDesc()
	if err != nil {
		t.Logf("GetDesc failed (may not be supported on newer adapters): %v", err)
		return
	}

	if desc == nil {
		t.Fatal("GetDesc returned nil")
	}

	t.Logf("Adapter Description: %v", desc.Description)
	t.Logf("VendorId: 0x%x", desc.VendorId)
	t.Logf("DeviceId: 0x%x", desc.DeviceId)
}

func TestDXGIAdapter_GetDesc1(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	adapter := getTestAdapter(t)
	defer adapter.Release()

	desc, err := adapter.GetDesc1()
	if err != nil {
		t.Fatalf("GetDesc1 failed: %v", err)
	}

	if desc == nil {
		t.Fatal("GetDesc1 returned nil")
	}

	t.Logf("Adapter Flags: 0x%x", desc.Flags)
	t.Logf("DedicatedVideoMemory: %d", desc.DedicatedVideoMemory)
	t.Logf("SharedSystemMemory: %d", desc.SharedSystemMemory)
}

func TestDXGIAdapter_GetDesc2(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	adapter := getTestAdapter(t)
	defer adapter.Release()

	desc, err := adapter.GetDesc2()
	if err != nil {
		t.Logf("GetDesc2 failed (may not be supported): %v", err)
		return
	}

	if desc == nil {
		t.Fatal("GetDesc2 returned nil")
	}

	t.Logf("GraphicsPreemptionGranularity: %d", desc.GraphicsPreemptionGranularity)
	t.Logf("ComputePreemptionGranularity: %d", desc.ComputePreemptionGranularity)
}

func TestDXGIAdapter_GetDesc3(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	adapter := getTestAdapter(t)
	defer adapter.Release()

	desc, err := adapter.GetDesc3()
	if err != nil {
		t.Logf("GetDesc3 failed (may not be supported): %v", err)
		return
	}

	if desc == nil {
		t.Fatal("GetDesc3 returned nil")
	}

	t.Log("GetDesc3 succeeded")
}

func TestDXGIAdapter_EnumOutputs(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	adapter := getTestAdapter(t)
	defer adapter.Release()

	outputCount := 0
	for i := uint32(0); i < 10; i++ {
		output, err := adapter.EnumOutputs(i)
		if err != nil {
			break
		}
		defer output.Release()

		desc, err := output.GetDesc()
		if err != nil {
			t.Logf("Failed to get output %d description: %v", i, err)
			continue
		}

		outputCount++
		t.Logf("Output %d: DeviceName=%v", i, desc.DeviceName)
	}

	t.Logf("Found %d outputs", outputCount)
}

func TestDXGIAdapter_CheckInterfaceSupport(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	adapter := getTestAdapter(t)
	defer adapter.Release()

	// 测试检查 IUnknown 接口支持
		err := adapter.CheckInterfaceSupport(IID_IUnknown, nil)
	if err != nil {
		t.Logf("CheckInterfaceSupport(IUnknown) failed: %v", err)
	} else {
		t.Log("IUnknown interface is supported")
	}
}

func TestDXGIAdapter_QueryVideoMemoryInfo(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	adapter := getTestAdapter(t)
	defer adapter.Release()

	// 查询本地视频内存信息
	info, err := adapter.QueryVideoMemoryInfo(0, DXGI_MEMORY_SEGMENT_GROUP_LOCAL.DXGI_MEMORY_SEGMENT_GROUP_LOCAL)
	if err != nil {
		t.Logf("QueryVideoMemoryInfo failed (may not be supported): %v", err)
		return
	}

	if info == nil {
		t.Fatal("QueryVideoMemoryInfo returned nil")
	}

	t.Logf("Video Memory Budget: %d", info.Budget)
	t.Logf("Video Memory CurrentUsage: %d", info.CurrentUsage)
	t.Logf("Video Memory AvailableForReservation: %d", info.AvailableForReservation)
}
