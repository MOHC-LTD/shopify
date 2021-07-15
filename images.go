package shopify

import "time"

// Images contains a list of images
type Images []Image

<<<<<<< HEAD
// Image is a basic image object with the essential content
type Image struct {
	// CreatedAt is the date and time when the product image was created.
=======
// Image is a image
type Image struct {
	// CreatedAt is the time and date when the image was added
>>>>>>> develop
	CreatedAt time.Time
	// SRC is the source URL that specifies the location of the image.
	SRC string
	// Width is the width of the image in pixels.
	Width int
	// Height is the height of the image in pixels.
	Height int
<<<<<<< HEAD
	// Alt is the alternative text that describes the collection image.
=======
	// Alt is the alternative text that describes the image.
>>>>>>> develop
	Alt string
}
