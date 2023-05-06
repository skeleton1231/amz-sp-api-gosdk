# InventorySummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asin** | **string** | The Amazon Standard Identification Number (ASIN) of an item. | [optional] [default to null]
**FnSku** | **string** | Amazon&#x27;s fulfillment network SKU identifier. | [optional] [default to null]
**SellerSku** | **string** | The seller SKU of the item. | [optional] [default to null]
**Condition** | **string** | The condition of the item as described by the seller (for example, New Item). | [optional] [default to null]
**InventoryDetails** | [***InventoryDetails**](InventoryDetails.md) |  | [optional] [default to null]
**LastUpdatedTime** | [**time.Time**](time.Time.md) | The date and time that any quantity was last updated. | [optional] [default to null]
**ProductName** | **string** | The localized language product title of the item within the specific marketplace. | [optional] [default to null]
**TotalQuantity** | **int32** | The total number of units in an inbound shipment or in Amazon fulfillment centers. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

