package dxgi

import (
	"syscall"
	"unsafe"
)

var (
	IID_IDXGISwapChain  = new(GUID)
	IID_IDXGISwapChain1 = new(GUID)
	IID_IDXGISwapChain2 = new(GUID)
	IID_IDXGISwapChain3 = new(GUID)
	IID_IDXGISwapChain4 = new(GUID)
)

func init() {
	IID_IDXGISwapChain.Setup("310d36a0-d2e7-4c0a-aa04-6a9d23b8886a")
	IID_IDXGISwapChain1.Setup("790a45f7-0d42-4876-983a-0a55cfe6f4aa")
	IID_IDXGISwapChain2.Setup("a8be2ac4-199f-4946-b331-79599fb98de7")
	IID_IDXGISwapChain3.Setup("94d99bdb-f1f8-4ab0-b236-7da0170edab1")
	IID_IDXGISwapChain4.Setup("3D585D5A-BD4A-489E-B1F4-3DBCB6452FFB")
}

type DXGISwapChain struct {
	lpVtbl *swapChainVtbl
}
type swapChainVtbl struct {
	/*** IUnknown methods ***/
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
	/*** IDXGIObject methods ***/
	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	GetPrivateData          uintptr
	GetParent               uintptr
	/*** IDXGISwapChain methods ***/
	Present             uintptr
	GetBuffer           uintptr
	SetFullscreenState  uintptr
	GetFullscreenState  uintptr
	GetDesc             uintptr
	ResizeBuffers       uintptr
	ResizeTarget        uintptr
	GetContainingOutput uintptr
	GetFrameStatistics  uintptr
	GetLastPresentCount uintptr
	/*** IDXGISwapChain1 methods ***/
	GetDesc1                 uintptr
	GetFullscreenDesc        uintptr
	GetHwnd                  uintptr
	GetCoreWindow            uintptr
	Present1                 uintptr
	IsTemporaryMonoSupported uintptr
	GetRestrictToOutput      uintptr
	SetBackgroundColor       uintptr
	GetBackgroundColor       uintptr
	SetRotation              uintptr
	GetRotation              uintptr
	/*** IDXGISwapChain2 methods ***/
	SetSourceSize                 uintptr
	GetSourceSize                 uintptr
	SetMaximumFrameLatency        uintptr
	GetMaximumFrameLatency        uintptr
	GetFrameLatencyWaitableObject uintptr
	SetMatrixTransform            uintptr
	GetMatrixTransform            uintptr
	/*** IDXGISwapChain3 methods ***/
	GetCurrentBackBufferIndex uintptr
	CheckColorSpaceSupport    uintptr
	SetColorSpace1            uintptr
	ResizeBuffers1            uintptr
	/*** IDXGISwapChain4 methods ***/
	SetHDRMetaData uintptr
}

/*** IDXGIObject methods ***/
func (this *DXGISwapChain) QueryInterface(riid *GUID, ppvObject *interface{}) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.QueryInterface, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObject)),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGISwapChain) AddRef() uint32 {
	ret, _, _ := syscall.SyscallN(this.lpVtbl.AddRef, uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}
func (this *DXGISwapChain) Release() uint32 {
	ret, _, _ := syscall.SyscallN(this.lpVtbl.Release, uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}

/*** IDXGIObject methods ***/
func (this *DXGISwapChain) SetPrivateData(Name *GUID, DataSize uint32, pData *interface{}) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.SetPrivateData, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(Name)),
		uintptr(DataSize),
		uintptr(unsafe.Pointer(pData)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGISwapChain) SetPrivateDataInterface(Name *GUID, pUnknown *IUnknown) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.SetPrivateDataInterface, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(Name)),
		uintptr(unsafe.Pointer(pUnknown)),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGISwapChain) GetPrivateData(Name *GUID, pDataSize *uint32) (*interface{}, error) {
	var err error
	var pData *interface{}
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetPrivateData, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(Name)),
		uintptr(unsafe.Pointer(pDataSize)),
		uintptr(unsafe.Pointer(pData)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return pData, err
}
func (this *DXGISwapChain) GetParent(riid *GUID) (*interface{}, error) {
	var err error
	var ppParent *interface{}
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetParent, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppParent)),
	)
	err = GetError(uint32(ret))
	return ppParent, err
}

