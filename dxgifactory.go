package dxgi

import (
	"syscall"
	"unsafe"
)

var (
	IID_IDXGIFactory  = new(GUID)
	IID_IDXGIFactory1 = new(GUID)
	IID_IDXGIFactory2 = new(GUID)
	IID_IDXGIFactory3 = new(GUID)
	IID_IDXGIFactory4 = new(GUID)
	IID_IDXGIFactory5 = new(GUID)
	IID_IDXGIFactory6 = new(GUID)
	IID_IDXGIFactory7 = new(GUID)
)

func init() {
	IID_IDXGIFactory.Setup("7b7166ec-21c7-44ae-b21a-c9ae321ae369")
	IID_IDXGIFactory1.Setup("770aae78-f26f-4dba-a829-253c83d1b387")
	IID_IDXGIFactory2.Setup("50c83a1c-e072-4c48-87b0-3630fa36a6d0")
	IID_IDXGIFactory3.Setup("25483823-cd46-4c7d-86ca-47aa95b837bd")
	IID_IDXGIFactory4.Setup("1bc6ea02-ef36-464f-bf0c-21ca39e5168a")
	IID_IDXGIFactory5.Setup("7632e1f5-ee65-4dca-87fd-84cd75f8838d")
	IID_IDXGIFactory6.Setup("c1b6694f-ff09-44a9-b03c-77900a0a1d17")
	IID_IDXGIFactory7.Setup("a4966eed-76db-44da-84c1-ee9a7afb20a8")
}

type DXGIFactory struct {
	lpVtbl *factoryVtbl
}
type factoryVtbl struct {
	/*** IUnknown methods ***/
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
	/*** IDXGIObject methods ***/
	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	GetPrivateData          uintptr
	GetParent               uintptr
	/*** IDXGIFactory methods ***/
	EnumAdapters          uintptr
	MakeWindowAssociation uintptr
	GetWindowAssociation  uintptr
	CreateSwapChain       uintptr
	CreateSoftwareAdapter uintptr
	/*** IDXGIFactory1 methods ***/
	EnumAdapters1 uintptr
	IsCurrent     uintptr
	/*** IDXGIFactory2 methods ***/
	IsWindowedStereoEnabled       uintptr
	CreateSwapChainForHwnd        uintptr
	CreateSwapChainForCoreWindow  uintptr
	GetSharedResourceAdapterLuid  uintptr
	RegisterOcclusionStatusWindow uintptr
	RegisterStereoStatusEvent     uintptr
	UnregisterStereoStatus        uintptr
	RegisterStereoStatusWindow    uintptr
	RegisterOcclusionStatusEvent  uintptr
	UnregisterOcclusionStatus     uintptr
	CreateSwapChainForComposition uintptr
	/*** IDXGIFactory3 methods ***/
	GetCreationFlags uintptr
	/*** IDXGIFactory4 methods ***/
	EnumAdapterByLuid uintptr
	EnumWarpAdapter   uintptr
	/*** IDXGIFactory5 methods ***/
	CheckFeatureSupport uintptr
	/*** IDXGIFactory6 methods ***/
	EnumAdapterByGpuPreference uintptr
	/*** IDXGIFactory7 methods ***/
	RegisterAdaptersChangedEvent   uintptr
	UnregisterAdaptersChangedEvent uintptr
}

/*** IDXGIObject methods ***/
func (this *DXGIFactory) QueryInterface(riid *GUID, ppvObject *interface{}) error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.QueryInterface,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObject)),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIFactory) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}
func (this *DXGIFactory) Release() uint32 {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.Release,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}

/*** IDXGIObject methods ***/
func (this *DXGIFactory) SetPrivateData(Name *GUID, DataSize uint32, pData *interface{}) error {
	var err error
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.SetPrivateData,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(Name)),
		uintptr(DataSize),
		uintptr(unsafe.Pointer(pData)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIFactory) SetPrivateDataInterface(Name *GUID, pUnknown *IUnknown) error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.SetPrivateDataInterface,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(Name)),
		uintptr(unsafe.Pointer(pUnknown)),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIFactory) GetPrivateData(Name *GUID, pDataSize *uint32) (*interface{}, error) {
	var err error
	var pData *interface{}
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.GetPrivateData,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(Name)),
		uintptr(unsafe.Pointer(pDataSize)),
		uintptr(unsafe.Pointer(pData)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return pData, err
}
func (this *DXGIFactory) GetParent(riid *GUID) (*interface{}, error) {
	var err error
	var ppParent *interface{}
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetParent,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppParent)),
	)
	err = GetError(uint32(ret))
	return ppParent, err
}

