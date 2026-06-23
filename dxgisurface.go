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
	ret, _, _ := syscall.SyscallN(this.lpVtbl.QueryInterface, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObject)),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGISurface) AddRef() uint32 {
	ret, _, _ := syscall.SyscallN(this.lpVtbl.AddRef, uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}
func (this *DXGISurface) Release() uint32 {
	ret, _, _ := syscall.SyscallN(this.lpVtbl.Release, uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}

/*** IDXGIObject methods ***/
func (this *DXGISurface) SetPrivateData(Name *GUID, DataSize uint32, pData *interface{}) error {
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
func (this *DXGISurface) SetPrivateDataInterface(Name *GUID, pUnknown *IUnknown) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.SetPrivateDataInterface, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(Name)),
		uintptr(unsafe.Pointer(pUnknown)),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGISurface) GetPrivateData(Name *GUID, pDataSize *uint32) (*interface{}, error) {
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
func (this *DXGISurface) GetParent(riid *GUID) (*interface{}, error) {
	var err error
	var ppParent *interface{}
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetParent, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppParent)),
	)
	err = GetError(uint32(ret))
	return ppParent, err
}

/*** IDXGISurface methods ***/
func (this *DXGISurface) GetDesc() (*DXGISurfaceDesc, error) {
	var desc DXGISurfaceDesc
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetDesc, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&desc)),
		0,
	)
	err := GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return &desc, nil
}
func (this *DXGISurface) Map(MapFlags uint32) (*DXGI_MAPPED_RECT, error) {
	var lockedRect DXGI_MAPPED_RECT
	ret, _, _ := syscall.SyscallN(this.lpVtbl.Map, uintptr(unsafe.Pointer(this)),
		uintptr(MapFlags),
		uintptr(unsafe.Pointer(&lockedRect)),
	)
	err := GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return &lockedRect, nil
}
func (this *DXGISurface) Unmap() error {
	ret, _, _ := syscall.SyscallN(this.lpVtbl.Unmap, uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	err := GetError(uint32(ret))
	return err
}

/*** IDXGISurface1 methods ***/
func (this *DXGISurface) GetDC(Discard bool) (HDC, error) {
	var phdc HDC
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetDC, uintptr(unsafe.Pointer(this)),
		boolToInt(Discard),
		uintptr(unsafe.Pointer(&phdc)),
	)
	err := GetError(uint32(ret))
	if err != nil {
		return 0, err
	}
	return phdc, nil
}
func (this *DXGISurface) ReleaseDC(pDirtyRect *RECT) error {
	ret, _, _ := syscall.SyscallN(this.lpVtbl.ReleaseDC, uintptr(unsafe.Pointer(this)),
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
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetResource, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppParentResource)),
		uintptr(unsafe.Pointer(pSubresourceIndex)),
		0,
		0,
	)
	err := GetError(uint32(ret))
	return ppParentResource, pSubresourceIndex, err
}
