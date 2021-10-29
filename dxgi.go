package dxgi

import "C"
import (
	"syscall"
	"unsafe"
)

//++++++++++++++++++++ api ++++++++++++++++++++
var (
	dxgi = syscall.NewLazyDLL("dxgi.dll")

	createDXGIFactory  = dxgi.NewProc("CreateDXGIFactory")
	createDXGIFactory1 = dxgi.NewProc("CreateDXGIFactory1")
	createDXGIFactory2 = dxgi.NewProc("CreateDXGIFactory2")

	getDebugInterface            = dxgi.NewProc("DXGIGetDebugInterface")
	getDebugInterface1           = dxgi.NewProc("DXGIGetDebugInterface1")
	declareAdapterRemovalSupport = dxgi.NewProc("DXGIDeclareAdapterRemovalSupport")
)

func CreateDXGIFactory(riid *GUID) (*DXGIFactory, error) {
	var factory uintptr
	ret, _, _ := createDXGIFactory.Call(
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&factory)),
	)
	err := GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return (*DXGIFactory)(unsafe.Pointer(factory)), nil
}
func CreateDXGIFactory1(riid *GUID) (*DXGIFactory, error) {
	var factory *DXGIFactory
	ret, _, _ := createDXGIFactory1.Call(
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&factory)),
	)
	err := GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return factory, nil
}
func CreateDXGIFactory2(flags uint32, riid *GUID) (*DXGIFactory, error) {
	var factory *DXGIFactory
	ret, _, _ := createDXGIFactory2.Call(
		uintptr(flags),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&factory)),
	)
	err := GetError(uint32(ret))
	if err != nil {
		return nil, err
	}
	return factory, nil
}

//++++++++++++++++++++ typedef ++++++++++++++++++++
type (
	DXGI_USAGE = uint32
)

//++++++++++++++++++++ enum ++++++++++++++++++++
//++++++++++++++++++++ dxgi.h ++++++++++++++++++++
type dxgiAdapterFlag struct {
	DXGI_ADAPTER_FLAG_NONE        uint32
	DXGI_ADAPTER_FLAG_REMOTE      uint32
	DXGI_ADAPTER_FLAG_SOFTWARE    uint32
	DXGI_ADAPTER_FLAG_FORCE_DWORD uint32
}
type dxgiResidency struct {
	DXGI_RESIDENCY_FULLY_RESIDENT            uint32
	DXGI_RESIDENCY_RESIDENT_IN_SHARED_MEMORY uint32
	DXGI_RESIDENCY_EVICTED_TO_DISK           uint32
}
type dxgiSwapEffect struct {
	DXGI_SWAP_EFFECT_DISCARD         uint32
	DXGI_SWAP_EFFECT_SEQUENTIAL      uint32
	DXGI_SWAP_EFFECT_FLIP_SEQUENTIAL uint32
	DXGI_SWAP_EFFECT_FLIP_DISCARD    uint32
}
type dxgiSwapChainFlag struct {
	DXGI_SWAP_CHAIN_FLAG_NONPREROTATED                          uint32
	DXGI_SWAP_CHAIN_FLAG_ALLOW_MODE_SWITCH                      uint32
	DXGI_SWAP_CHAIN_FLAG_GDI_COMPATIBLE                         uint32
	DXGI_SWAP_CHAIN_FLAG_RESTRICTED_CONTENT                     uint32
	DXGI_SWAP_CHAIN_FLAG_RESTRICT_SHARED_RESOURCE_DRIVER        uint32
	DXGI_SWAP_CHAIN_FLAG_DISPLAY_ONLY                           uint32
	DXGI_SWAP_CHAIN_FLAG_FRAME_LATENCY_WAITABLE_OBJECT          uint32
	DXGI_SWAP_CHAIN_FLAG_FOREGROUND_LAYER                       uint32
	DXGI_SWAP_CHAIN_FLAG_FULLSCREEN_VIDEO                       uint32
	DXGI_SWAP_CHAIN_FLAG_YUV_VIDEO                              uint32
	DXGI_SWAP_CHAIN_FLAG_HW_PROTECTED                           uint32
	DXGI_SWAP_CHAIN_FLAG_ALLOW_TEARING                          uint32
	DXGI_SWAP_CHAIN_FLAG_RESTRICTED_TO_ALL_HOLOGRAPHIC_DISPLAYS uint32
}

