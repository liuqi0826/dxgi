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
	ret, _, _ := syscall.SyscallN(this.lpVtbl.QueryInterface, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObject)),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIOutputDuplication) AddRef() uint32 {
	ret, _, _ := syscall.SyscallN(this.lpVtbl.AddRef, uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}
func (this *DXGIOutputDuplication) Release() uint32 {
	ret, _, _ := syscall.SyscallN(this.lpVtbl.Release, uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}

/*** IDXGIObject methods ***/
func (this *DXGIOutputDuplication) SetPrivateData(Name *GUID, DataSize uint32, pData *interface{}) error {
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
func (this *DXGIOutputDuplication) SetPrivateDataInterface(Name *GUID, pUnknown *IUnknown) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.SetPrivateDataInterface, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(Name)),
		uintptr(unsafe.Pointer(pUnknown)),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIOutputDuplication) GetPrivateData(Name *GUID, pDataSize *uint32) (*interface{}, error) {
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
func (this *DXGIOutputDuplication) GetParent(riid *GUID) (*interface{}, error) {
	var err error
	var ppParent *interface{}
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetParent, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppParent)),
	)
	err = GetError(uint32(ret))
	return ppParent, err
}

/*** IDXGIOutputDuplication methods ***/
func (this *DXGIOutputDuplication) GetDesc() (*DXGI_OUTDUPL_DESC, error) {
	var err error
	var desc DXGI_OUTDUPL_DESC
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
func (this *DXGIOutputDuplication) AcquireNextFrame(TimeoutInMilliseconds uint32) (*DXGI_OUTDUPL_FRAME_INFO, *DXGIResource, error) {
	var err error
	var frameInfo DXGI_OUTDUPL_FRAME_INFO
	var ppDesktopResource *DXGIResource
	ret, _, _ := syscall.SyscallN(this.lpVtbl.AcquireNextFrame, uintptr(unsafe.Pointer(this)),
		uintptr(TimeoutInMilliseconds),
		uintptr(unsafe.Pointer(&frameInfo)),
		uintptr(unsafe.Pointer(&ppDesktopResource)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	if err != nil {
		return nil, nil, err
	}
	return &frameInfo, ppDesktopResource, nil
}
func (this *DXGIOutputDuplication) GetFrameDirtyRects(DirtyRectsBufferSize uint32) (*RECT, *uint32, error) {
	var err error
	var pDirtyRectsBuffer *RECT
	var pDirtyRectsBufferSizeRequired *uint32
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetFrameDirtyRects, uintptr(unsafe.Pointer(this)),
		uintptr(DirtyRectsBufferSize),
		uintptr(unsafe.Pointer(pDirtyRectsBuffer)),
		uintptr(unsafe.Pointer(pDirtyRectsBufferSizeRequired)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return pDirtyRectsBuffer, pDirtyRectsBufferSizeRequired, err
}
func (this *DXGIOutputDuplication) GetFrameMoveRects(MoveRectsBufferSize uint32) (*DXGI_OUTDUPL_MOVE_RECT, *uint32, error) {
	var err error
	var pMoveRectBuffer *DXGI_OUTDUPL_MOVE_RECT
	var pMoveRectsBufferSizeRequired *uint32
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetFrameMoveRects, uintptr(unsafe.Pointer(this)),
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
	ret, _, _ := syscall.SyscallN(this.lpVtbl.GetFramePointerShape, uintptr(unsafe.Pointer(this)),
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
	var lockedRect DXGI_MAPPED_RECT
	ret, _, _ := syscall.SyscallN(this.lpVtbl.MapDesktopSurface, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&lockedRect)),
		0,
	)
	err = GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return &lockedRect, nil
}
func (this *DXGIOutputDuplication) UnMapDesktopSurface() error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.UnMapDesktopSurface, uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIOutputDuplication) ReleaseFrame() error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.ReleaseFrame, uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return err
}
