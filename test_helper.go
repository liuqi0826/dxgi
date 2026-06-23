package dxgi

import (
	"runtime"
	"testing"
)

// skipIfNotWindows 如果不是 Windows 系统则跳过测试
func skipIfNotWindows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("Skipping test on non-Windows platform")
	}
}

// getTestFactory 获取测试用的工厂对象
func getTestFactory(t *testing.T) *DXGIFactory {
	skipIfNotWindows(t)

	factory, err := CreateDXGIFactory1(IID_IDXGIFactory1)
	if err != nil {
		t.Fatalf("CreateDXGIFactory1 failed: %v", err)
	}
	return factory
}

// getTestAdapter 获取测试用的适配器对象
func getTestAdapter(t *testing.T) *DXGIAdapter {
	factory := getTestFactory(t)
	defer factory.Release()

	adapter, err := factory.EnumAdapters1(0)
	if err != nil {
		t.Fatalf("EnumAdapters1(0) failed: %v", err)
	}
	return adapter
}

// getTestOutput 获取测试用的输出对象（遍历所有适配器）
func getTestOutput(t *testing.T) *DXGIOutput {
	factory := getTestFactory(t)
	defer factory.Release()

	for i := uint32(0); i < 10; i++ {
		adapter, err := factory.EnumAdapters1(i)
		if err != nil {
			break
		}

		var output *DXGIOutput
		for j := uint32(0); j < 10; j++ {
			output, err = adapter.EnumOutputs(j)
			if err != nil {
				output = nil
				break
			}
			break
		}
		adapter.Release()
		if output != nil {
			return output
		}
	}
	t.Skip("No display outputs found on any adapter")
	return nil
}