var DXGI_ADAPTER_FLAG = dxgiAdapterFlag{
	DXGI_ADAPTER_FLAG_NONE:        0,
	DXGI_ADAPTER_FLAG_REMOTE:      1,
	DXGI_ADAPTER_FLAG_SOFTWARE:    2,
	DXGI_ADAPTER_FLAG_FORCE_DWORD: 0xffffffff,
}
var DXGI_RESIDENCY = dxgiResidency{
	DXGI_RESIDENCY_FULLY_RESIDENT:            1,
	DXGI_RESIDENCY_RESIDENT_IN_SHARED_MEMORY: 2,
	DXGI_RESIDENCY_EVICTED_TO_DISK:           3,
}
var DXGI_SWAP_EFFECT = dxgiSwapEffect{
	DXGI_SWAP_EFFECT_DISCARD:         1,
	DXGI_SWAP_EFFECT_SEQUENTIAL:      2,
	DXGI_SWAP_EFFECT_FLIP_SEQUENTIAL: 3,
	DXGI_SWAP_EFFECT_FLIP_DISCARD:    4,
}
var DXGI_SWAP_CHAIN_FLAG = dxgiSwapChainFlag{
	DXGI_SWAP_CHAIN_FLAG_NONPREROTATED:                          1,
	DXGI_SWAP_CHAIN_FLAG_ALLOW_MODE_SWITCH:                      2,
	DXGI_SWAP_CHAIN_FLAG_GDI_COMPATIBLE:                         4,
	DXGI_SWAP_CHAIN_FLAG_RESTRICTED_CONTENT:                     8,
	DXGI_SWAP_CHAIN_FLAG_RESTRICT_SHARED_RESOURCE_DRIVER:        16,
	DXGI_SWAP_CHAIN_FLAG_DISPLAY_ONLY:                           32,
	DXGI_SWAP_CHAIN_FLAG_FRAME_LATENCY_WAITABLE_OBJECT:          64,
	DXGI_SWAP_CHAIN_FLAG_FOREGROUND_LAYER:                       128,
	DXGI_SWAP_CHAIN_FLAG_FULLSCREEN_VIDEO:                       256,
	DXGI_SWAP_CHAIN_FLAG_YUV_VIDEO:                              512,
	DXGI_SWAP_CHAIN_FLAG_HW_PROTECTED:                           1024,
	DXGI_SWAP_CHAIN_FLAG_ALLOW_TEARING:                          2048,
	DXGI_SWAP_CHAIN_FLAG_RESTRICTED_TO_ALL_HOLOGRAPHIC_DISPLAYS: 4096,
}

//++++++++++++++++++++ dxgi1_2.h ++++++++++++++++++++
type dxgiAlphaMode struct {
	DXGI_ALPHA_MODE_UNSPECIFIED   uint32
	DXGI_ALPHA_MODE_PREMULTIPLIED uint32
	DXGI_ALPHA_MODE_STRAIGHT      uint32
	DXGI_ALPHA_MODE_IGNORE        uint32
	DXGI_ALPHA_MODE_FORCE_DWORD   uint32
}
type dxgiComputePreemptionGranularity struct {
	DXGI_COMPUTE_PREEMPTION_DMA_BUFFER_BOUNDARY   uint32
	DXGI_COMPUTE_PREEMPTION_DISPATCH_BOUNDARY     uint32
	DXGI_COMPUTE_PREEMPTION_THREAD_GROUP_BOUNDARY uint32
	DXGI_COMPUTE_PREEMPTION_THREAD_BOUNDARY       uint32
	DXGI_COMPUTE_PREEMPTION_INSTRUCTION_BOUNDARY  uint32
}
type dxgiGraphicsPreemptionGranularity struct {
	DXGI_GRAPHICS_PREEMPTION_DMA_BUFFER_BOUNDARY  uint32
	DXGI_GRAPHICS_PREEMPTION_PRIMITIVE_BOUNDARY   uint32
	DXGI_GRAPHICS_PREEMPTION_TRIANGLE_BOUNDARY    uint32
	DXGI_GRAPHICS_PREEMPTION_PIXEL_BOUNDARY       uint32
	DXGI_GRAPHICS_PREEMPTION_INSTRUCTION_BOUNDARY uint32
}
type dxgiOfferResourcePriority struct {
	DXGI_OFFER_RESOURCE_PRIORITY_LOW    uint32
	DXGI_OFFER_RESOURCE_PRIORITY_NORMAL uint32
	DXGI_OFFER_RESOURCE_PRIORITY_HIGH   uint32
}
type dxgiOutduplPointerShapeType struct {
	DXGI_OUTDUPL_POINTER_SHAPE_TYPE_MONOCHROME   uint32
	DXGI_OUTDUPL_POINTER_SHAPE_TYPE_COLOR        uint32
	DXGI_OUTDUPL_POINTER_SHAPE_TYPE_MASKED_COLOR uint32
}
type dxgiScaling struct {
	DXGI_SCALING_STRETCH              uint32
	DXGI_SCALING_NONE                 uint32
	DXGI_SCALING_ASPECT_RATIO_STRETCH uint32
}