/*** IDXGISwapChain methods ***/
func (this *DXGISwapChain) Present(SyncInterval, Flags uint32) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.Present, uintptr(unsafe.Pointer(this)),
		uintptr(SyncInterval),
		uintptr(Flags),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGISwapChain) GetBuffer(Buffer uint32, riid *GUID) (interface{}, error) {
	var err error
	var ppSurface interface{}
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetBuffer, uintptr(unsafe.Pointer(this)),
		uintptr(Buffer),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppSurface)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return ppSurface, err
}
func (this *DXGISwapChain) SetFullscreenState(Fullscreen bool, pTarget *DXGIOutput) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.SetFullscreenState, uintptr(unsafe.Pointer(this)),
		boolToInt(Fullscreen),
		uintptr(unsafe.Pointer(pTarget)),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGISwapChain) GetFullscreenState(pFullscreen *bool) (*DXGIOutput, error) {
	var err error
	var ppTarget *DXGIOutput
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetFullscreenState, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pFullscreen)),
		uintptr(unsafe.Pointer(&ppTarget)),
	)
	err = GetError(uint32(ret))
	return ppTarget, err
}
func (this *DXGISwapChain) GetDesc() (*DXGISwapChainDesc, error) {
	var err error
	var desc DXGISwapChainDesc
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetDesc, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&desc)),
		0,
	)
	err = GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return &desc, nil
}
func (this *DXGISwapChain) ResizeBuffers(BufferCount, Width, Height uint32, NewFormat uint32, SwapChainFlags uint32) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.ResizeBuffers, uintptr(unsafe.Pointer(this)),
		uintptr(BufferCount),
		uintptr(Width),
		uintptr(Height),
		uintptr(NewFormat),
		uintptr(SwapChainFlags),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGISwapChain) ResizeTarget(pNewTargetParameters *DXGIModeDesc) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.ResizeTarget, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pNewTargetParameters)),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGISwapChain) GetContainingOutput() (*DXGIOutput, error) {
	var err error
	var ppOutput *DXGIOutput
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetContainingOutput, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&ppOutput)),
		0,
	)
	err = GetError(uint32(ret))
	return ppOutput, err
}
func (this *DXGISwapChain) GetFrameStatistics() (*DXGI_FRAME_STATISTICS, error) {
	var err error
	var stats DXGI_FRAME_STATISTICS
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetFrameStatistics, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&stats)),
		0,
	)
	err = GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return &stats, nil
}
func (this *DXGISwapChain) GetLastPresentCount() (uint32, error) {
	var err error
	var lastPresentCount uint32
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetLastPresentCount, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&lastPresentCount)),
		0,
	)
	err = GetError(uint32(ret))
	if err != nil {
		return 0, err
	}
	return lastPresentCount, nil
}

/*** IDXGISwapChain1 methods ***/
func (this *DXGISwapChain) GetDesc1() (*DXGI_SWAP_CHAIN_DESC1, error) {
	var err error
	var desc DXGI_SWAP_CHAIN_DESC1
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetDesc1, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&desc)),
		0,
	)
	err = GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return &desc, nil
}
func (this *DXGISwapChain) GetFullscreenDesc() (*DXGI_SWAP_CHAIN_FULLSCREEN_DESC, error) {
	var err error
	var desc DXGI_SWAP_CHAIN_FULLSCREEN_DESC
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetFullscreenDesc, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&desc)),
		0,
	)
	err = GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return &desc, nil
}
func (this *DXGISwapChain) GetHwnd() (HWND, error) {
	var err error
	var hwnd HWND
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetHwnd, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&hwnd)),
		0,
	)
	err = GetError(uint32(ret))
	if err != nil {
		return 0, err
	}
	return hwnd, nil
}
func (this *DXGISwapChain) GetCoreWindow(refiid *GUID) (*interface{}, error) {
	var err error
	var ppUnk *interface{}
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetCoreWindow, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(refiid)),
		uintptr(unsafe.Pointer(&ppUnk)),
	)
	err = GetError(uint32(ret))
	return ppUnk, err
}
func (this *DXGISwapChain) Present1(SyncInterval, PresentFlags uint32, pPresentParameters *DXGI_PRESENT_PARAMETERS) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.Present1, uintptr(unsafe.Pointer(this)),
		uintptr(SyncInterval),
		uintptr(PresentFlags),
		uintptr(unsafe.Pointer(pPresentParameters)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGISwapChain) IsTemporaryMonoSupported() bool {
	ret, _, _ := syscall.SyscallN(this.lpVtbl.IsTemporaryMonoSupported, uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return intToBool(ret)
}
func (this *DXGISwapChain) GetRestrictToOutput() (*DXGIOutput, error) {
	var err error
	var ppRestrictToOutput *DXGIOutput
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetRestrictToOutput, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&ppRestrictToOutput)),
		0,
	)
	err = GetError(uint32(ret))
	return ppRestrictToOutput, err
}
func (this *DXGISwapChain) SetBackgroundColor(pColor *DXGI_RGBA) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.SetBackgroundColor, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pColor)),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGISwapChain) GetBackgroundColor() (*DXGI_RGBA, error) {
	var err error
	var color DXGI_RGBA
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetBackgroundColor, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&color)),
		0,
	)
	err = GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return &color, nil
}
func (this *DXGISwapChain) SetRotation(Rotation uint32) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.SetRotation, uintptr(unsafe.Pointer(this)),
		uintptr(Rotation),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGISwapChain) GetRotation() (uint32, error) {
	var err error
	var rotation uint32
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetRotation, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&rotation)),
		0,
	)
	err = GetError(uint32(ret))
	if err != nil {
		return 0, err
	}
	return rotation, nil
}

