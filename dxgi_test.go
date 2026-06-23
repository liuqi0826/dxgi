package dxgi

import (
	"testing"
)

func TestCreateDXGIFactory1(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	factory, err := CreateDXGIFactory1(IID_IDXGIFactory1)
	if err != nil {
		t.Fatalf("CreateDXGIFactory1 failed: %v", err)
	}
	if factory == nil {
		t.Fatal("CreateDXGIFactory1 returned nil factory")
	}
	defer factory.Release()

	// 测试 IsCurrent
	isCurrent := factory.IsCurrent()
	t.Logf("Factory IsCurrent: %v", isCurrent)
}

func TestCreateDXGIFactory2(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	factory, err := CreateDXGIFactory2(0, IID_IDXGIFactory2)
	if err != nil {
		t.Fatalf("CreateDXGIFactory2 failed: %v", err)
	}
	if factory == nil {
		t.Fatal("CreateDXGIFactory2 returned nil factory")
	}
	defer factory.Release()

	// 测试 GetCreationFlags
	flags := factory.GetCreationFlags()
	t.Logf("Factory creation flags: 0x%x", flags)
}

func TestEnumAdapters1(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	factory, err := CreateDXGIFactory1(IID_IDXGIFactory1)
	if err != nil {
		t.Fatalf("CreateDXGIFactory1 failed: %v", err)
	}
	defer factory.Release()

	// 枚举适配器
	adapterCount := 0
	for i := uint32(0); i < 10; i++ {
		adapter, err := factory.EnumAdapters1(i)
		if err != nil {
			break
		}
		defer adapter.Release()

		desc, err := adapter.GetDesc1()
		if err != nil {
			t.Logf("Failed to get adapter %d description: %v", i, err)
			continue
		}

		adapterCount++
		t.Logf("Adapter %d: Flags=0x%x", i, desc.Flags)
	}

	if adapterCount == 0 {
		t.Fatal("No adapters found")
	}
	t.Logf("Found %d adapters", adapterCount)
}