var DXGI_ALPHA_MODE = dxgiAlphaMode{
	DXGI_ALPHA_MODE_UNSPECIFIED:   0,
	DXGI_ALPHA_MODE_PREMULTIPLIED: 1,
	DXGI_ALPHA_MODE_STRAIGHT:      2,
	DXGI_ALPHA_MODE_IGNORE:        3,
	DXGI_ALPHA_MODE_FORCE_DWORD:   0xffffffff,
}
var DXGI_COMPUTE_PREEMPTION_GRANULARITY = dxgiComputePreemptionGranularity{
	DXGI_COMPUTE_PREEMPTION_DMA_BUFFER_BOUNDARY:   0,
	DXGI_COMPUTE_PREEMPTION_DISPATCH_BOUNDARY:     1,
	DXGI_COMPUTE_PREEMPTION_THREAD_GROUP_BOUNDARY: 2,
	DXGI_COMPUTE_PREEMPTION_THREAD_BOUNDARY:       3,
	DXGI_COMPUTE_PREEMPTION_INSTRUCTION_BOUNDARY:  4,
}
var DXGI_GRAPHICS_PREEMPTION_GRANULARITY = dxgiGraphicsPreemptionGranularity{
	DXGI_GRAPHICS_PREEMPTION_DMA_BUFFER_BOUNDARY:  0,
	DXGI_GRAPHICS_PREEMPTION_PRIMITIVE_BOUNDARY:   1,
	DXGI_GRAPHICS_PREEMPTION_TRIANGLE_BOUNDARY:    2,
	DXGI_GRAPHICS_PREEMPTION_PIXEL_BOUNDARY:       3,
	DXGI_GRAPHICS_PREEMPTION_INSTRUCTION_BOUNDARY: 4,
}
var DXGI_OFFER_RESOURCE_PRIORITY = dxgiOfferResourcePriority{
	DXGI_OFFER_RESOURCE_PRIORITY_LOW:    1,
	DXGI_OFFER_RESOURCE_PRIORITY_NORMAL: 2,
	DXGI_OFFER_RESOURCE_PRIORITY_HIGH:   3,
}
var DXGI_OUTDUPL_POINTER_SHAPE_TYPE = dxgiOutduplPointerShapeType{
	DXGI_OUTDUPL_POINTER_SHAPE_TYPE_MONOCHROME:   0x1,
	DXGI_OUTDUPL_POINTER_SHAPE_TYPE_COLOR:        0x2,
	DXGI_OUTDUPL_POINTER_SHAPE_TYPE_MASKED_COLOR: 0x4,
}
var DXGI_SCALING = dxgiScaling{
	DXGI_SCALING_STRETCH:              0,
	DXGI_SCALING_NONE:                 1,
	DXGI_SCALING_ASPECT_RATIO_STRETCH: 2,
}

