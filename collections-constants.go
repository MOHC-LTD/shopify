package shopify

const (
	// CollectionTypeSmart - specifies if a collection is a smart collection.
	CollectionTypeSmart = "smart"
	// CollectionTypeCustom - specifies if a collection is a custom collection.
	CollectionTypeCustom = "custom"
)

const (
	// CollectionSortOrderAlphaAsc - Alphabetically, in ascending order (A - Z).
	CollectionSortOrderAlphaAsc = "alpha-asc"
	// CollectionSortOrderAlphaDesc - Alphabetically, in descending order (Z - A).
	CollectionSortOrderAlphaDesc = "alpha-desc"
	// CollectionSortOrderBestSelling - By best-selling products.
	CollectionSortOrderBestSelling = "best-selling"
	// CollectionSortOrderCreated - By date created, in ascending order (oldest - newest).
	CollectionSortOrderCreated = "created"
	// CollectionSortOrderCreatedDesc - By date created, in descending order (newest - oldest).
	CollectionSortOrderCreatedDesc = "created-desc"
	// CollectionSortOrderManual - In the order set manually by the shop owner.
	CollectionSortOrderManual = "manual"
	// CollectionSortOrderPriceAsc - By price, in ascending order (lowest - highest).
	CollectionSortOrderPriceAsc = "price-asc"
	// CollectionSortOrderPriceDesc - By price, in descending order (highest - lowest).
	CollectionSortOrderPriceDesc = "price-desc"
)

const (
	// CollectionPublishedAtWeb - The collection is published to the Online Store channel but not published to the Point of Sale channel.
	CollectionPublishedAtWeb = "web"
	// CollectionPublishedAtGlobal - The collection is published to both the Online Store channel and the Point of Sale channel.
	CollectionPublishedAtGlobal = "global"
)
