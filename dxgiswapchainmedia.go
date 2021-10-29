package dxgi

import (
	"syscall"
	"unsafe"
)

var IID_IDXGISwapChainMedia = new(GUID)

func init() {
	IID_IDXGISwapChainMedia.Setup("dd95b90b-f05f-4f6a-bd65-25bfb264bd84")
}

type DXGISwapChainMedia struct {
	Unknown
	lpVtbl *swapChainMediaVtbl
}

type swapChainMediaVtbl struct {
	/*** IUnknown methods ***/
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
	/*** IDXGIObject methods ***/
	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	GetPrivateData          uintptr
	GetParent               uintptr
	/*** IDXGISwapChainMedia methods ***/
	GetFrameStatisticsMedia     uintptr
	SetPresentDuration          uintptr
	CheckPresentDurationSupport uintptr
}

/*** IDXGIObject methods ***/
func (this *DXGISwapChainMedia) QueryInterface(riid *GUID, ppvObject *interface{}) error {
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
func (this *DXGISwapChainMedia) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}
func (this *DXGISwapChainMedia) Release() uint32 {
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
func (this *DXGISwapChainMedia) SetPrivateData(Name *GUID, DataSize uint32, pData *interface{}) error {
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
func (this *DXGISwapChainMedia) SetPrivateDataInterface(Name *GUID, pUnknown *IUnknown) error {
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
func (this *DXGISwapChainMedia) GetPrivateData(Name *GUID, pDataSize *uint32) (*interface{}, error) {
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
func (this *DXGISwapChainMedia) GetParent(riid *GUID) (*interface{}, error) {
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

/*** IDXGISwapChainMedia methods ***/
func (this *DXGISwapChainMedia) GetFrameStatisticsMedia() (*DXGI_FRAME_STATISTICS_MEDIA, error) {
	var err error
	var pStats *DXGI_FRAME_STATISTICS_MEDIA
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetFrameStatisticsMedia,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pStats)),
		0,
	)
	err = GetError(uint32(ret))
	return pStats, err
}
func (this *DXGISwapChainMedia) SetPresentDuration(Duration uint32) error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.SetPresentDuration,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(Duration),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGISwapChainMedia) CheckPresentDurationSupport(DesiredPresentDuration uint32) (*uint32, *uint32, error) {
	var err error
	var pClosestSmallerPresentDuration, pClosestLargerPresentDuration *uint32
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.CheckPresentDurationSupport,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(DesiredPresentDuration),
		uintptr(unsafe.Pointer(pClosestSmallerPresentDuration)),
		uintptr(unsafe.Pointer(pClosestLargerPresentDuration)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return pClosestSmallerPresentDuration, pClosestLargerPresentDuration, err
}
