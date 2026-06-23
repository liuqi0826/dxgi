package dxgi

import (
	"testing"
)

func TestGetError(t *testing.T) {
	tests := []struct {
		name     string
		hresult  uint32
		wantErr  bool
		wantMsg  string
	}{
		{
			name:    "Success (S_OK)",
			hresult: 0,
			wantErr: false,
		},
		{
			name:     "DXGI_ERROR_INVALID_CALL",
			hresult:  0x887A0001,
			wantErr:  true,
			wantMsg:  "DXGI_ERROR_INVALID_CALL",
		},
		{
			name:     "DXGI_ERROR_NOT_FOUND",
			hresult:  0x887A0002,
			wantErr:  true,
			wantMsg:  "DXGI_ERROR_NOT_FOUND",
		},
		{
			name:     "DXGI_ERROR_DEVICE_REMOVED",
			hresult:  0x887A0005,
			wantErr:  true,
			wantMsg:  "DXGI_ERROR_DEVICE_REMOVED",
		},
		{
			name:     "DXGI_STATUS_OCCLUDED",
			hresult:  0x087a0001,
			wantErr:  true,
			wantMsg:  "DXGI_STATUS_OCCLUDED",
		},
		{
			name:    "Unknown error code",
			hresult: 0x12345678,
			wantErr: false, // 未知成功 HRESULT 返回 nil
		},
		{
			name:     "Unknown failure HRESULT",
			hresult:  0x80000001,
			wantErr:  true,
			wantMsg:  "HRESULT 0x80000001",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := GetError(tt.hresult)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetError() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err != nil {
				if err.Error() != tt.wantMsg {
					t.Errorf("GetError() error message = %v, want %v", err.Error(), tt.wantMsg)
				}
			}
		})
	}
}

func TestBoolToInt(t *testing.T) {
	tests := []struct {
		name  string
		value bool
		want  uintptr
	}{
		{
			name:  "true",
			value: true,
			want:  1,
		},
		{
			name:  "false",
			value: false,
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := boolToInt(tt.value)
			if got != tt.want {
				t.Errorf("boolToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToBool(t *testing.T) {
	tests := []struct {
		name  string
		value uintptr
		want  bool
	}{
		{
			name:  "zero",
			value: 0,
			want:  false,
		},
		{
			name:  "one",
			value: 1,
			want:  true,
		},
		{
			name:  "non-zero",
			value: 42,
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intToBool(tt.value)
			if got != tt.want {
				t.Errorf("intToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToChar(t *testing.T) {
	tests := []struct {
		name string
		str  string
	}{
		{
			name: "normal string",
			str:  "test",
		},
		{
			name: "empty string",
			str:  "",
		},
		{
			name: "string with null terminator",
			str:  "test\x00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stringToChar(tt.str)
			if result == nil && tt.str != "" {
				t.Error("stringToChar returned nil for non-empty string")
			}
		})
	}
}

func TestCharToString(t *testing.T) {
	// 注意：这个测试需要实际的 C 字符串，在测试环境中可能不可用
	// 这里只测试函数不会 panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("charToString panicked: %v", r)
		}
	}()

	// 测试 nil 指针（应该安全处理）
	result := charToString(nil)
	_ = result
}
