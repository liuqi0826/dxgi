package dxgi

import (
	"syscall"
	"unsafe"
)

var IID_IDXGIAdapter = new(GUID)
var IID_IDXGIAdapter1 = new(GUID)
var IID_IDXGIAdapter2 = new(GUID)
var IID_IDXGIAdapter3 = new(GUID)
var IID_IDXGIAdapter4 = new(GUID)

func init() {
	IID_IDXGIAdapter.Setup("2411e7e1-12ac-4ccf-bd14-9798e8534dc0")
	IID_IDXGIAdapter1.Setup("770aae78-f26f-4dba-a829-253c83d1b387")
	IID_IDXGIAdapter2.Setup("0AA1AE0A-FA0E-4B84-8644-E05FF8E5ACB5")
	IID_IDXGIAdapter3.Setup("645967A4-1392-4310-A798-8053CE3E93FD")
	IID_IDXGIAdapter4.Setup("3c8d99d1-4fbf-4181-a82c-af66bf7bd24e")
}

type DXGIAdapter struct {
	Unknown
	lpVtbl *adapterVtbl
}

type adapterVtbl struct {
	/*** IUnknown methods ***/
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
	/*** IDXGIObject methods ***/
	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	GetPrivateData          uintptr
	GetParent               uintptr
	/*** IDXGIAdapter methods ***/
	CheckInterfaceSupport uintptr
	EnumOutputs           uintptr
	GetDesc               uintptr
	/*** IDXGIAdapter1 methods ***/
	GetDesc1 uintptr
	/*** IDXGIAdapter2 methods ***/
	GetDesc2 uintptr
	/*** IDXGIAdapter3 methods ***/
	RegisterHardwareContentProtectionTeardownStatusEvent uintptr
	UnregisterHardwareContentProtectionTeardownStatus    uintptr
	QueryVideoMemoryInfo                                 uintptr
	SetVideoMemoryReservation                            uintptr
	RegisterVideoMemoryBudgetChangeNotificationEvent     uintptr
	UnregisterVideoMemoryBudgetChangeNotification        uintptr
	/*** IDXGIAdapter4 methods ***/
	GetDesc3 uintptr
}

/*** IDXGIObject methods ***/
func (this *DXGIAdapter) QueryInterface(riid *GUID, ppvObject *interface{}) error {
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
func (this *DXGIAdapter) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return uint32(ret)
}
func (this *DXGIAdapter) Release() uint32 {
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
func (this *DXGIAdapter) SetPrivateData(Name *GUID, DataSize uint32, pData *interface{}) error {
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
func (this *DXGIAdapter) SetPrivateDataInterface(Name *GUID, pUnknown *IUnknown) error {
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
func (this *DXGIAdapter) GetPrivateData(Name *GUID, pDataSize *uint32) (*interface{}, error) {
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
func (this *DXGIAdapter) GetParent(riid *GUID) (*interface{}, error) {
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

/*** IDXGIAdapter methods ***/
func (this *DXGIAdapter) CheckInterfaceSupport(riid *GUID, pUMDVersion *interface{}) error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.CheckInterfaceSupport,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(pUMDVersion)),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIAdapter) EnumOutputs(Output uint32) (*interface{}, error) {
	var err error
	var ppOutput *interface{}
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.EnumOutputs,
		1,
		uintptr(unsafe.Pointer(this)),
		uintptr(Output),
		uintptr(unsafe.Pointer(&ppOutput)),
	)
	err = GetError(uint32(ret))
	return ppOutput, err
}
func (this *DXGIAdapter) GetDesc() (*DXGI_ADAPTER_DESC, error) {
	var err error
	var pDesc *DXGI_ADAPTER_DESC
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

/*** IDXGIAdapter1 methods ***/
func (this *DXGIAdapter) GetDesc1() (*DXGI_ADAPTER_DESC1, error) {
	var err error
	var pDesc *DXGI_ADAPTER_DESC1
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetDesc1,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pDesc)),
		0,
	)
	err = GetError(uint32(ret))
	return pDesc, err
}

/*** IDXGIAdapter methods ***/

func (this *DXGIAdapter) GetDesc2() (*DXGI_ADAPTER_DESC2, error) {
	var err error
	var pDesc *DXGI_ADAPTER_DESC2
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetDesc2,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pDesc)),
		0,
	)
	err = GetError(uint32(ret))
	return pDesc, err
}

/*** IDXGIAdapter3 methods ***/
func (this *DXGIAdapter) RegisterHardwareContentProtectionTeardownStatusEvent(hEvent HANDLE) (*DWORD, error) {
	var err error
	var pdwCookie *DWORD
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.RegisterHardwareContentProtectionTeardownStatusEvent,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(hEvent)),
		uintptr(unsafe.Pointer(pdwCookie)),
	)
	err = GetError(uint32(ret))
	return pdwCookie, err
}
func (this *DXGIAdapter) UnregisterHardwareContentProtectionTeardownStatus(dwCookie DWORD) error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.UnregisterHardwareContentProtectionTeardownStatus,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(dwCookie),
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIAdapter) QueryVideoMemoryInfo(NodeIndex uint32, MemorySegmentGroup uint32) (*DXGI_QUERY_VIDEO_MEMORY_INFO, error) {
	var err error
	var pVideoMemoryInfo *DXGI_QUERY_VIDEO_MEMORY_INFO
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.QueryVideoMemoryInfo,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(NodeIndex),
		uintptr(MemorySegmentGroup),
		uintptr(unsafe.Pointer(pVideoMemoryInfo)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return pVideoMemoryInfo, err
}
func (this *DXGIAdapter) SetVideoMemoryReservation(NodeIndex uint32, MemorySegmentGroup uint32, Reservation uint64) error {
	var err error
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.SetVideoMemoryReservation,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(NodeIndex),
		uintptr(MemorySegmentGroup),
		uintptr(Reservation),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIAdapter) RegisterVideoMemoryBudgetChangeNotificationEvent(hEvent HANDLE) (*DWORD, error) {
	var err error
	var pdwCookie *DWORD
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.RegisterVideoMemoryBudgetChangeNotificationEvent,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(hEvent),
		uintptr(unsafe.Pointer(pdwCookie)),
	)
	err = GetError(uint32(ret))
	return pdwCookie, err
}
func (this *DXGIAdapter) UnregisterVideoMemoryBudgetChangeNotification(dwCookie DWORD) error {
	var err error
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.UnregisterVideoMemoryBudgetChangeNotification,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(dwCookie),
		0,
	)
	err = GetError(uint32(ret))
	return err
}

/*** IDXGIAdapter4 methods ***/
func (this *DXGIAdapter) GetDesc3() (*DXGI_ADAPTER_DESC3, error) {
	var err error
	var pDesc *DXGI_ADAPTER_DESC3
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetDesc3,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pDesc)),
		0,
	)
	err = GetError(uint32(ret))
	return pDesc, err
}
