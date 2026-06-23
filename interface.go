package dxgi

import (
	"syscall"
	"unsafe"
)

type IUnknown interface {
	QueryInterface(*GUID, *interface{}) error
	AddRef() uint32
	Release() uint32
}

type Unknown struct {
	lpVtbl *unknownVtbl
}
type unknownVtbl struct {
	/*** IUnknown methods ***/
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

/*** IUnknown methods ***/
func (this *Unknown) QueryInterface(riid *GUID, ppvObject *interface{}) error {
	var err error
	ret, _, _ := syscall.SyscallN(this.lpVtbl.QueryInterface, uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObject)),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *Unknown) AddRef() uint32 {
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.AddRef,
		uintptr(unsafe.Pointer(this)),
	)
	return uint32(ret)
}
func (this *Unknown) Release() uint32 {
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.Release,
		uintptr(unsafe.Pointer(this)),
	)
	return uint32(ret)
}
