package dxgi

import (
	"syscall"
	"unsafe"
)

var (
	IID_IDXGISurface  = new(GUID)
	IID_IDXGISurface1 = new(GUID)
	IID_IDXGISurface2 = new(GUID)
)

func init() {
	IID_IDXGISurface.Setup("cafcb56c-6ac3-4889-bf47-9e23bbd260ec")
	IID_IDXGISurface1.Setup("4AE63092-6327-4c1b-80AE-BFE12EA32B86")
	IID_IDXGISurface2.Setup("ba496dd-b617-4cb8-a866-bc44d7eb1fa2")
}

type DXGISurface struct {
	Unknown
	lpVtbl *surfaceVtbl
}
type surfaceVtbl struct {
	/*** IUnknown methods ***/
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
	/*** IDXGIObject methods ***/
	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	GetPrivateData          uintptr
	GetParent               uintptr
	/*** IDXGISurface methods ***/
	GetDesc uintptr
	Map     uintptr
	Unmap   uintptr
	/*** IDXGISurface1 methods ***/
	GetDC     uintptr
	ReleaseDC uintptr
	/*** IDXGISurface2 methods ***/
	GetResource uintptr
}

/*** IDXGIObject methods ***/
func (this *DXGISurface) QueryInterface(riid *GUID, ppvObject *interface{}) error {
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
func (this *DXGISurface) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}
func (this *DXGISurface) Release() uint32 {
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
func (this *DXGISurface) SetPrivateData(Name *GUID, DataSize uint32, pData *interface{}) error {
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
func (this *DXGISurface) SetPrivateDataInterface(Name *GUID, pUnknown *IUnknown) error {
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
func (this *DXGISurface) GetPrivateData(Name *GUID, pDataSize *uint32) (*interface{}, error) {
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
func (this *DXGISurface) GetParent(riid *GUID) (*interface{}, error) {
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

/*** IDXGISurface methods ***/
func (this *DXGISurface) GetDesc() (*DXGISurfaceDesc, error) {
	var desc *DXGISurfaceDesc
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetDesc,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(desc)),
		0,
	)
	err := GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return desc, nil
}
func (this *DXGISurface) Map(lockedRect *DXGISurfaceDesc, MapFlags uint32) error {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.Map,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(lockedRect)),
		uintptr(MapFlags),
	)
	err := GetError(uint32(ret))
	return err
}
func (this *DXGISurface) Unmap() error {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.Unmap,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	err := GetError(uint32(ret))
	return err
}

/*** IDXGISurface1 methods ***/
func (this *DXGISurface) GetDC(Discard bool) (*HDC, error) {
	var phdc *HDC
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetDC,
		3,
		uintptr(unsafe.Pointer(this)),
		boolToInt(Discard),
		uintptr(unsafe.Pointer(phdc)),
	)
	err := GetError(uint32(ret))
	return phdc, err
}
func (this *DXGISurface) ReleaseDC(pDirtyRect *RECT) error {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.ReleaseDC,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pDirtyRect)),
		0,
	)
	err := GetError(uint32(ret))
	return err
}

/*** IDXGISurface2 methods ***/
func (this *DXGISurface) GetResource(riid *GUID) (*interface{}, *uint32, error) {
	var ppParentResource *interface{}
	var pSubresourceIndex *uint32
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.GetResource,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppParentResource)),
		uintptr(unsafe.Pointer(pSubresourceIndex)),
		0,
		0,
	)
	err := GetError(uint32(ret))
	return ppParentResource, pSubresourceIndex, err
}
