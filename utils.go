package dxgi

// #include <stdlib.h>
import "C"
import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func GetError(hresult uint32) error {
	// fmt.Println(hresult)
	switch hresult {
	case 0x087a0001:
		err := errors.New("DXGI_STATUS_OCCLUDED")
		return err
	case 0x087a0002:
		err := errors.New("DXGI_STATUS_CLIPPED")
		return err
	case 0x087a0004:
		err := errors.New("DXGI_STATUS_NO_REDIRECTION")
		return err
	case 0x087a0005:
		err := errors.New("DXGI_STATUS_NO_DESKTOP_ACCESS")
		return err
	case 0x087a0006:
		err := errors.New("DXGI_STATUS_GRAPHICS_VIDPN_SOURCE_IN_USE")
		return err
	case 0x087a0007:
		err := errors.New("DXGI_STATUS_MODE_CHANGED")
		return err
	case 0x087a0008:
		err := errors.New("DXGI_STATUS_MODE_CHANGE_IN_PROGRESS")
		return err
	case 0x087a0009:
		err := errors.New("DXGI_STATUS_UNOCCLUDED")
		return err
	case 0x087a000a:
		err := errors.New("DXGI_STATUS_DDA_WAS_STILL_DRAWING")
		return err
	case 0x087a002f:
		err := errors.New("DXGI_STATUS_PRESENT_REQUIRED")
		return err
	case 0x887A0001:
		err := errors.New("DXGI_ERROR_INVALID_CALL")
		return err
	case 0x887A0002:
		err := errors.New("DXGI_ERROR_NOT_FOUND")
		return err
	case 0x887A0003:
		err := errors.New("DXGI_ERROR_MORE_DATA")
		return err
	case 0x887A0004:
		err := errors.New("DXGI_ERROR_UNSUPPORTED")
		return err
	case 0x887A0005:
		err := errors.New("DXGI_ERROR_DEVICE_REMOVED")
		return err
	case 0x887A0006:
		err := errors.New("DXGI_ERROR_DEVICE_HUNG")
		return err
	case 0x887A0007:
		err := errors.New("DXGI_ERROR_DEVICE_RESET")
		return err
	case 0x887A000A:
		err := errors.New("DXGI_ERROR_WAS_STILL_DRAWING")
		return err
	case 0x887A000B:
		err := errors.New("DXGI_ERROR_FRAME_STATISTICS_DISJOINT")
		return err
	case 0x887A000C:
		err := errors.New("DXGI_ERROR_GRAPHICS_VIDPN_SOURCE_IN_USE")
		return err
	case 0x887A0020:
		err := errors.New("DXGI_ERROR_DRIVER_INTERNAL_ERROR")
		return err
	case 0x887A0021:
		err := errors.New("DXGI_ERROR_NONEXCLUSIVE")
		return err
	case 0x887A0022:
		err := errors.New("DXGI_ERROR_NOT_CURRENTLY_AVAILABLE")
		return err
	case 0x887A0023:
		err := errors.New("DXGI_ERROR_REMOTE_CLIENT_DISCONNECTED")
		return err
	case 0x887A0024:
		err := errors.New("DXGI_ERROR_REMOTE_OUTOFMEMORY")
		return err
	case 0x887A0026:
		err := errors.New("DXGI_ERROR_ACCESS_LOST")
		return err
	case 0x887A0027:
		err := errors.New("DXGI_ERROR_WAIT_TIMEOUT")
		return err
	case 0x887A0028:
		err := errors.New("DXGI_ERROR_SESSION_DISCONNECTED")
		return err
	case 0x887A0029:
		err := errors.New("DXGI_ERROR_RESTRICT_TO_OUTPUT_STALE")
		return err
	case 0x887A002A:
		err := errors.New("DXGI_ERROR_CANNOT_PROTECT_CONTENT")
		return err
	case 0x887A002B:
		err := errors.New("DXGI_ERROR_ACCESS_DENIED")
		return err
	case 0x887A002C:
		err := errors.New("DXGI_ERROR_NAME_ALREADY_EXISTS")
		return err
	case 0x887A002D:
		err := errors.New("DXGI_ERROR_SDK_COMPONENT_MISSING")
		return err
	}
	return nil
}

func boolToInt(value bool) uintptr {
	if value {
		return 1
	}
	return 0
}
func intToBool(value uintptr) bool {
	if value != 0 {
		return true
	}
	return false
}
func stringToChar(str string) *uint8 {
	if !strings.HasSuffix(str, "\x00") {
		str = str + "\x00"
		// panic("str argument missing null terminator: " + str)
	}
	header := (*reflect.StringHeader)(unsafe.Pointer(&str))
	return (*uint8)(unsafe.Pointer(header.Data))
}

func charToString(cstr *uint8) string {
	return C.GoString((*C.char)(unsafe.Pointer(cstr)))
}
func stringArrayToCharArray(strs ...string) (cstrs **uint8, free func()) {
	if len(strs) == 0 {
		panic("Strs: expected at least 1 string")
	}
	var list []string
	for _, s := range strs {
		list = append(list, s+"\x00")
	}
	// Allocate a contiguous array large enough to hold all the strings' contents.
	n := 0
	for i := range list {
		n += len(list[i])
	}
	data := C.malloc(C.size_t(n))
	// Copy all the strings into data.
	dataSlice := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(data),
		Len:  n,
		Cap:  n,
	}))
	css := make([]*uint8, len(list)) // Populated with pointers to each string.
	offset := 0
	for i := range list {
		copy(dataSlice[offset:offset+len(list[i])], list[i][:]) // Copy strs[i] into proper data location.
		css[i] = (*uint8)(unsafe.Pointer(&dataSlice[offset]))   // Set a pointer to it.
		offset += len(list[i])
	}
	return (**uint8)(&css[0]), func() { C.free(data) }
}

func numberToPointer(value interface{}) unsafe.Pointer {
	if value == nil {
		return unsafe.Pointer(nil)
	}
	v := reflect.ValueOf(value)
	return unsafe.Pointer(v.UnsafeAddr())
}

func sliceToPointer(value interface{}) unsafe.Pointer {
	if value == nil {
		return unsafe.Pointer(nil)
	}
	v := reflect.ValueOf(value)
	return unsafe.Pointer(v.Index(0).UnsafeAddr())
}

func offsetPointer(ptr int32) unsafe.Pointer {
	return unsafe.Pointer(uintptr(ptr))
}
func ptr(data interface{}) unsafe.Pointer {
	if data == nil {
		return unsafe.Pointer(nil)
	}
	var addr unsafe.Pointer
	v := reflect.ValueOf(data)
	switch v.Type().Kind() {
	case reflect.Ptr:
		e := v.Elem()
		switch e.Kind() {
		case
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			addr = unsafe.Pointer(e.UnsafeAddr())
		default:
			panic(fmt.Errorf("unsupported pointer to type %s; must be a slice or pointer to a singular scalar value or the first element of an array or slice", e.Kind()))
		}
	case reflect.Uintptr:
		addr = unsafe.Pointer(v.Pointer())
	case reflect.Slice:
		addr = unsafe.Pointer(v.Index(0).UnsafeAddr())
	default:
		panic(fmt.Errorf("unsupported type %s; must be a slice or pointer to a singular scalar value or the first element of an array or slice", v.Type()))
	}
	return addr
}
func ptrOffset(offset int32) unsafe.Pointer {
	return unsafe.Pointer(uintptr(offset))
}
func goStr(cstr *uint8) string {
	return C.GoString((*C.char)(unsafe.Pointer(cstr)))
}
