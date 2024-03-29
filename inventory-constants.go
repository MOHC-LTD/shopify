package shopify

// InventoryPolicyAllowWhenOutOfStock is an inventory policy to indicate that customers ARE allowed to place orders for a product variant if it's out of stock.
const InventoryPolicyAllowWhenOutOfStock = "continue"

// InventoryPolicyDenyWhenOutOfStock is an inventory policy to indicate that customers are NOT allowed to place orders for a product variant if it's out of stock.
const InventoryPolicyDenyWhenOutOfStock = "deny"
