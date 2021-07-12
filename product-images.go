package shopify

import "time"

// ProductImages is a collection of Images
type ProductImages []ProductImage

// ProductImage is an image of a the product
type ProductImage struct {
	// ID is a unique numeric identifier for the product image.
	ID int64
	// CreatedAt is the date and time when the product image was created.
	CreatedAt time.Time
	// Position is used to order the product image in the list.
	/*
	   The first product image is at position 1 and is the "main" image for the product.
	*/
	Position int
	// ProductID is unique numeric identifier of the product associated with the image.
	ProductID int64
	// VariantIDs is a array of variant ids associated with the image.
	VariantIDs []int64
	// SRC Specifies the location of the product image.
	/*
	   This parameter supports URL filters that you can use to retrieve modified copies of the image.
	   For example, add _small, to the filename to retrieve a scaled copy of the image at 100 x 100 px
	   (for example, ipod-nano_small.png), or add _2048x2048 to retrieve a copy of the image constrained
	   at 2048 x 2048 px resolution (for example, ipod-nano_2048x2048.png).
	*/
	SRC string
	// Width dimension of the image which is determined on upload.
	Width int
	// Height dimension of the image which is determined on upload.
	Height int
	// UpdatedAt is the date and time when the product image was last modified.
	UpdatedAt time.Time
}

// ProductImageQuery are properties that can be used to filter the returned Images
// See https://shopify.dev/api/admin/rest/reference/products/product-image
type ProductImageQuery struct {
	/*
		Return only products images .
	*/
	SinceIDs []int64
}

// ProductImageRepository maintains the Images of a shop.
type ProductImageRepository interface {
	// List gets all of the Images
	List(query ProductImageQuery) (ProductImages, error)
}
