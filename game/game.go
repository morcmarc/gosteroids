package game

const (
	WIDTH  int    = 600
	HEIGHT int    = 600
	TITLE  string = "Gosteroids"
)

func Start() {
	ctrlChnl := make(chan uint8)
	InitControls(ctrlChnl)
	InitWindow(WIDTH, HEIGHT, TITLE, ctrlChnl)
}