/*** IDXGISwapChain2 methods ***/
func (this *DXGISwapChain) SetSourceSize(Width, Height uint32) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.SetSourceSize, uintptr(unsafe.Pointer(this)),
		uintptr(Width),
		uintptr(Height),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGISwapChain) GetSourceSize() (Width, Height uint32, err error) {
	var pWidth, pHeight uint32
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetSourceSize, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&pWidth)),
		uintptr(unsafe.Pointer(&pHeight)),
		0,
		0,
		0,
	)
	err = GetError(uint32(ret))
	if err != nil {
		return 0, 0, err
	}
	return pWidth, pHeight, nil
}
func (this *DXGISwapChain) SetMaximumFrameLatency(MaxLatency uint32) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.SetMaximumFrameLatency, uintptr(unsafe.Pointer(this)),
		uintptr(MaxLatency),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGISwapChain) GetMaximumFrameLatency() (uint32, error) {
	var err error
	var maxLatency uint32
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetMaximumFrameLatency, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&maxLatency)),
		0,
	)
	err = GetError(uint32(ret))
	if err != nil {
		return 0, err
	}
	return maxLatency, nil
}
func (this *DXGISwapChain) GetFrameLatencyWaitableObject() (HANDLE, error) {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetFrameLatencyWaitableObject, uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	if err != nil {
		return 0, err
	}
	return HANDLE(ret), nil
}
func (this *DXGISwapChain) SetMatrixTransform(pMatrix *DXGI_MATRIX_3X2_F) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.SetMatrixTransform, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pMatrix)),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGISwapChain) GetMatrixTransform() (*DXGI_MATRIX_3X2_F, error) {
	var err error
	var matrix DXGI_MATRIX_3X2_F
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetMatrixTransform, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&matrix)),
		0,
	)
	err = GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return &matrix, nil
}

/*** IDXGISwapChain3 methods ***/
func (this *DXGISwapChain) GetCurrentBackBufferIndex() uint32 {
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetCurrentBackBufferIndex, uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}
func (this *DXGISwapChain) CheckColorSpaceSupport(ColorSpace uint32) (uint32, error) {
	var err error
	var colorSpaceSupport uint32
	ret, _, _ := syscall.SyscallN(this.lpVtbl.CheckColorSpaceSupport, uintptr(unsafe.Pointer(this)),
		uintptr(ColorSpace),
		uintptr(unsafe.Pointer(&colorSpaceSupport)),
	)
	err = GetError(uint32(ret))
	if err != nil {
		return 0, err
	}
	return colorSpaceSupport, nil
}
func (this *DXGISwapChain) SetColorSpace1(ColorSpace uint32) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.SetColorSpace1, uintptr(unsafe.Pointer(this)),
		uintptr(ColorSpace),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGISwapChain) ResizeBuffers1(BufferCount, Width, Height, Format, SwapChainFlags uint32, pCreationNodeMask *uint32, ppPresentQueue **interface{}) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.ResizeBuffers1, uintptr(unsafe.Pointer(this)),
		uintptr(BufferCount),
		uintptr(Width),
		uintptr(Height),
		uintptr(Format),
		uintptr(SwapChainFlags),
		uintptr(unsafe.Pointer(pCreationNodeMask)),
		uintptr(unsafe.Pointer(ppPresentQueue)),
		0,
	)
	err = GetError(uint32(ret))
	return err
}

/*** IDXGISwapChain4 methods ***/
func (this *DXGISwapChain) SetHDRMetaData(Type uint32, Size uint32, pMetaData *interface{}) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.SetHDRMetaData, uintptr(unsafe.Pointer(this)),
		uintptr(Type),
		uintptr(Size),
		uintptr(unsafe.Pointer(pMetaData)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return err
}
