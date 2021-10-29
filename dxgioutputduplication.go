package dxgi

import (
	"syscall"
	"unsafe"
)

var IID_IDXGIOutputDuplication = new(GUID)

func init() {
	IID_IDXGIOutputDuplication.Setup("191cfac3-a341-470d-b26e-a864f428319c")
}

type DXGIOutputDuplication struct {
	Unknown
	lpVtbl *outputDuplicationVtbl
}

type outputDuplicationVtbl struct {
	/*** IUnknown methods ***/
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
	/*** IDXGIObject methods ***/
	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	GetPrivateData          uintptr
	GetParent               uintptr
	/*** IDXGIOutputDuplication methods ***/
	GetDesc              uintptr
	AcquireNextFrame     uintptr
	GetFrameDirtyRects   uintptr
	GetFrameMoveRects    uintptr
	GetFramePointerShape uintptr
	MapDesktopSurface    uintptr
	UnMapDesktopSurface  uintptr
	ReleaseFrame         uintptr
}

/*** IDXGIObject methods ***/
func (this *DXGIOutputDuplication) QueryInterface(riid *GUID, ppvObject *interface{}) error {
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
func (this *DXGIOutputDuplication) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}
func (this *DXGIOutputDuplication) Release() uint32 {
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
func (this *DXGIOutputDuplication) SetPrivateData(Name *GUID, DataSize uint32, pData *interface{}) error {
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
func (this *DXGIOutputDuplication) SetPrivateDataInterface(Name *GUID, pUnknown *IUnknown) error {
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
func (this *DXGIOutputDuplication) GetPrivateData(Name *GUID, pDataSize *uint32) (*interface{}, error) {
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
func (this *DXGIOutputDuplication) GetParent(riid *GUID) (*interface{}, error) {
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

/*** IDXGIOutputDuplication methods ***/
func (this *DXGIOutputDuplication) GetDesc() (*DXGI_OUTDUPL_DESC, error) {
	var err error
	var pDesc *DXGI_OUTDUPL_DESC
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetDesc,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pDesc)),
		0,
	)
	err = GetError(uint32(ret))
	return pDesc, err
}
func (this *DXGIOutputDuplication) AcquireNextFrame(TimeoutInMilliseconds uint32) (*DXGI_OUTDUPL_FRAME_INFO, *DXGIResource, error) {
	var err error
	var pFrameInfo *DXGI_OUTDUPL_FRAME_INFO
	var ppDesktopResource *DXGIResource
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.AcquireNextFrame,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(TimeoutInMilliseconds),
		uintptr(unsafe.Pointer(pFrameInfo)),
		uintptr(unsafe.Pointer(&ppDesktopResource)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return pFrameInfo, ppDesktopResource, err
}
func (this *DXGIOutputDuplication) GetFrameDirtyRects(DirtyRectsBufferSize uint32) (*RECT, *uint32, error) {
	var err error
	var pDirtyRectsBuffer *RECT
	var pDirtyRectsBufferSizeRequired *uint32
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.GetFrameDirtyRects,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(DirtyRectsBufferSize),
		uintptr(unsafe.Pointer(pDirtyRectsBuffer)),
		uintptr(unsafe.Pointer(pDirtyRectsBufferSizeRequired)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return pDirtyRectsBuffer, pDirtyRectsBufferSizeRequired, err
}
func (this *DXGIOutputDuplication) GetFrameMoveRects(MoveRectsBufferSize uint32) (*RECT, *uint32, error) {
	var err error
	var pMoveRectBuffer *RECT
	var pMoveRectsBufferSizeRequired *uint32
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.GetFrameMoveRects,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(MoveRectsBufferSize),
		uintptr(unsafe.Pointer(pMoveRectBuffer)),
		uintptr(unsafe.Pointer(pMoveRectsBufferSizeRequired)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return pMoveRectBuffer, pMoveRectsBufferSizeRequired, err
}
func (this *DXGIOutputDuplication) GetFramePointerShape(PointerShapeBufferSize uint32) (*interface{}, *uint32, *DXGI_OUTDUPL_POINTER_SHAPE_INFO, error) {
	var err error
	var pPointerShapeBuffer *interface{}
	var pPointerShapeBufferSizeRequired *uint32
	var pPointerShapeInfo *DXGI_OUTDUPL_POINTER_SHAPE_INFO
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.GetFramePointerShape,
		5,
		uintptr(unsafe.Pointer(this)),
		uintptr(PointerShapeBufferSize),
		uintptr(unsafe.Pointer(pPointerShapeBuffer)),
		uintptr(unsafe.Pointer(pPointerShapeBufferSizeRequired)),
		uintptr(unsafe.Pointer(pPointerShapeInfo)),
		0,
	)
	err = GetError(uint32(ret))
	return pPointerShapeBuffer, pPointerShapeBufferSizeRequired, pPointerShapeInfo, err
}
func (this *DXGIOutputDuplication) MapDesktopSurface() (*DXGI_MAPPED_RECT, error) {
	var err error
	var pLockedRect *DXGI_MAPPED_RECT
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.MapDesktopSurface,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pLockedRect)),
		0,
	)
	err = GetError(uint32(ret))
	return pLockedRect, err
}
func (this *DXGIOutputDuplication) UnMapDesktopSurface() error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.UnMapDesktopSurface,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIOutputDuplication) ReleaseFrame() error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.ReleaseFrame,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return err
}
