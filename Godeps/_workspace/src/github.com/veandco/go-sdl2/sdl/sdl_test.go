package sdl

import "testing"
import "runtime"

func TestInitQuit(t *testing.T) {
	Init(0)
	subs := []uint32{INIT_TIMER, INIT_AUDIO, INIT_VIDEO, INIT_JOYSTICK,
		INIT_HAPTIC, INIT_GAMECONTROLLER}

	for i := 0; i < len(subs); i++ {
		if (runtime.GOOS == "freebsd") && (subs[i] == INIT_HAPTIC) {
			// FreeBSD does not support the haptic subsystem
			continue
		}
		if Init(subs[i]) != 0 {
			t.Errorf("Error on Init(%d): %s", subs[i], GetError())
		}
		if WasInit(subs[i]) != subs[i] {
			t.Errorf("Init(%d): subsystem not initialized", subs[i])
		}
		QuitSubSystem(subs[i])
		if WasInit(subs[i]) == subs[i] {
			t.Errorf("QuitSubSystem(%d): subsystem still initialized", subs[i])
		}
		if InitSubSystem(subs[i]) != 0 {
			t.Errorf("Error on Init(%d): %s", subs[i], GetError())
		}
		if WasInit(subs[i]) != subs[i] {
			t.Errorf("InitSubSystem(%d): subsystem not initialized", subs[i])
		}
		QuitSubSystem(subs[i])
		if WasInit(subs[i]) == subs[i] {
			t.Errorf("QuitSubSystem(%d): subsystem still initialized", subs[i])
		}
	}
	Quit()
}

func TestGetPlatform(t *testing.T) {
	goos := runtime.GOOS
	value := GetPlatform()
	if (goos == "linux" && value != "Linux") ||
		(goos == "freebsd" && value != "FreeBSD") ||
		(goos == "windows" && value != "Windows") ||
		(goos == "darwin" && value != "Mac OS X") {
		t.Errorf("platform mismatch: '%s' != '%s'", goos, value)
	}
}
