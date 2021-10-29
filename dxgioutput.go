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
	Unknown
	lpVtbl *outputVtbl
}
type outputVtbl struct {
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
	var outputDesc *DXGIOutputDesc
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetDesc,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(outputDesc)),
		0,
	)
	err := GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return outputDesc, nil
}
func (this *DXGIOutput) GetDisplayModeList(EnumFormat uint32, Flags uint32, NumModes *uint32) (*DXGIModeDesc, error) {
	var modeDesc *DXGIModeDesc
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.GetDisplayModeList,
		5,
		uintptr(unsafe.Pointer(this)),
		uintptr(EnumFormat),
		uintptr(Flags),
		uintptr(unsafe.Pointer(NumModes)),
		uintptr(unsafe.Pointer(modeDesc)),
		0,
	)
	err := GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return modeDesc, nil
}
func (this *DXGIOutput) FindClosestMatchingMode(ModeToMatch *DXGIModeDesc, ClosestMatch *DXGIModeDesc, ConcernedDevice *IUnknown) error {
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.FindClosestMatchingMode,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(ModeToMatch)),
		uintptr(unsafe.Pointer(ClosestMatch)),
		uintptr(unsafe.Pointer(ConcernedDevice)),
		0,
		0,
	)
	err := GetError(uint32(ret))
	if err != nil {
		return err
	}
	return nil
}
func (this *DXGIOutput) WaitForVBlank() error {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.WaitForVBlank,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	err := GetError(uint32(ret))
	if err != nil {
		return err
	}
	return nil
}
func (this *DXGIOutput) TakeOwnership(Device *IUnknown, Exclusive bool) error {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.TakeOwnership,
		3,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(Device)),
		uintptr(unsafe.Pointer(&Exclusive)),
	)
	err := GetError(uint32(ret))
	if err != nil {
		return err
	}
	return nil
}
func (this *DXGIOutput) ReleaseOwnership() {
	syscall.Syscall(
		this.lpVtbl.ReleaseOwnership,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
}
func (this *DXGIOutput) GetGammaControlCapabilities() (*DXGIGammaControlCapabilities, error) {
	var gammaCaps *DXGIGammaControlCapabilities
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetGammaControlCapabilities,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(gammaCaps)),
		0,
	)
	err := GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return gammaCaps, nil
}
func (this *DXGIOutput) SetGammaControl(array *DXGIGammaControl) error {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.SetGammaControl,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(array)),
		0,
	)
	err := GetError(uint32(ret))
	if err != nil {
		return err
	}
	return nil
}
func (this *DXGIOutput) GetGammaControl() (*DXGIGammaControl, error) {
	var array *DXGIGammaControl
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetGammaControl,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(array)),
		0,
	)
	err := GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return array, nil
}
func (this *DXGIOutput) SetDisplaySurface(ScanoutSurface *DXGISurface) error {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.SetDisplaySurface,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(ScanoutSurface)),
		0,
	)
	err := GetError(uint32(ret))
	if err != nil {
		return err
	}
	return nil
}
func (this *DXGIOutput) GetDisplaySurfaceData() (*DXGISurface, error) {
	var destination *DXGISurface
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetDisplaySurfaceData,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(destination)),
		0,
	)
	err := GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return destination, nil
}
func (this *DXGIOutput) GetFrameStatistics() (*DXGI_FRAME_STATISTICS, error) {
	var stats *DXGI_FRAME_STATISTICS
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.GetFrameStatistics,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(stats)),
		0,
	)
	err := GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return stats, nil
}

/*** IDXGIOutput2 methods ***/
func (this *DXGIOutput) SupportsOverlays() bool {
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.SupportsOverlays,
		1,
		uintptr(unsafe.Pointer(this)),
		0,
		0,
	)
	return intToBool(ret)
}

/*** IDXGIOutput3 methods ***/
func (this *DXGIOutput) CheckOverlaySupport(EnumFormat uint32, pConcernedDevice *interface{}) (*uint32, error) {
	var err error
	var pFlags *uint32
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.CheckOverlaySupport,
		4,
		uintptr(unsafe.Pointer(this)),
		uintptr(EnumFormat),
		uintptr(unsafe.Pointer(pConcernedDevice)),
		uintptr(unsafe.Pointer(pFlags)),
		0,
		0,
	)
	err = GetError(uint32(ret))
	return pFlags, err
}

/*** IDXGIOutput4 methods ***/
func (this *DXGIOutput) CheckOverlayColorSpaceSupport(Format uint32, ColorSpace uint32, pConcernedDevice *DXGIDevice) (*uint32, error) {
	var err error
	var pFlags *uint32
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.CheckOverlayColorSpaceSupport,
		5,
		uintptr(unsafe.Pointer(this)),
		uintptr(Format),
		uintptr(ColorSpace),
		uintptr(unsafe.Pointer(pConcernedDevice)),
		uintptr(unsafe.Pointer(pFlags)),
		0,
	)
	err = GetError(uint32(ret))
	return pFlags, err
}

/*** IDXGIOutput5 methods ***/
func (this *DXGIOutput) DuplicateOutput1(pDevice *DXGIDevice, SupportedFormatsCount uint32, pSupportedFormats *uint32) (*DXGIOutputDuplication, error) {
	var err error
	var ppOutputDuplication *DXGIOutputDuplication
	ret, _, _ := syscall.Syscall6(
		this.lpVtbl.DuplicateOutput1,
		5,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pDevice)),
		uintptr(SupportedFormatsCount),
		uintptr(unsafe.Pointer(pSupportedFormats)),
		uintptr(unsafe.Pointer(&ppOutputDuplication)),
		0,
	)
	err = GetError(uint32(ret))
	return ppOutputDuplication, err
}

/*** IDXGIOutput6 methods ***/
func (this *DXGIOutput) GetDesc1() (*DXGI_OUTPUT_DESC1, error) {
	var err error
	var pDesc *DXGI_OUTPUT_DESC1
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
func (this *DXGIOutput) CheckHardwareCompositionSupport() (*uint32, error) {
	var err error
	var pFlags *uint32
	ret, _, _ := syscall.Syscall(
		this.lpVtbl.CheckHardwareCompositionSupport,
		2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pFlags)),
		0,
	)
	err = GetError(uint32(ret))
	return pFlags, err
}
