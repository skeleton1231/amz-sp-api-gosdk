/*
 * Selling Partner API for Direct Fulfillment Inventory Updates
 *
 * The Selling Partner API for Direct Fulfillment Inventory Updates provides programmatic access to a direct fulfillment vendor's inventory updates.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InventoryUpdate struct {
	SellingParty *PartyIdentification `json:"sellingParty"`
	// When true, this request contains a full feed. Otherwise, this request contains a partial feed. When sending a full feed, you must send information about all items in the warehouse. Any items not in the full feed are updated as not available. When sending a partial feed, only include the items that need an update to inventory. The status of other items will remain unchanged.
	IsFullUpdate bool `json:"isFullUpdate"`
	// A list of inventory items with updated details, including quantity available.
	Items []ItemDetails `json:"items"`
}
