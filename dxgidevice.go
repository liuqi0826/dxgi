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
	ret, _, _ := syscall.SyscallN(this.lpVtbl.QueryInterface, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObject)),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIDevice) AddRef() uint32 {
	ret, _, _ := syscall.SyscallN(this.lpVtbl.AddRef, uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}
func (this *DXGIDevice) Release() uint32 {
	ret, _, _ := syscall.SyscallN(this.lpVtbl.Release, uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}

/*** IDXGIObject methods ***/
func (this *DXGIDevice) SetPrivateData(Name *GUID, DataSize uint32, pData *interface{}) error {
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
func (this *DXGIDevice) SetPrivateDataInterface(Name *GUID, pUnknown *IUnknown) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.SetPrivateDataInterface, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(Name)),
		uintptr(unsafe.Pointer(pUnknown)),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIDevice) GetPrivateData(Name *GUID, pDataSize *uint32) (*interface{}, error) {
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
func (this *DXGIDevice) GetParent(riid *GUID) (*interface{}, error) {
	var err error
	var ppParent *interface{}
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetParent, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppParent)),
	)
	err = GetError(uint32(ret))
	return ppParent, err
}

/*** IDXGIDevice methods ***/
func (this *DXGIDevice) GetAdapter() (*DXGIAdapter, error) {
	var err error
	var ppAdapter *DXGIAdapter
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetAdapter, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&ppAdapter)),
		0,
	)
	err = GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return ppAdapter, nil
}
func (this *DXGIDevice) CreateSurface(pDesc *DXGISharedResource, NumSurfaces uint32, Usage DXGIUsage, pSharedResource *DXGISharedResource) (*DXGISurface, error) {
	var err error
	var ppSurface *DXGISurface
	ret, _, _ := syscall.SyscallN(this.lpVtbl.CreateSurface, uintptr(unsafe.Pointer(this)),
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
	ret, _, _ := syscall.SyscallN(this.lpVtbl.QueryResourceResidency, uintptr(unsafe.Pointer(this)),
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
	ret, _, _ := syscall.SyscallN(this.lpVtbl.SetGPUThreadPriority, uintptr(unsafe.Pointer(this)),
		uintptr(Priority),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIDevice) GetGPUThreadPriority() (int32, error) {
	var err error
	var priority int32
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetGPUThreadPriority, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&priority)),
		0,
	)
	err = GetError(uint32(ret))
	if err != nil {
		return 0, err
	}
	return priority, nil
}

/*** IDXGIDevice1 methods ***/
func (this *DXGIDevice) SetMaximumFrameLatency(MaxLatency uint32) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.SetMaximumFrameLatency, uintptr(unsafe.Pointer(this)),
		uintptr(MaxLatency),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIDevice) GetMaximumFrameLatency() (uint32, error) {
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

/*** IDXGIDevice2 methods ***/
func (this *DXGIDevice) EnqueueSetEvent(hEvent HANDLE) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.EnqueueSetEvent, uintptr(unsafe.Pointer(this)),
		uintptr(hEvent),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIDevice) OfferResources(NumResources uint32, ppResources *DXGIResource, Priority uint32) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.OfferResources, uintptr(unsafe.Pointer(this)),
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
	ret, _, _ := syscall.SyscallN(this.lpVtbl.ReclaimResources, uintptr(unsafe.Pointer(this)),
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
	ret, _, _ := syscall.SyscallN(this.lpVtbl.Trim, uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return err
}

/*** IDXGIDevice4 methods ***/
func (this *DXGIDevice) OfferResources1(NumResources uint32, ppResources *DXGIResource, Priority uint32, Flags uint32) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.OfferResources1, uintptr(unsafe.Pointer(this)),
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
	ret, _, _ := syscall.SyscallN(this.lpVtbl.ReclaimResources1, uintptr(unsafe.Pointer(this)),
		uintptr(NumResources),
		uintptr(unsafe.Pointer(&ppResources)),
		uintptr(unsafe.Pointer(pResults)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return err
}
