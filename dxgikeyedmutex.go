package dxgi

import (
	"syscall"
	"unsafe"
)

var IID_IDXGIKeyedMutex = new(GUID)

func init() {
	IID_IDXGIKeyedMutex.Setup("9d8e1289-d7b3-465f-8126-250e349af85d")
}

type DXGIKeyedMutex struct {
	Unknown
	lpVtbl *keyedMutexVtbl
}

type keyedMutexVtbl struct {
	/*** IUnknown methods ***/
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
	/*** IDXGIObject methods ***/
	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	GetPrivateData          uintptr
	GetParent               uintptr
	/*** IDXGIKeyedMutex methods ***/
	AcquireSync uintptr
	ReleaseSync uintptr
}

/*** IDXGIObject methods ***/
func (this *DXGIKeyedMutex) QueryInterface(riid *GUID, ppvObject *interface{}) error {
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
func (this *DXGIKeyedMutex) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}
func (this *DXGIKeyedMutex) Release() uint32 {
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
func (this *DXGIKeyedMutex) SetPrivateData(Name *GUID, DataSize uint32, pData *interface{}) error {
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
func (this *DXGIKeyedMutex) SetPrivateDataInterface(Name *GUID, pUnknown *IUnknown) error {
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
func (this *DXGIKeyedMutex) GetPrivateData(Name *GUID, pDataSize *uint32) (*interface{}, error) {
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
func (this *DXGIKeyedMutex) GetParent(riid *GUID) (*interface{}, error) {
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

/*** IDXGIKeyedMutex methods ***/
func (this *DXGIKeyedMutex) AcquireSync(Key uint64, dwMilliseconds DWORD) error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.AcquireSync,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(Key),
		uintptr(dwMilliseconds),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIKeyedMutex) ReleaseSync(Key uint64) error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.ReleaseSync,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(Key),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
