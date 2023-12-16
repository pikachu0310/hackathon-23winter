package images

import (
	_ "embed"
)

//go:embed aqua.png
var TestKemonoImageAqua []byte

//go:embed fire.png
var TestKemonoImageFire []byte

//go:embed aquafire.png
var TestKemonoImageAquaFire []byte
