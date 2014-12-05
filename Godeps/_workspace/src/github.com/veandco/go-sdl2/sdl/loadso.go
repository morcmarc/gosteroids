package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

func LoadObject(sofile string) unsafe.Pointer {
	_sofile := C.CString(sofile)
	defer C.SDL_free(unsafe.Pointer(_sofile))
	return (unsafe.Pointer)(C.SDL_LoadObject(_sofile))
}

func LoadFunction(handle unsafe.Pointer, name string) unsafe.Pointer {
	_name := C.CString(name)
	defer C.SDL_free(unsafe.Pointer(_name))
	return (unsafe.Pointer)(C.SDL_LoadFunction(handle, _name))
}

func UnloadObject(handle unsafe.Pointer) {
	C.SDL_UnloadObject(handle)
}
