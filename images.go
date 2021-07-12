package shopify

import "time"

// Images contains a list of images
type Images []Image

// Image is a basic image object with the essential content
type Image struct {
	// CreatedAt is the date and time when the product image was created.
	CreatedAt time.Time
	// SRC Specifies the location of the product image.
	/*
	   This parameter supports URL filters that you can use to retrieve modified copies of the image.
	   For example, add _small, to the filename to retrieve a scaled copy of the image at 100 x 100 px
	   (for example, ipod-nano_small.png), or add _2048x2048 to retrieve a copy of the image constrained
	   at 2048 x 2048 px resolution (for example, ipod-nano_2048x2048.png).
	*/
	SRC string
	// Width is the width of the image in pixels.
	Width int
	// Height is the height of the image in pixels.
	Height int
	// Alt is the alternative text that describes the collection image.
	Alt string
}
