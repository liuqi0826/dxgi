package dxgi

import (
	"syscall"
	"unsafe"
)

var (
	IID_IDXGIOutput  = new(GUID)
	IID_IDXGIOutput1 = new(GUID)
	IID_IDXGIOutput2 = new(GUID)
	IID_IDXGIOutput3 = new(GUID)
	IID_IDXGIOutput4 = new(GUID)
	IID_IDXGIOutput5 = new(GUID)
	IID_IDXGIOutput6 = new(GUID)
)

func init() {
	IID_IDXGIOutput.Setup("ae02eedb-c735-4690-8d52-5a8dc20213aa")
	IID_IDXGIOutput1.Setup("00cddea8-939b-4b83-a340-a685226666cc")
	IID_IDXGIOutput2.Setup("595e39d1-2724-4663-99b1-da969de28364")
	IID_IDXGIOutput3.Setup("8a6bb301-7e7e-41F4-a8e0-5b32f7f99b18")
	IID_IDXGIOutput4.Setup("dc7dca35-2196-414d-9F53-617884032a60")
	IID_IDXGIOutput5.Setup("80A07424-AB52-42EB-833C-0C42FD282D98")
	IID_IDXGIOutput6.Setup("068346e8-aaec-4b84-add7-137f513f77a1")
}

type DXGIOutput struct {
	lpVtbl *outputVtbl
}
type outputVtbl struct {
	/*** IUnknown methods ***/
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
	/*** IDXGIObject methods ***/
	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	GetPrivateData          uintptr
	GetParent               uintptr
	/*** IDXGIOutput methods ***/
	GetDesc                     uintptr
	GetDisplayModeList          uintptr
	FindClosestMatchingMode     uintptr
	WaitForVBlank               uintptr
	TakeOwnership               uintptr
	ReleaseOwnership            uintptr
	GetGammaControlCapabilities uintptr
	SetGammaControl             uintptr
	GetGammaControl             uintptr
	SetDisplaySurface           uintptr
	GetDisplaySurfaceData       uintptr
	GetFrameStatistics          uintptr
	/*** IDXGIOutput1 methods ***/
	DuplicateOutput uintptr
	/*** IDXGIOutput2 methods ***/
	SupportsOverlays uintptr
	/*** IDXGIOutput3 methods ***/
	CheckOverlaySupport uintptr
	/*** IDXGIOutput4 methods ***/
	CheckOverlayColorSpaceSupport uintptr
	/*** IDXGIOutput5 methods ***/
	DuplicateOutput1 uintptr
	/*** IDXGIOutput6 methods ***/
	GetDesc1                        uintptr
	CheckHardwareCompositionSupport uintptr
}

/*** IDXGIOutput methods ***/
func (this *DXGIOutput) GetDesc() (*DXGIOutputDesc, error) {
	var outputDesc DXGIOutputDesc
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.GetDesc,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&outputDesc)),
	)
	err := GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return &outputDesc, nil
}
func (this *DXGIOutput) GetDisplayModeList(EnumFormat uint32, Flags uint32, NumModes *uint32, pDesc *DXGIModeDesc) error {
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.GetDisplayModeList,
		uintptr(unsafe.Pointer(this)),
		uintptr(EnumFormat),
		uintptr(Flags),
		uintptr(unsafe.Pointer(NumModes)),
		uintptr(unsafe.Pointer(pDesc)),
	)
	err := GetError(uint32(ret))
	return err
}
func (this *DXGIOutput) FindClosestMatchingMode(ModeToMatch *DXGIModeDesc, ClosestMatch *DXGIModeDesc, ConcernedDevice *IUnknown) error {
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.FindClosestMatchingMode,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(ModeToMatch)),
		uintptr(unsafe.Pointer(ClosestMatch)),
		uintptr(unsafe.Pointer(ConcernedDevice)),
	)
	err := GetError(uint32(ret))
	if err != nil {
		return err
	}
	return nil
}
func (this *DXGIOutput) WaitForVBlank() error {
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.WaitForVBlank,
		uintptr(unsafe.Pointer(this)),
	)
	err := GetError(uint32(ret))
	if err != nil {
		return err
	}
	return nil
}
func (this *DXGIOutput) TakeOwnership(Device *IUnknown, Exclusive bool) error {
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.TakeOwnership,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(Device)),
		boolToInt(Exclusive),
	)
	err := GetError(uint32(ret))
	if err != nil {
		return err
	}
	return nil
}
func (this *DXGIOutput) ReleaseOwnership() {
	syscall.SyscallN(
		this.lpVtbl.ReleaseOwnership,
		uintptr(unsafe.Pointer(this)),
	)
}
func (this *DXGIOutput) GetGammaControlCapabilities() (*DXGIGammaControlCapabilities, error) {
	var gammaCaps DXGIGammaControlCapabilities
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.GetGammaControlCapabilities,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&gammaCaps)),
	)
	err := GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return &gammaCaps, nil
}
func (this *DXGIOutput) SetGammaControl(array *DXGIGammaControl) error {
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.SetGammaControl,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(array)),
	)
	err := GetError(uint32(ret))
	if err != nil {
		return err
	}
	return nil
}
func (this *DXGIOutput) GetGammaControl() (*DXGIGammaControl, error) {
	var gammaControl DXGIGammaControl
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.GetGammaControl,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&gammaControl)),
	)
	err := GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return &gammaControl, nil
}
func (this *DXGIOutput) SetDisplaySurface(ScanoutSurface *DXGISurface) error {
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.SetDisplaySurface,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(ScanoutSurface)),
	)
	err := GetError(uint32(ret))
	if err != nil {
		return err
	}
	return nil
}
func (this *DXGIOutput) GetDisplaySurfaceData() (*DXGISurface, error) {
	var destination *DXGISurface
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.GetDisplaySurfaceData,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&destination)),
	)
	err := GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return destination, nil
}
func (this *DXGIOutput) GetFrameStatistics() (*DXGI_FRAME_STATISTICS, error) {
	var stats DXGI_FRAME_STATISTICS
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.GetFrameStatistics,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&stats)),
	)
	err := GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return &stats, nil
}