//++++++++++++++++++++ dxgi1_3.h ++++++++++++++++++++
type dxgiFramePresentationMode struct {
	DXGI_FRAME_PRESENTATION_MODE_COMPOSED            uint32
	DXGI_FRAME_PRESENTATION_MODE_OVERLAY             uint32
	DXGI_FRAME_PRESENTATION_MODE_NONE                uint32
	DXGI_FRAME_PRESENTATION_MODE_COMPOSITION_FAILURE uint32
}
type dxgiMultiplaneOverlayYcbcrFlags struct {
	DXGI_MULTIPLANE_OVERLAY_YCbCr_FLAG_NOMINAL_RANGE uint32
	DXGI_MULTIPLANE_OVERLAY_YCbCr_FLAG_BT709         uint32
	DXGI_MULTIPLANE_OVERLAY_YCbCr_FLAG_xvYCC         uint32
}
type dxgiOverlaySupportFlag struct {
	DXGI_OVERLAY_SUPPORT_FLAG_DIRECT  uint32
	DXGI_OVERLAY_SUPPORT_FLAG_SCALING uint32
}

var DXGI_FRAME_PRESENTATION_MODE = dxgiFramePresentationMode{
	DXGI_FRAME_PRESENTATION_MODE_COMPOSED:            0,
	DXGI_FRAME_PRESENTATION_MODE_OVERLAY:             1,
	DXGI_FRAME_PRESENTATION_MODE_NONE:                2,
	DXGI_FRAME_PRESENTATION_MODE_COMPOSITION_FAILURE: 3,
}
var DXGI_MULTIPLANE_OVERLAY_YCbCr_FLAGS = dxgiMultiplaneOverlayYcbcrFlags{
	DXGI_MULTIPLANE_OVERLAY_YCbCr_FLAG_NOMINAL_RANGE: 0x1,
	DXGI_MULTIPLANE_OVERLAY_YCbCr_FLAG_BT709:         0x2,
	DXGI_MULTIPLANE_OVERLAY_YCbCr_FLAG_xvYCC:         0x4,
}
var DXGI_OVERLAY_SUPPORT_FLAG = dxgiOverlaySupportFlag{
	DXGI_OVERLAY_SUPPORT_FLAG_DIRECT:  0x1,
	DXGI_OVERLAY_SUPPORT_FLAG_SCALING: 0x2,
}

//++++++++++++++++++++ dxgi1_4.h ++++++++++++++++++++
type dxgiMemorySegmentGroup struct {
	DXGI_MEMORY_SEGMENT_GROUP_LOCAL     uint32
	DXGI_MEMORY_SEGMENT_GROUP_NON_LOCAL uint32
}
type dxgiOverlayColorSpaceSupportFlag struct {
	DXGI_OVERLAY_COLOR_SPACE_SUPPORT_FLAG_PRESENT uint32
}
type dxgiSwapChainColorSpaceSupportFlag struct {
	DXGI_SWAP_CHAIN_COLOR_SPACE_SUPPORT_FLAG_PRESENT         uint32
	DXGI_SWAP_CHAIN_COLOR_SPACE_SUPPORT_FLAG_OVERLAY_PRESENT uint32
}

var DXGI_MEMORY_SEGMENT_GROUP_LOCAL = dxgiMemorySegmentGroup{
	DXGI_MEMORY_SEGMENT_GROUP_LOCAL:     0,
	DXGI_MEMORY_SEGMENT_GROUP_NON_LOCAL: 1,
}
var DXGI_OVERLAY_COLOR_SPACE_SUPPORT_FLAG = dxgiOverlayColorSpaceSupportFlag{
	DXGI_OVERLAY_COLOR_SPACE_SUPPORT_FLAG_PRESENT: 0x1,
}
var DXGI_SWAP_CHAIN_COLOR_SPACE_SUPPORT_FLAG = dxgiSwapChainColorSpaceSupportFlag{
	DXGI_SWAP_CHAIN_COLOR_SPACE_SUPPORT_FLAG_PRESENT:         0x1,
	DXGI_SWAP_CHAIN_COLOR_SPACE_SUPPORT_FLAG_OVERLAY_PRESENT: 0x2,
}

