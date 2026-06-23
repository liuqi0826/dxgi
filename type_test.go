package dxgi

import (
	"testing"
)

func TestRECT(t *testing.T) {
	rect := RECT{
		Left:   10,
		Top:    20,
		Right:  100,
		Bottom: 200,
	}

	if rect.Left != 10 {
		t.Errorf("RECT Left = %d, want 10", rect.Left)
	}
	if rect.Top != 20 {
		t.Errorf("RECT Top = %d, want 20", rect.Top)
	}
	if rect.Right != 100 {
		t.Errorf("RECT Right = %d, want 100", rect.Right)
	}
	if rect.Bottom != 200 {
		t.Errorf("RECT Bottom = %d, want 200", rect.Bottom)
	}
}

func TestPOINT(t *testing.T) {
	point := POINT{
		X: 50,
		Y: 75,
	}

	if point.X != 50 {
		t.Errorf("POINT X = %d, want 50", point.X)
	}
	if point.Y != 75 {
		t.Errorf("POINT Y = %d, want 75", point.Y)
	}
}

func TestLUID(t *testing.T) {
	luid := LUID{
		LowPart:  0x12345678,
		HighPart: 0x12345678, // 使用有效的 int32 值
	}

	if luid.LowPart != 0x12345678 {
		t.Errorf("LUID LowPart = 0x%x, want 0x12345678", luid.LowPart)
	}
	if luid.HighPart != 0x12345678 {
		t.Errorf("LUID HighPart = 0x%x, want 0x12345678", luid.HighPart)
	}
}

func TestLARGE_INTEGER(t *testing.T) {
	largeInt := LARGE_INTEGER{
		LowPart:  0x12345678,
		HighPart: 0x12345678, // 使用有效的 int32 值
	}

	if largeInt.LowPart != 0x12345678 {
		t.Errorf("LARGE_INTEGER LowPart = 0x%x, want 0x12345678", largeInt.LowPart)
	}
	if largeInt.HighPart != 0x12345678 {
		t.Errorf("LARGE_INTEGER HighPart = 0x%x, want 0x12345678", largeInt.HighPart)
	}
}
