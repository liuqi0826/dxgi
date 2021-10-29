package dxgi

//++++++++++++++++++++ typedef ++++++++++++++++++++
type (
	DXGI_RGBA = D3DCOLORVALUE
)

//++++++++++++++++++++ enum ++++++++++++++++++++

type dxgiModeScanlineOrder struct {
	DXGI_MODE_SCANLINE_ORDER_UNSPECIFIED       uint32
	DXGI_MODE_SCANLINE_ORDER_PROGRESSIVE       uint32
	DXGI_MODE_SCANLINE_ORDER_UPPER_FIELD_FIRST uint32
	DXGI_MODE_SCANLINE_ORDER_LOWER_FIELD_FIRST uint32
}
type dxgiModeScaling struct {
	DXGI_MODE_SCALING_UNSPECIFIED uint32
	DXGI_MODE_SCALING_CENTERED    uint32
	DXGI_MODE_SCALING_STRETCHED   uint32
}
type dxgiModeRotation struct {
	DXGI_MODE_ROTATION_UNSPECIFIED uint32
	DXGI_MODE_ROTATION_IDENTITY    uint32
	DXGI_MODE_ROTATION_ROTATE90    uint32
	DXGI_MODE_ROTATION_ROTATE180   uint32
	DXGI_MODE_ROTATION_ROTATE270   uint32
}

var DXGI_MODE_SCANLINE_ORDER = dxgiModeScanlineOrder{
	DXGI_MODE_SCANLINE_ORDER_UNSPECIFIED:       0,
	DXGI_MODE_SCANLINE_ORDER_PROGRESSIVE:       1,
	DXGI_MODE_SCANLINE_ORDER_UPPER_FIELD_FIRST: 2,
	DXGI_MODE_SCANLINE_ORDER_LOWER_FIELD_FIRST: 3,
}
var DXGI_MODE_SCALING = dxgiModeScaling{
	DXGI_MODE_SCALING_UNSPECIFIED: 0,
	DXGI_MODE_SCALING_CENTERED:    1,
	DXGI_MODE_SCALING_STRETCHED:   2,
}
var DXGI_MODE_ROTATION = dxgiModeRotation{
	DXGI_MODE_ROTATION_UNSPECIFIED: 0,
	DXGI_MODE_ROTATION_IDENTITY:    1,
	DXGI_MODE_ROTATION_ROTATE90:    2,
	DXGI_MODE_ROTATION_ROTATE180:   3,
	DXGI_MODE_ROTATION_ROTATE270:   4,
}

//++++++++++++++++++++ struct ++++++++++++++++++++
type DXGI_RGB struct {
	Red   float32
	Green float32
	Blue  float32
}
type D3DCOLORVALUE struct {
	R float32
	G float32
	B float32
	A float32
}
type DXGI_GAMMA_CONTROL struct {
	Scale      DXGI_RGB
	Offset     DXGI_RGB
	GammaCurve [1025]DXGI_RGB
}
type DXGI_GAMMA_CONTROL_CAPABILITIES struct {
	ScaleAndOffsetSupported bool
	MaxConvertedValue       float32
	MinConvertedValue       float32
	NumGammaControlPoints   uint32
	ControlPointPositions   [1025]float32
}
type DXGI_MODE_DESC struct {
	Width            uint32
	Height           uint32
	RefreshRate      DXGI_RATIONAL
	Format           uint32
	ScanlineOrdering uint32
	Scaling          uint32
}
type DXGI_JPEG_DC_HUFFMAN_TABLE struct {
	CodeCounts [12]byte
	CodeValues [12]byte
}
type DXGI_JPEG_AC_HUFFMAN_TABLE struct {
	CodeCounts [16]byte
	CodeValues [162]byte
}
type DXGI_JPEG_QUANTIZATION_TABLE struct {
	Elements [64]byte
}
