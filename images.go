package shopify

import "time"

// Images contains a list of images
type Images []Image

// Image is a basic image object with the essential content
type Image struct {
	// CreatedAt is the date and time when the product image was created.
	CreatedAt time.Time
	// SRC is the source URL that specifies the location of the image.
	SRC string
	// Width is the width of the image in pixels.
	Width int
	// Height is the height of the image in pixels.
	Height int
	// Alt is the alternative text that describes the collection image.
	Alt string
}
