# ItemSummaryByMarketplace

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceId** | **string** | A marketplace identifier. Identifies the Amazon marketplace for the listings item. | [default to null]
**Asin** | **string** | Amazon Standard Identification Number (ASIN) of the listings item. | [default to null]
**ProductType** | **string** | The Amazon product type of the listings item. | [default to null]
**ConditionType** | **string** | Identifies the condition of the listings item. | [optional] [default to null]
**Status** | **[]string** | Statuses that apply to the listings item. | [default to null]
**FnSku** | **string** | Fulfillment network stock keeping unit is an identifier used by Amazon fulfillment centers to identify each unique item. | [optional] [default to null]
**ItemName** | **string** | Name, or title, associated with an Amazon catalog item. | [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) | Date the listings item was created, in ISO 8601 format. | [default to null]
**LastUpdatedDate** | [**time.Time**](time.Time.md) | Date the listings item was last updated, in ISO 8601 format. | [default to null]
**MainImage** | [***ItemImage**](ItemImage.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

