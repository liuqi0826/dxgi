package dxgi

import (
	"syscall"
	"unsafe"
)

var IID_IDXGIDevice = new(GUID)
var IID_IDXGIDevice1 = new(GUID)
var IID_IDXGIDevice2 = new(GUID)
var IID_IDXGIDevice3 = new(GUID)
var IID_IDXGIDevice4 = new(GUID)

func init() {
	IID_IDXGIDevice.Setup("54ec77fa-1377-44e6-8c32-88fd5f44c84c")
	IID_IDXGIDevice1.Setup("77db970f-6276-48ba-ba28-070143b4392c")
	IID_IDXGIDevice2.Setup("05008617-fbfd-4051-a790-144884b4f6a9")
	IID_IDXGIDevice3.Setup("6007896c-3244-4afd-bf18-a6d3beda5023")
	IID_IDXGIDevice4.Setup("95B4F95F-D8DA-4CA4-9EE6-3B76D5968A10")
}

type DXGIDevice struct {
	Unknown
	lpVtbl *deviceVtbl
}

type deviceVtbl struct {
	/*** IUnknown methods ***/
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
	/*** IDXGIObject methods ***/
	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	GetPrivateData          uintptr
	GetParent               uintptr
	/*** IDXGIDevice methods ***/
	GetAdapter             uintptr
	CreateSurface          uintptr
	QueryResourceResidency uintptr
	SetGPUThreadPriority   uintptr
	GetGPUThreadPriority   uintptr
	/*** IDXGIDevice1 methods ***/
	SetMaximumFrameLatency uintptr
	GetMaximumFrameLatency uintptr
	/*** IDXGIDevice2 methods ***/
	EnqueueSetEvent  uintptr
	OfferResources   uintptr
	ReclaimResources uintptr
	/*** IDXGIDevice3 methods ***/
	Trim uintptr
	/*** IDXGIDevice4 methods ***/
	OfferResources1   uintptr
	ReclaimResources1 uintptr
}

/*** IDXGIObject methods ***/
func (this *DXGIDevice) QueryInterface(riid *GUID, ppvObject *interface{}) error {
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
func (this *DXGIDevice) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}
func (this *DXGIDevice) Release() uint32 {
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
func (this *DXGIDevice) SetPrivateData(Name *GUID, DataSize uint32, pData *interface{}) error {
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
func (this *DXGIDevice) SetPrivateDataInterface(Name *GUID, pUnknown *IUnknown) error {
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
func (this *DXGIDevice) GetPrivateData(Name *GUID, pDataSize *uint32) (*interface{}, error) {
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
func (this *DXGIDevice) GetParent(riid *GUID) (*interface{}, error) {
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

/*** IDXGIDevice methods ***/
func (this *DXGIDevice) GetAdapter() (*DXGIAdapter, error) {
	var err error
	var pAdapter *DXGIAdapter
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetAdapter,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pAdapter)),
		0,
	)
	err = GetError(uint32(ret))
	return pAdapter, err
}
func (this *DXGIDevice) CreateSurface(pDesc *DXGISharedResource, NumSurfaces uint32, Usage DXGIUsage, pSharedResource *DXGISharedResource) (*DXGISurface, error) {
	var err error
	var ppSurface *DXGISurface
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.CreateSurface,
		6,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pDesc)),
		uintptr(NumSurfaces),
		uintptr(Usage),
		uintptr(unsafe.Pointer(pSharedResource)),
		uintptr(unsafe.Pointer(&ppSurface)),
	)
	err = GetError(uint32(ret))
	return ppSurface, err
}
func (this *DXGIDevice) QueryResourceResidency(ppResources *Unknown) (*uint32, uint32, error) {
	var err error
	var pResidencyStatus *uint32
	var NumResources uint32
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.QueryResourceResidency,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&ppResources)),
		uintptr(unsafe.Pointer(pResidencyStatus)),
		uintptr(NumResources),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return pResidencyStatus, NumResources, err
}
func (this *DXGIDevice) SetGPUThreadPriority(Priority uint32) error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.SetGPUThreadPriority,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(Priority),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIDevice) GetGPUThreadPriority() (*uint32, error) {
	var err error
	var pPriority *uint32
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetGPUThreadPriority,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pPriority)),
		0,
	)
	err = GetError(uint32(ret))
	return pPriority, err
}

/*** IDXGIDevice1 methods ***/
func (this *DXGIDevice) SetMaximumFrameLatency(MaxLatency uint32) error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.SetMaximumFrameLatency,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(MaxLatency),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIDevice) GetMaximumFrameLatency() (*uint32, error) {
	var err error
	var pMaxLatency *uint32
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetMaximumFrameLatency,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pMaxLatency)),
		0,
	)
	err = GetError(uint32(ret))
	return pMaxLatency, err
}

/*** IDXGIDevice2 methods ***/
func (this *DXGIDevice) EnqueueSetEvent(hEvent HANDLE) error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.EnqueueSetEvent,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(hEvent),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIDevice) OfferResources(NumResources uint32, ppResources *DXGIResource, Priority uint32) error {
	var err error
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.OfferResources,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(NumResources),
		uintptr(unsafe.Pointer(&ppResources)),
		uintptr(Priority),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIDevice) ReclaimResources(NumResources uint32, ppResources *DXGIResource, pDiscarded *bool) error {
	var err error
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.ReclaimResources,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(NumResources),
		uintptr(unsafe.Pointer(&ppResources)),
		uintptr(unsafe.Pointer(pDiscarded)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return err
}

/*** IDXGIDevice3 methods ***/
func (this *DXGIDevice) Trim() error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.Trim,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return err
}

/*** IDXGIDevice4 methods ***/
func (this *DXGIDevice) OfferResources1(NumResources uint32, ppResources *DXGIResource, Priority uint32, Flags uint32) error {
	var err error
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.OfferResources1,
		5,
		uintptr(unsafe.Pointer(this)),
		uintptr(NumResources),
		uintptr(unsafe.Pointer(&ppResources)),
		uintptr(Priority),
		uintptr(Flags),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIDevice) ReclaimResources1(NumResources uint32, ppResources *DXGIResource, pResults *uint32) error {
	var err error
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.ReclaimResources1,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(NumResources),
		uintptr(unsafe.Pointer(&ppResources)),
		uintptr(unsafe.Pointer(pResults)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return err
}
