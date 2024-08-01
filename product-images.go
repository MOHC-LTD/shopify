package shopify

import "time"

// ProductImages is a collection of product images
type ProductImages []ProductImage

// ProductImage is an image of the product
type ProductImage struct {
	Image
	// ID is a unique numeric identifier for the product image.
	ID int64
	// Position is used to order the product image in the list.
	/*
	   The first product image is at position 1 and is the "main" image for the product.
	*/
	Position int
	// ProductID is unique numeric identifier of the product associated with the image.
	ProductID int64
	// Src specifies the location of the product image
	Src string
	// VariantIDs is a array of variant ids associated with the image.
	VariantIDs []int64
	// UpdatedAt is the date and time when the product image was last modified.
	UpdatedAt time.Time
}

// ProductImageQuery are properties that can be used to filter the returned product images
// See https://shopify.dev/api/admin/rest/reference/products/product-image
type ProductImageQuery struct {
	/*
		Restrict product images to after the specified ID
	*/
	SinceID int64
}

// ProductImageRepository maintains the product images for products in the shop.
type ProductImageRepository interface {
	// List gets all of the product images by product id.
	List(productID int64, query ProductImageQuery) (ProductImages, error)
}
