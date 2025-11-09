package wayland

type Pixels int

type PixelSize struct {
	Width  Pixels
	Height Pixels
}

var VirtualMonitorSize = PixelSize{
	Width:  640,
	Height: 480,
}