/*** IDXGIObject methods ***/
func (this *DXGIOutput) QueryInterface(riid *GUID, ppvObject *interface{}) error {
	var err error
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.QueryInterface,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObject)),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIOutput) AddRef() uint32 {
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.AddRef,
		uintptr(unsafe.Pointer(this)),
	)
	return uint32(ret)
}
func (this *DXGIOutput) Release() uint32 {
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.Release,
		uintptr(unsafe.Pointer(this)),
	)
	return uint32(ret)
}
func (this *DXGIOutput) SetPrivateData(Name *GUID, DataSize uint32, pData *interface{}) error {
	var err error
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.SetPrivateData,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(Name)),
		uintptr(DataSize),
		uintptr(unsafe.Pointer(pData)),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIOutput) SetPrivateDataInterface(Name *GUID, pUnknown *IUnknown) error {
	var err error
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.SetPrivateDataInterface,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(Name)),
		uintptr(unsafe.Pointer(pUnknown)),
	)
	err = GetError(uint32(ret))
	return err
}
func (this *DXGIOutput) GetPrivateData(Name *GUID, pDataSize *uint32) (*interface{}, error) {
	var err error
	var pData *interface{}
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.GetPrivateData,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(Name)),
		uintptr(unsafe.Pointer(pDataSize)),
		uintptr(unsafe.Pointer(pData)),
	)
	err = GetError(uint32(ret))
	return pData, err
}
func (this *DXGIOutput) GetParent(riid *GUID) (*interface{}, error) {
	var err error
	var ppParent *interface{}
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.GetParent,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppParent)),
	)
	err = GetError(uint32(ret))
	return ppParent, err
}

/*** IDXGIOutput1 methods ***/
func (this *DXGIOutput) DuplicateOutput(pDevice *IUnknown) (*DXGIOutputDuplication, error) {
	var err error
	var ppOutputDuplication *DXGIOutputDuplication
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.DuplicateOutput,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pDevice)),
		uintptr(unsafe.Pointer(&ppOutputDuplication)),
	)
	err = GetError(uint32(ret))
	return ppOutputDuplication, err
}

/*** IDXGIOutput2 methods ***/
func (this *DXGIOutput) SupportsOverlays() bool {
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.SupportsOverlays,
		uintptr(unsafe.Pointer(this)),
	)
	return intToBool(ret)
}

/*** IDXGIOutput3 methods ***/
func (this *DXGIOutput) CheckOverlaySupport(EnumFormat uint32, pConcernedDevice *interface{}) (uint32, error) {
	var err error
	var flags uint32
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.CheckOverlaySupport,
		uintptr(unsafe.Pointer(this)),
		uintptr(EnumFormat),
		uintptr(unsafe.Pointer(pConcernedDevice)),
		uintptr(unsafe.Pointer(&flags)),
	)
	err = GetError(uint32(ret))
	if err != nil {
		return 0, err
	}
	return flags, nil
}

/*** IDXGIOutput4 methods ***/
func (this *DXGIOutput) CheckOverlayColorSpaceSupport(Format uint32, ColorSpace uint32, pConcernedDevice *DXGIDevice) (uint32, error) {
	var err error
	var flags uint32
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.CheckOverlayColorSpaceSupport,
		uintptr(unsafe.Pointer(this)),
		uintptr(Format),
		uintptr(ColorSpace),
		uintptr(unsafe.Pointer(pConcernedDevice)),
		uintptr(unsafe.Pointer(&flags)),
	)
	err = GetError(uint32(ret))
	if err != nil {
		return 0, err
	}
	return flags, nil
}

/*** IDXGIOutput5 methods ***/
func (this *DXGIOutput) DuplicateOutput1(pDevice *DXGIDevice, SupportedFormatsCount uint32, pSupportedFormats *uint32) (*DXGIOutputDuplication, error) {
	var err error
	var ppOutputDuplication *DXGIOutputDuplication
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.DuplicateOutput1,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pDevice)),
		uintptr(SupportedFormatsCount),
		uintptr(unsafe.Pointer(pSupportedFormats)),
		uintptr(unsafe.Pointer(&ppOutputDuplication)),
	)
	err = GetError(uint32(ret))
	return ppOutputDuplication, err
}

/*** IDXGIOutput6 methods ***/
func (this *DXGIOutput) GetDesc1() (*DXGI_OUTPUT_DESC1, error) {
	var err error
	var desc DXGI_OUTPUT_DESC1
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.GetDesc1,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&desc)),
	)
	err = GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return &desc, nil
}
func (this *DXGIOutput) CheckHardwareCompositionSupport() (uint32, error) {
	var err error
	var flags uint32
	ret, _, _ := syscall.SyscallN(
		this.lpVtbl.CheckHardwareCompositionSupport,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&flags)),
	)
	err = GetError(uint32(ret))
	if err != nil {
		return 0, err
	}
	return flags, nil
}
