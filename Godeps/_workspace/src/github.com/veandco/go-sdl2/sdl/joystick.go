package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

const (
	HAT_CENTERED  = C.SDL_HAT_CENTERED
	HAT_UP        = C.SDL_HAT_UP
	HAT_RIGHT     = C.SDL_HAT_RIGHT
	HAT_DOWN      = C.SDL_HAT_DOWN
	HAT_LEFT      = C.SDL_HAT_LEFT
	HAT_RIGHTUP   = C.SDL_HAT_RIGHTUP
	HAT_RIGHTDOWN = C.SDL_HAT_RIGHTDOWN
	HAT_LEFTUP    = C.SDL_HAT_LEFTUP
	HAT_LEFTDOWN  = C.SDL_HAT_LEFTDOWN
)

type Joystick C.SDL_Joystick
type JoystickGUID C.SDL_JoystickGUID
type JoystickID C.SDL_JoystickID

func (joy *Joystick) cptr() *C.SDL_Joystick {
    return (*C.SDL_Joystick)(unsafe.Pointer(joy))
}

func (guid JoystickGUID) c() C.SDL_JoystickGUID {
    return C.SDL_JoystickGUID(guid)
}

func NumJoysticks() int {
	return (int)(C.SDL_NumJoysticks())
}

func JoystickNameForIndex(index int) string {
	return (C.GoString)(C.SDL_JoystickNameForIndex(C.int(index)))
}

func JoystickOpen(index int) *Joystick {
	return (*Joystick)(C.SDL_JoystickOpen(C.int(index)))
}

func (joy *Joystick) Name() string {
	return (C.GoString)(C.SDL_JoystickName(joy.cptr()))
}

func JoystickGetDeviceGUID(index int) JoystickGUID {
	return (JoystickGUID)(C.SDL_JoystickGetDeviceGUID(C.int(index)))
}

func (joy *Joystick) GetGUID() JoystickGUID {
	return (JoystickGUID)(C.SDL_JoystickGetGUID(joy.cptr()))
}

func JoystickGetGUIDString(guid JoystickGUID, pszGUID string, cbGUID int) {
	_pszGUID := C.CString(pszGUID)
	defer C.SDL_free(unsafe.Pointer(_pszGUID))
	C.SDL_JoystickGetGUIDString(guid.c(), _pszGUID, C.int(cbGUID))
}

func JoystickGetGUIDFromString(pchGUID string) JoystickGUID {
	_pchGUID := C.CString(pchGUID)
	defer C.SDL_free(unsafe.Pointer(_pchGUID))
	return (JoystickGUID)(C.SDL_JoystickGetGUIDFromString(_pchGUID))
}

func (joy *Joystick) GetAttached() bool {
	return C.SDL_JoystickGetAttached(joy.cptr()) > 0
}

func (joy *Joystick) InstanceID() JoystickID {
	return (JoystickID)(C.SDL_JoystickInstanceID(joy.cptr()))
}

func (joy *Joystick) NumAxes() int {
	return (int)(C.SDL_JoystickNumAxes(joy.cptr()))
}

func (joy *Joystick) NumBalls() int {
	return (int)(C.SDL_JoystickNumBalls(joy.cptr()))
}

func (joy *Joystick) NumHats() int {
	return (int)(C.SDL_JoystickNumHats(joy.cptr()))
}

func (joy *Joystick) NumButtons() int {
	return (int)(C.SDL_JoystickNumButtons(joy.cptr()))
}

func Update() {
	C.SDL_JoystickUpdate()
}

func JoystickEventState(state int) int {
	return (int)(C.SDL_JoystickEventState(C.int(state)))
}

func (joy *Joystick) GetAxis(axis int) int16 {
	return (int16)(C.SDL_JoystickGetAxis(joy.cptr(), C.int(axis)))
}

func (joy *Joystick) GetHat(hat int) byte {
	return (byte)(C.SDL_JoystickGetHat(joy.cptr(), C.int(hat)))
}

func (joy *Joystick) GetBall(ball int, dx, dy *int) int {
	_dx := (*C.int)(unsafe.Pointer(dx))
	_dy := (*C.int)(unsafe.Pointer(dy))
	return (int)(C.SDL_JoystickGetBall(joy.cptr(), C.int(ball), _dx, _dy))
}

func (joy *Joystick) GetButton(button int) byte {
	return (byte)(C.SDL_JoystickGetButton(joy.cptr(), C.int(button)))
}

func (joy *Joystick) Close() {
	C.SDL_JoystickClose(joy.cptr())
}
