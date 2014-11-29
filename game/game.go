package game

const (
	Width  int    = 600
	Height int    = 600
	Title  string = "Gosteroids"
)

func Start() {
	ctrlChnl := make(chan uint8)
	InitControls(ctrlChnl)
	InitWindow(Width, Height, Title, ctrlChnl)
}