//++++++++++++++++++++ dxgi1_5.h ++++++++++++++++++++
type dxgiFeature struct {
	DXGI_FEATURE_PRESENT_ALLOW_TEARING uint32
}
type dxgiHDRMetadataType struct {
	DXGI_HDR_METADATA_TYPE_NONE  uint32
	DXGI_HDR_METADATA_TYPE_HDR10 uint32
}
type dataOfferResourceFlags struct {
	DXGI_OFFER_RESOURCE_FLAG_ALLOW_DECOMMIT uint32
}
type dxgiReclaimResourceResults struct {
	DXGI_RECLAIM_RESOURCE_RESULT_OK            uint32
	DXGI_RECLAIM_RESOURCE_RESULT_DISCARDED     uint32
	DXGI_RECLAIM_RESOURCE_RESULT_NOT_COMMITTED uint32
}

var DXGI_FEATURE = dxgiFeature{
	DXGI_FEATURE_PRESENT_ALLOW_TEARING: 0,
}
var DXGI_HDR_METADATA_TYPE = dxgiHDRMetadataType{
	DXGI_HDR_METADATA_TYPE_NONE:  0,
	DXGI_HDR_METADATA_TYPE_HDR10: 1,
}
var DXGI_OFFER_RESOURCE_FLAGS = dataOfferResourceFlags{
	DXGI_OFFER_RESOURCE_FLAG_ALLOW_DECOMMIT: 0x1,
}
var DXGI_RECLAIM_RESOURCE_RESULTS = dxgiReclaimResourceResults{
	DXGI_RECLAIM_RESOURCE_RESULT_OK:            0,
	DXGI_RECLAIM_RESOURCE_RESULT_DISCARDED:     1,
	DXGI_RECLAIM_RESOURCE_RESULT_NOT_COMMITTED: 2,
}

//++++++++++++++++++++ dxgi1_6.h ++++++++++++++++++++
type dxgiAdapterFlag3 struct {
	DXGI_ADAPTER_FLAG3_NONE                         uint32
	DXGI_ADAPTER_FLAG3_REMOTE                       uint32
	DXGI_ADAPTER_FLAG3_SOFTWARE                     uint32
	DXGI_ADAPTER_FLAG3_ACG_COMPATIBLE               uint32
	DXGI_ADAPTER_FLAG3_SUPPORT_MONITORED_FENCES     uint32
	DXGI_ADAPTER_FLAG3_SUPPORT_NON_MONITORED_FENCES uint32
	DXGI_ADAPTER_FLAG3_KEYED_MUTEX_CONFORMANCE      uint32
	DXGI_ADAPTER_FLAG3_FORCE_DWORD                  uint32
}
type dxgiGPUPreference struct {
	DXGI_GPU_PREFERENCE_UNSPECIFIED      uint32
	DXGI_GPU_PREFERENCE_MINIMUM_POWER    uint32
	DXGI_GPU_PREFERENCE_HIGH_PERFORMANCE uint32
}
type dxgiHardwareCompositionSupportFlags struct {
	DXGI_HARDWARE_COMPOSITION_SUPPORT_FLAG_FULLSCREEN       uint32
	DXGI_HARDWARE_COMPOSITION_SUPPORT_FLAG_WINDOWED         uint32
	DXGI_HARDWARE_COMPOSITION_SUPPORT_FLAG_CURSOR_STRETCHED uint32
}

var DXGI_ADAPTER_FLAG3 = dxgiAdapterFlag3{
	DXGI_ADAPTER_FLAG3_NONE:                         0,
	DXGI_ADAPTER_FLAG3_REMOTE:                       1,
	DXGI_ADAPTER_FLAG3_SOFTWARE:                     2,
	DXGI_ADAPTER_FLAG3_ACG_COMPATIBLE:               4,
	DXGI_ADAPTER_FLAG3_SUPPORT_MONITORED_FENCES:     8,
	DXGI_ADAPTER_FLAG3_SUPPORT_NON_MONITORED_FENCES: 0x10,
	DXGI_ADAPTER_FLAG3_KEYED_MUTEX_CONFORMANCE:      0x20,
	DXGI_ADAPTER_FLAG3_FORCE_DWORD:                  0xffffffff,
}
var DXGI_GPU_PREFERENCE = dxgiGPUPreference{
	DXGI_GPU_PREFERENCE_UNSPECIFIED:      0,
	DXGI_GPU_PREFERENCE_MINIMUM_POWER:    1,
	DXGI_GPU_PREFERENCE_HIGH_PERFORMANCE: 2,
}
var DXGI_HARDWARE_COMPOSITION_SUPPORT_FLAGS = dxgiHardwareCompositionSupportFlags{
	DXGI_HARDWARE_COMPOSITION_SUPPORT_FLAG_FULLSCREEN:       1,
	DXGI_HARDWARE_COMPOSITION_SUPPORT_FLAG_WINDOWED:         2,
	DXGI_HARDWARE_COMPOSITION_SUPPORT_FLAG_CURSOR_STRETCHED: 4,
}

