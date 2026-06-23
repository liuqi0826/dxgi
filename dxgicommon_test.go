package dxgi

import (
	"testing"
)

func TestDXGI_RATIONAL(t *testing.T) {
	rational := DXGI_RATIONAL{
		Numerator:   60,
		Denominator: 1,
	}

	if rational.Numerator != 60 {
		t.Errorf("DXGI_RATIONAL Numerator = %d, want 60", rational.Numerator)
	}
	if rational.Denominator != 1 {
		t.Errorf("DXGI_RATIONAL Denominator = %d, want 1", rational.Denominator)
	}
}

func TestDXGI_SAMPLE_DESC(t *testing.T) {
	sampleDesc := DXGI_SAMPLE_DESC{
		Count:   4,
		Quality: 0,
	}

	if sampleDesc.Count != 4 {
		t.Errorf("DXGI_SAMPLE_DESC Count = %d, want 4", sampleDesc.Count)
	}
	if sampleDesc.Quality != 0 {
		t.Errorf("DXGI_SAMPLE_DESC Quality = %d, want 0", sampleDesc.Quality)
	}
}

func TestDXGI_COLOR_SPACE_TYPE(t *testing.T) {
	// 测试颜色空间类型常量
	if DXGI_COLOR_SPACE_TYPE.DXGI_COLOR_SPACE_RGB_FULL_G22_NONE_P709 != 0 {
		t.Error("DXGI_COLOR_SPACE_RGB_FULL_G22_NONE_P709 should be 0")
	}
}
