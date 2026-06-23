package dxgi

import (
	"syscall"
	"unsafe"
)

var IID_IDXGIResource = new(GUID)
var IID_IDXGIResource1 = new(GUID)

func init() {
	IID_IDXGIResource.Setup("035f3ab4-482e-4e50-b41f-8a7f8bd8960b")
	IID_IDXGIResource1.Setup("30961379-4609-4a41-998e-54fe567ee0c1")
}

type DXGIResource struct {
	lpVtbl *resourceVtbl
}

type resourceVtbl struct {
	/*** IUnknown methods ***/
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
	/*** IDXGIObject methods ***/
	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	GetPrivateData          uintptr
	GetParent               uintptr
	/*** IDXGIResource methods ***/
	GetSharedHandle     uintptr
	GetUsage            uintptr
	SetEvictionPriority uintptr
	GetEvictionPriority uintptr
	/*** IDXGIResource1 methods ***/
	CreateSubresourceSurface uintptr
	CreateSharedHandle       uintptr
}

/*** IDXGIObject methods ***/
func (this *DXGIResource) QueryInterface(riid *GUID, ppvObject *interface{}) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.QueryInterface, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObject)),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIResource) AddRef() uint32 {
	ret, _, _ := syscall.SyscallN(this.lpVtbl.AddRef, uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}
func (this *DXGIResource) Release() uint32 {
	ret, _, _ := syscall.SyscallN(this.lpVtbl.Release, uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}

/*** IDXGIObject methods ***/
func (this *DXGIResource) SetPrivateData(Name *GUID, DataSize uint32, pData *interface{}) error {
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
func (this *DXGIResource) SetPrivateDataInterface(Name *GUID, pUnknown *IUnknown) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.SetPrivateDataInterface, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(Name)),
		uintptr(unsafe.Pointer(pUnknown)),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIResource) GetPrivateData(Name *GUID, pDataSize *uint32) (*interface{}, error) {
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
func (this *DXGIResource) GetParent(riid *GUID) (*interface{}, error) {
	var err error
	var ppParent *interface{}
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetParent, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppParent)),
	)
	err = GetError(uint32(ret))
	return ppParent, err
}

/*** IDXGIResource methods ***/
func (this *DXGIResource) GetSharedHandle() (*HANDLE, error) {
	var err error
	var pSharedHandle *HANDLE
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetSharedHandle, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pSharedHandle)),
		0,
	)
	err = GetError(uint32(ret))
	return pSharedHandle, err
}
func (this *DXGIResource) GetUsage() (DXGI_USAGE, error) {
	var err error
	var usage DXGI_USAGE
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetUsage, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&usage)),
		0,
	)
	err = GetError(uint32(ret))
	if err != nil {
		return 0, err
	}
	return usage, nil
}
func (this *DXGIResource) SetEvictionPriority(EvictionPriority uint32) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.SetEvictionPriority, uintptr(unsafe.Pointer(this)),
		uintptr(EvictionPriority),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIResource) GetEvictionPriority() (uint32, error) {
	var err error
	var evictionPriority uint32
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetEvictionPriority, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&evictionPriority)),
		0,
	)
	err = GetError(uint32(ret))
	if err != nil {
		return 0, err
	}
	return evictionPriority, nil
}

/*** IDXGIResource1 methods ***/
func (this *DXGIResource) CreateSubresourceSurface(index uint32) (*DXGISurface, error) {
	var err error
	var ppSurface *DXGISurface
	ret, _, _ := syscall.SyscallN(this.lpVtbl.CreateSubresourceSurface, uintptr(unsafe.Pointer(this)),
		uintptr(index),
		uintptr(unsafe.Pointer(&ppSurface)),
	)
	err = GetError(uint32(ret))
	return ppSurface, err
}
func (this *DXGIResource) CreateSharedHandle(pAttributes *SECURITY_ATTRIBUTES, dwAccess DWORD, lpName *[]byte) (*HANDLE, error) {
	var err error
	var pHandle *HANDLE
	ret, _, _ := syscall.SyscallN(this.lpVtbl.CreateSharedHandle, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pAttributes)),
		uintptr(dwAccess),
		uintptr(unsafe.Pointer(lpName)),
		uintptr(unsafe.Pointer(pHandle)),
		0,
	)
	err = GetError(uint32(ret))
	return pHandle, err
}