//++++++++++++++++++++ struct ++++++++++++++++++++
//++++++++++++++++++++ dxgi.h ++++++++++++++++++++
type DXGI_ADAPTER_DESC struct {
	Description           [128]C.wchar_t
	VendorId              uint32
	DeviceId              uint32
	SubSysId              uint32
	Revision              uint32
	DedicatedVideoMemory  SIZE_T
	DedicatedSystemMemory SIZE_T
	SharedSystemMemory    SIZE_T
	AdapterLuid           LUID
}
type DXGI_ADAPTER_DESC1 struct {
	Description           [128]C.wchar_t
	VendorId              uint32
	DeviceId              uint32
	SubSysId              uint32
	Revision              uint32
	DedicatedVideoMemory  SIZE_T
	DedicatedSystemMemory SIZE_T
	SharedSystemMemory    SIZE_T
	AdapterLuid           LUID
	Flags                 uint32
}
type DXGI_DISPLAY_COLOR_SPACE struct {
	PrimaryCoordinates [8][2]float32
	WhitePoints        [16][2]float32
}
type DXGI_FRAME_STATISTICS struct {
	PresentCount        uint32
	PresentRefreshCount uint32
	SyncRefreshCount    uint32
	SyncQPCTime         LargeInteger
	SyncGPUTime         LargeInteger
}
type DXGI_MAPPED_RECT struct {
	Pitch int32
	Bits  *byte
}
type DXGI_MAPPED_RECTstruct struct {
	Pitch uint32
	Bits  *byte
}
type DXGIOutputDesc struct {
	DeviceName         [32]C.wchar_t
	DesktopCoordinates RECT
	AttachedToDesktop  bool
	Rotation           uint32
	Monitor            unsafe.Pointer
}
type DXGISharedResource struct {
	Handle HANDLE
}
type DXGISurfaceDesc struct {
	Width      uint32
	Height     uint32
	Format     uint32
	SampleDesc DXGISampleDesc
}
type DXGISwapChainDesc struct {
	BufferDesc   DXGIModeDesc
	SampleDesc   DXGISampleDesc
	BufferUsage  DXGIUsage
	BufferCount  uint32
	OutputWindow HWND
	Windowed     bool
	SwapEffect   uint32
	Flags        uint32
}