/*** IDXGIFactory methods ***/
func (this *DXGIFactory) EnumAdapters(adapteIdx uint32) (*DXGIAdapter, error) {
	var err error
	var adapter uintptr
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.EnumAdapters,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(adapteIdx),
		uintptr(unsafe.Pointer(&adapter)),
	)
	err = GetError(uint32(ret))
	return (*DXGIAdapter)(unsafe.Pointer(adapter)), err
}
func (this *DXGIFactory) MakeWindowAssociation(windowHandle HANDLE, flags uint32) error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.MakeWindowAssociation,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(windowHandle)),
		uintptr(flags),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIFactory) GetWindowAssociation() (HANDLE, error) {
	var err error
	var windowHandle HANDLE
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetWindowAssociation,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&windowHandle)),
		0,
	)
	err = GetError(uint32(ret))
	return windowHandle, err
}
func (this *DXGIFactory) CreateSwapChain(Device *Unknown, Desc *DXGISwapChainDesc) (*DXGISwapChain, error) {
	var err error
	var swapChain uintptr
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.CreateSwapChain,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(Device)),
		uintptr(unsafe.Pointer(Desc)),
		uintptr(unsafe.Pointer(&swapChain)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return (*DXGISwapChain)(unsafe.Pointer(swapChain)), err
}
func (this *DXGIFactory) CreateSoftwareAdapter(Module *interface{}) (*DXGIAdapter, error) {
	var err error
	var adapter uintptr
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.CreateSoftwareAdapter,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(Module)),
		uintptr(unsafe.Pointer(&adapter)),
	)
	err = GetError(uint32(ret))
	return (*DXGIAdapter)(unsafe.Pointer(adapter)), err
}

/*** IDXGIFactory1 methods ***/
func (this *DXGIFactory) EnumAdapters1(adapteIdx uint32) (*DXGIAdapter, error) {
	var err error
	var adapter *DXGIAdapter
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.EnumAdapters1,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(adapteIdx),
		uintptr(unsafe.Pointer(&adapter)),
	)
	err = GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return adapter, nil
}
func (this *DXGIFactory) IsCurrent() bool {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.IsCurrent,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return intToBool(ret)
}

/*** IDXGIFactory2 methods ***/
func (this *DXGIFactory) IsWindowedStereoEnabled() bool {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.IsWindowedStereoEnabled,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return intToBool(ret)
}
func (this *DXGIFactory) CreateSwapChainForHwnd(pDevice *interface{}, hWnd HWND, pDesc *DXGI_SWAP_CHAIN_DESC1, pFullscreenDesc *DXGI_SWAP_CHAIN_FULLSCREEN_DESC, pRestrictToOutput *DXGIOutput) (*DXGISwapChain, error) {
	var err error
	var ppSwapChain *DXGISwapChain
	ret, _, _ := syscall.Syscall9(
		this.lpVtbl.CreateSwapChainForHwnd,
		7,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pDevice)),
		uintptr(unsafe.Pointer(hWnd)),
		uintptr(unsafe.Pointer(pDesc)),
		uintptr(unsafe.Pointer(pFullscreenDesc)),
		uintptr(unsafe.Pointer(pRestrictToOutput)),
		uintptr(unsafe.Pointer(&ppSwapChain)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return ppSwapChain, err
}
func (this *DXGIFactory) CreateSwapChainForCoreWindow(pDevice, pWindow *interface{}, pDesc *DXGI_SWAP_CHAIN_DESC1, pRestrictToOutput *DXGIOutput) (*DXGISwapChain, error) {
	var err error
	var ppSwapChain *DXGISwapChain
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.CreateSwapChainForCoreWindow,
		6,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pDevice)),
		uintptr(unsafe.Pointer(pWindow)),
		uintptr(unsafe.Pointer(pDesc)),
		uintptr(unsafe.Pointer(pRestrictToOutput)),
		uintptr(unsafe.Pointer(&ppSwapChain)),
	)
	err = GetError(uint32(ret))
	return ppSwapChain, err
}
func (this *DXGIFactory) GetSharedResourceAdapterLuid(hResource HANDLE) (*LUID, error) {
	var err error
	var pLuid *LUID
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetSharedResourceAdapterLuid,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(hResource),
		uintptr(unsafe.Pointer(pLuid)),
	)
	err = GetError(uint32(ret))
	return pLuid, err
}
func (this *DXGIFactory) RegisterOcclusionStatusWindow(WindowHandle HWND, wMsg uint32) error {
	var err error
	var pdwCookie *DWORD
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.RegisterOcclusionStatusWindow,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(WindowHandle),
		uintptr(wMsg),
		uintptr(unsafe.Pointer(pdwCookie)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIFactory) RegisterStereoStatusEvent(hEvent HANDLE) (*DWORD, error) {
	var err error
	var pdwCookie *DWORD
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.RegisterStereoStatusEvent,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(hEvent),
		uintptr(unsafe.Pointer(pdwCookie)),
	)
	err = GetError(uint32(ret))
	return pdwCookie, err
}
func (this *DXGIFactory) UnregisterStereoStatus(dwCookie DWORD) error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.UnregisterStereoStatus,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(dwCookie),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIFactory) RegisterStereoStatusWindow(WindowHandle HWND, wMsg uint32) (*DWORD, error) {
	var err error
	var pdwCookie *DWORD
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.RegisterStereoStatusWindow,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(WindowHandle),
		uintptr(wMsg),
		uintptr(unsafe.Pointer(pdwCookie)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return pdwCookie, err
}
func (this *DXGIFactory) RegisterOcclusionStatusEvent(hEvent HANDLE) (*DWORD, error) {
	var err error
	var pdwCookie *DWORD
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.RegisterOcclusionStatusEvent,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(hEvent),
		uintptr(unsafe.Pointer(pdwCookie)),
	)
	err = GetError(uint32(ret))
	return pdwCookie, err
}
func (this *DXGIFactory) UnregisterOcclusionStatus(dwCookie DWORD) error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.UnregisterOcclusionStatus,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(dwCookie),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIFactory) CreateSwapChainForComposition(pDevice *interface{}, pDesc *DXGI_SWAP_CHAIN_DESC1, pRestrictToOutput *DXGIOutput) (*DXGISwapChain, error) {
	var err error
	var ppSwapChain *DXGISwapChain
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.CreateSwapChainForComposition,
		5,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pDevice)),
		uintptr(unsafe.Pointer(pDesc)),
		uintptr(unsafe.Pointer(pRestrictToOutput)),
		uintptr(unsafe.Pointer(&ppSwapChain)),
		0,
	)
	err = GetError(uint32(ret))
	return ppSwapChain, err
}

