/*
 * Selling Partner API for Retail Procurement Shipments
 *
 * The Selling Partner API for Retail Procurement Shipments provides programmatic access to retail shipping data for vendors.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Details of the Pallet/Tare being shipped.
type Pallet struct {
	// A list of pallet identifiers.
	PalletIdentifiers []ContainerIdentification `json:"palletIdentifiers"`
	// Number of layers per pallet. Only applicable to container type Pallet.
	Tier int32 `json:"tier,omitempty"`
	// Number of cartons per layer on the pallet. Only applicable to container type Pallet.
	Block int32 `json:"block,omitempty"`
	Dimensions *Dimensions `json:"dimensions,omitempty"`
	Weight *Weight `json:"weight,omitempty"`
	CartonReferenceDetails *CartonReferenceDetails `json:"cartonReferenceDetails,omitempty"`
	// A list of container item details.
	Items []ContainerItem `json:"items,omitempty"`
}