//++++++++++++++++++++ dxgi1_2.h ++++++++++++++++++++
type DXGI_ADAPTER_DESC2 struct {
	Description                   [128]C.wchar_t
	VendorId                      uint32
	DeviceId                      uint32
	SubSysId                      uint32
	Revision                      uint32
	DedicatedVideoMemory          SIZE_T
	DedicatedSystemMemory         SIZE_T
	SharedSystemMemory            SIZE_T
	AdapterLuid                   LUID
	Flags                         uint32
	GraphicsPreemptionGranularity uint32
	ComputePreemptionGranularity  uint32
}
type DXGI_MODE_DESC1 struct {
	Width            uint32
	Height           uint32
	RefreshRate      DXGI_RATIONAL
	Format           uint32
	ScanlineOrdering uint32
	Scaling          uint32
	Stereo           bool
}
type DXGI_OUTDUPL_DESC struct {
	ModeDesc                   DXGI_MODE_DESC
	Rotation                   uint32
	DesktopImageInSystemMemory bool
}
type DXGI_OUTDUPL_FRAME_INFO struct {
	LastPresentTime           uint64
	LastMouseUpdateTime       uint64
	AccumulatedFrames         uint32
	RectsCoalesced            bool
	ProtectedContentMaskedOut bool
	PointerPosition           DXGI_OUTDUPL_POINTER_POSITION
	TotalMetadataBufferSize   uint32
	PointerShapeBufferSize    uint32
}
type DXGI_OUTDUPL_MOVE_RECT struct {
	SourcePoint     POINT
	DestinationRect RECT
}
type DXGI_OUTDUPL_POINTER_POSITION struct {
	Position POINT
	Visible  bool
}
type DXGI_OUTDUPL_POINTER_SHAPE_INFO struct {
	Type    uint32
	Width   uint32
	Height  uint32
	Pitch   uint32
	HotSpot POINT
}
type DXGI_PRESENT_PARAMETERS struct {
	DirtyRectsCount uint32
	DirtyRects      *RECT
	ScrollRect      *RECT
	ScrollOffset    *POINT
}
type DXGI_SWAP_CHAIN_DESC1 struct {
	Width       uint32
	Height      uint32
	Format      uint32
	Stereo      bool
	SampleDesc  DXGI_SAMPLE_DESC
	BufferUsage DXGI_USAGE
	BufferCount uint32
	Scaling     uint32
	SwapEffect  uint32
	AlphaMode   uint32
	Flags       uint32
}
type DXGI_SWAP_CHAIN_FULLSCREEN_DESC struct {
	RefreshRate      DXGI_RATIONAL
	ScanlineOrdering uint32
	Scaling          uint32
	Windowed         bool
}

//++++++++++++++++++++ dxgi1_3.h ++++++++++++++++++++
type DXGI_DECODE_SWAP_CHAIN_DESC struct {
	Flags uint32
}
type DXGI_FRAME_STATISTICS_MEDIA struct {
	PresentCount            uint32
	PresentRefreshCount     uint32
	SyncRefreshCount        uint32
	SyncQPCTime             LARGE_INTEGER
	SyncGPUTime             LARGE_INTEGER
	CompositionMode         uint32
	ApprovedPresentDuration uint32
}
type DXGI_MATRIX_3X2_F struct {
	_11 uint32
	_12 uint32
	_21 uint32
	_22 uint32
	_31 uint32
	_32 uint32
}

//++++++++++++++++++++ dxgi1_4.h ++++++++++++++++++++
type DXGI_QUERY_VIDEO_MEMORY_INFO struct {
	Budget                  uint64
	CurrentUsage            uint64
	AvailableForReservation uint64
	CurrentReservation      uint64
}

//++++++++++++++++++++ dxgi1_5.h ++++++++++++++++++++
type DXGI_HDR_METADATA_HDR10 struct {
	RedPrimary                [2]uint16
	GreenPrimary              [2]uint16
	BluePrimary               [2]uint16
	WhitePoint                [2]uint16
	MaxMasteringLuminance     uint32
	MinMasteringLuminance     uint32
	MaxContentLightLevel      uint16
	MaxFrameAverageLightLevel uint16
}

//++++++++++++++++++++ dxgi1_6.h ++++++++++++++++++++
type DXGI_ADAPTER_DESC3 struct {
	Description                   [128]C.wchar_t
	VendorId                      uint32
	DeviceId                      uint32
	SubSysId                      uint32
	Revision                      uint32
	DedicatedVideoMemory          SIZE_T
	DedicatedSystemMemory         SIZE_T
	SharedSystemMemory            SIZE_T
	AdapterLuid                   *GUID
	Flags                         uint32
	GraphicsPreemptionGranularity uint32
	ComputePreemptionGranularity  uint32
}
type DXGI_OUTPUT_DESC1 struct {
	DeviceName            [32]C.wchar_t
	DesktopCoordinates    RECT
	AttachedToDesktop     BOOL
	Rotation              uint32
	Monitor               HMONITOR
	BitsPerColor          uint32
	ColorSpace            uint32
	RedPrimary            [2]float32
	GreenPrimary          [2]float32
	BluePrimary           [2]float32
	WhitePoint            [2]float32
	MinLuminance          float32
	MaxLuminance          float32
	MaxFullFrameLuminance float32
}
