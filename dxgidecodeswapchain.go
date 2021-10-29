package dxgi

import (
	"syscall"
	"unsafe"
)

var IID_IDXGIDecodeSwapChain = new(GUID)

func init() {
	IID_IDXGIDecodeSwapChain.Setup("2633066b-4514-4c7a-8fd8-12ea98059d18")
}

type DXGIDecodeSwapChain struct {
	Unknown
	lpVtbl *decodeSwapChainVtbl
}

type decodeSwapChainVtbl struct {
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
	/*** IDXGIDecodeSwapChain methods ***/
	PresentBuffer uintptr
	SetSourceRect uintptr
	SetTargetRect uintptr
	SetDestSize   uintptr
	GetSourceRect uintptr
	GetTargetRect uintptr
	GetDestSize   uintptr
	SetColorSpace uintptr
	GetColorSpace uintptr
}

/*** IDXGIObject methods ***/
func (this *DXGIDecodeSwapChain) QueryInterface(riid *GUID, ppvObject *interface{}) error {
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
func (this *DXGIDecodeSwapChain) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}
func (this *DXGIDecodeSwapChain) Release() uint32 {
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
func (this *DXGIDecodeSwapChain) SetPrivateData(Name *GUID, DataSize uint32, pData *interface{}) error {
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
func (this *DXGIDecodeSwapChain) SetPrivateDataInterface(Name *GUID, pUnknown *IUnknown) error {
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
func (this *DXGIDecodeSwapChain) GetPrivateData(Name *GUID, pDataSize *uint32) (*interface{}, error) {
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
func (this *DXGIDecodeSwapChain) GetParent(riid *GUID) (*interface{}, error) {
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

/*** IDXGIDecodeSwapChain methods ***/
func (this *DXGIDecodeSwapChain) PresentBuffer(BufferToPresent, SyncInterval, Flags uint32) error {
	var err error
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.PresentBuffer,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(BufferToPresent),
		uintptr(SyncInterval),
		uintptr(Flags),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIDecodeSwapChain) SetSourceRect(pRect *RECT) error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.SetSourceRect,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pRect)),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIDecodeSwapChain) SetTargetRect(pRect *RECT) error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.SetTargetRect,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pRect)),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIDecodeSwapChain) SetDestSize(Width, Height uint32) error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.SetDestSize,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(Width),
		uintptr(Height),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIDecodeSwapChain) GetSourceRect() (*RECT, error) {
	var err error
	var pRect *RECT
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetSourceRect,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pRect)),
		0,
	)
	err = GetError(uint32(ret))
	return pRect, err
}
func (this *DXGIDecodeSwapChain) GetTargetRect() (*RECT, error) {
	var err error
	var pRect *RECT
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetTargetRect,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pRect)),
		0,
	)
	err = GetError(uint32(ret))
	return pRect, err
}
func (this *DXGIDecodeSwapChain) GetDestSize() (uint32, uint32, error) {
	var err error
	var Width, Height uint32
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetDestSize,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(Width),
		uintptr(Height),
	)
	err = GetError(uint32(ret))
	return Width, Height, err
}
func (this *DXGIDecodeSwapChain) SetColorSpace(ColorSpace uint32) error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.SetColorSpace,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(ColorSpace),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIDecodeSwapChain) GetColorSpace() uint32 {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetColorSpace,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}
