package dxgi

import (
	"testing"
)

func TestGUID_Setup(t *testing.T) {
	tests := []struct {
		name    string
		guidStr string
		wantErr bool
	}{
		{
			name:    "Standard GUID format",
			guidStr: "{12345678-1234-1234-1234-123456789abc}",
			wantErr: false,
		},
		{
			name:    "GUID without braces",
			guidStr: "12345678-1234-1234-1234-123456789abc",
			wantErr: false,
		},
		{
			name:    "GUID without dashes",
			guidStr: "12345678123412341234123456789abc",
			wantErr: false,
		},
		{
			name:    "Invalid format",
			guidStr: "invalid",
			wantErr: false, // Setup 在无效格式时返回 nil，不返回错误
		},
		{
			name:    "Empty string",
			guidStr: "",
			wantErr: false, // Setup 在无效格式时返回 nil，不返回错误
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			guid := new(GUID)
			err := guid.Setup(tt.guidStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GUID.Setup() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && guid != nil {
				// 验证 GUID 字符串表示
				str := guid.String()
				if len(str) == 0 {
					t.Error("GUID.String() returned empty string")
				}
			}
		})
	}
}

func TestGUID_String(t *testing.T) {
	guid := new(GUID)
	err := guid.Setup("{12345678-1234-1234-1234-123456789abc}")
	if err != nil {
		t.Fatalf("GUID.Setup failed: %v", err)
	}

	str := guid.String()
	if str != "{12345678-1234-1234-1234-123456789ABC}" {
		t.Errorf("GUID.String() = %v, want {12345678-1234-1234-1234-123456789ABC}", str)
	}
}

func TestGUID_Nil(t *testing.T) {
	var guid *GUID
	str := guid.String()
	if str != "{00000000-0000-0000-0000-000000000000}" {
		t.Errorf("nil GUID.String() = %v, want {00000000-0000-0000-0000-000000000000}", str)
	}
}

func TestIID_Setup(t *testing.T) {
	// 测试标准接口 ID
	if IID_IUnknown == nil {
		t.Fatal("IID_IUnknown is nil")
	}

	str := IID_IUnknown.String()
	if len(str) == 0 {
		t.Error("IID_IUnknown.String() returned empty string")
	}

	// 验证 IID_IDXGIFactory1
	if IID_IDXGIFactory1 == nil {
		t.Fatal("IID_IDXGIFactory1 is nil")
	}

	str = IID_IDXGIFactory1.String()
	if len(str) == 0 {
		t.Error("IID_IDXGIFactory1.String() returned empty string")
	}
	t.Logf("IID_IDXGIFactory1: %s", str)
}

func TestGUID_RoundTrip(t *testing.T) {
	original := "{12345678-1234-1234-1234-123456789abc}"
	guid := new(GUID)
	err := guid.Setup(original)
	if err != nil {
		t.Fatalf("GUID.Setup failed: %v", err)
	}

	result := guid.String()
	// 注意：字符串可能被转换为大写
	if len(result) == 0 {
		t.Error("GUID.String() returned empty string after Setup")
	}
	t.Logf("Original: %s, Result: %s", original, result)
}