/*** IDXGIFactory3 methods ***/
func (this *DXGIFactory) GetCreationFlags() uint32 {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetCreationFlags,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}

/*** IDXGIFactory4 methods ***/
func (this *DXGIFactory) EnumAdapterByLuid(AdapterLuid *GUID, riid *GUID) (*DXGIAdapter, error) {
	var err error
	var ppvAdapter *DXGIAdapter
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.EnumAdapterByLuid,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(AdapterLuid)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppvAdapter)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return ppvAdapter, err
}
func (this *DXGIFactory) EnumWarpAdapter(riid *GUID) (*DXGIAdapter, error) {
	var err error
	var ppvAdapter *DXGIAdapter
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.EnumWarpAdapter,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppvAdapter)),
	)
	err = GetError(uint32(ret))
	return ppvAdapter, err
}

/*** IDXGIFactory5 methods ***/
func (this *DXGIFactory) CheckFeatureSupport(Feature uint32, pFeatureSupportData *interface{}, FeatureSupportDataSize uint32) error {
	var err error
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.CheckFeatureSupport,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(Feature),
		uintptr(unsafe.Pointer(pFeatureSupportData)),
		uintptr(FeatureSupportDataSize),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return err
}

/*** IDXGIFactory6 methods ***/
func (this *DXGIFactory) EnumAdapterByGpuPreference(Adapter uint32, GpuPreference uint32, riid *GUID) (*DXGIAdapter, error) {
	var err error
	var ppvAdapter *DXGIAdapter
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.EnumAdapterByGpuPreference,
		5,
		uintptr(unsafe.Pointer(this)),
		uintptr(Adapter),
		uintptr(GpuPreference),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppvAdapter)),
		0,
	)
	err = GetError(uint32(ret))
	return ppvAdapter, err
}

/*** IDXGIFactory7 methods ***/
func (this *DXGIFactory) RegisterAdaptersChangedEvent(hEvent HANDLE) (*DWORD, error) {
	var err error
	var pdwCookie *DWORD
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.RegisterAdaptersChangedEvent,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(hEvent),
		uintptr(unsafe.Pointer(pdwCookie)),
	)
	err = GetError(uint32(ret))
	return pdwCookie, err
}
func (this *DXGIFactory) UnregisterAdaptersChangedEvent(dwCookie DWORD) error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.UnregisterAdaptersChangedEvent,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(dwCookie),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
