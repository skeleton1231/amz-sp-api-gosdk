/*
 * Selling Partner API for Fulfillment Inbound
 *
 * The Selling Partner API for Fulfillment Inbound lets you create applications that create and update inbound shipments of inventory to Amazon's fulfillment network.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// Condition : The condition of the item.
type Condition string

// List of Condition
const (
	NEW_ITEM_Condition Condition = "NewItem"
	NEW_WITH_WARRANTY_Condition Condition = "NewWithWarranty"
	NEW_OEM_Condition Condition = "NewOEM"
	NEW_OPEN_BOX_Condition Condition = "NewOpenBox"
	USED_LIKE_NEW_Condition Condition = "UsedLikeNew"
	USED_VERY_GOOD_Condition Condition = "UsedVeryGood"
	USED_GOOD_Condition Condition = "UsedGood"
	USED_ACCEPTABLE_Condition Condition = "UsedAcceptable"
	USED_POOR_Condition Condition = "UsedPoor"
	USED_REFURBISHED_Condition Condition = "UsedRefurbished"
	COLLECTIBLE_LIKE_NEW_Condition Condition = "CollectibleLikeNew"
	COLLECTIBLE_VERY_GOOD_Condition Condition = "CollectibleVeryGood"
	COLLECTIBLE_GOOD_Condition Condition = "CollectibleGood"
	COLLECTIBLE_ACCEPTABLE_Condition Condition = "CollectibleAcceptable"
	COLLECTIBLE_POOR_Condition Condition = "CollectiblePoor"
	REFURBISHED_WITH_WARRANTY_Condition Condition = "RefurbishedWithWarranty"
	REFURBISHED_Condition Condition = "Refurbished"
	CLUB_Condition Condition = "Club"
)
