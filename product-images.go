package shopify

import "time"

// ProductImages is a collection of Images
type ProductImages []ProductImage

// ProductImage is an image of a the product
type ProductImage struct {
	// ID is a unique numeric identifier for the product image.
	ID int64
	// Image contains the basic image content
	Image
	// Position is used to order the product image in the list.
	/*
	   The first product image is at position 1 and is the "main" image for the product.
	*/
	Position int
	// ProductID is unique numeric identifier of the product associated with the image.
	ProductID int64
	// VariantIDs is a array of variant ids associated with the image.
	VariantIDs []int64
	// UpdatedAt is the date and time when the product image was last modified.
	UpdatedAt time.Time
}

// ProductImageQuery are properties that can be used to filter the returned Images
// See https://shopify.dev/api/admin/rest/reference/products/product-image
type ProductImageQuery struct {
	/*
		Return products images since a set of Image IDs.
	*/
	SinceIDs []int64
}

// ProductImageRepository maintains the Images of a shop.
type ProductImageRepository interface {
	// List gets all of the product images.
	List(query ProductImageQuery) (ProductImages, error)
}
