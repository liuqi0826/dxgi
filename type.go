package dxgi

import "C"
import "unsafe"

type (
	ATOM            = uint16
	BOOL            = int32
	COLORREF        = uint32
	DWM_FRAME_COUNT = uint64
	DWORD           = uint32
	HACCEL          = HANDLE
	HANDLE          = uintptr
	HBITMAP         = HANDLE
	HBRUSH          = HANDLE
	HCURSOR         = HANDLE
	HDC             = HANDLE
	HDROP           = HANDLE
	HDWP            = HANDLE
	HENHMETAFILE    = HANDLE
	HFONT           = HANDLE
	HGDIOBJ         = HANDLE
	HGLOBAL         = HANDLE
	HGLRC           = HANDLE
	HHOOK           = HANDLE
	HICON           = HANDLE
	HIMAGELIST      = HANDLE
	HINSTANCE       = HANDLE
	HKEY            = HANDLE
	HKL             = HANDLE
	HMENU           = HANDLE
	HMODULE         = HANDLE
	HMONITOR        = HANDLE
	HPEN            = HANDLE
	HRESULT         = int32
	HRGN            = HANDLE
	HRSRC           = HANDLE
	HTHUMBNAIL      = HANDLE
	HWND            = HANDLE
	LPARAM          = uintptr
	LPCVOID         = unsafe.Pointer
	LRESULT         = uintptr
	PVOID           = unsafe.Pointer
	QPC_TIME        = uint64
	ULONG_PTR       = uintptr
	WPARAM          = uintptr
	HRAWINPUT       = HANDLE
	SIZE_T          = uint32
)

//++++++++++++++++++++  ++++++++++++++++++++
type LargeInteger unsafe.Pointer
type LUID struct {
	LowPart  uint32
	HighPart int32
}
type LARGE_INTEGER struct {
	LowPart  uint32
	HighPart int32
}
type RECT struct {
	Left   int64
	Top    int64
	Right  int64
	Bottom int64
}
type POINT struct {
	X int64
	Y int64
}

//++++++++++++++++++++  ++++++++++++++++++++
type DXGI_DEBUG_ID GUID
type DXGIUsage uint32
type DXGIRGB struct {
	Red   float32
	Green float32
	Blue  float32
}
type DXGIGammaControl struct {
	Scale      DXGIRGB
	Offset     DXGIRGB
	GammaCurve [1025]DXGIRGB
}
type DXGIGammaControlCapabilities struct {
	ScaleAndOffsetSupported bool
	MaxConvertedValue       float32
	MinConvertedValue       float32
	NumGammaControlPoints   uint32
	ControlPointPositions   [1025]float32
}
type DXGIRational struct {
	Numerator   uint32
	Denominator uint32
}
type DXGIModeDesc struct {
	Width       uint32
	Height      uint32
	RefreshRate DXGIRational
}
type DXGISampleDesc struct {
	Count   uint32
	Quality uint32
}

type SECURITY_ATTRIBUTES struct {
	Length              DWORD
	lSecurityDescriptor *interface{}
	InheritHandle       bool
}
