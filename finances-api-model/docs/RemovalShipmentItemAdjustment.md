# RemovalShipmentItemAdjustment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemovalShipmentItemId** | **string** | An identifier for an item in a removal shipment. | [optional] [default to null]
**TaxCollectionModel** | **string** | The tax collection model applied to the item.  Possible values:  * MarketplaceFacilitator - Tax is withheld and remitted to the taxing authority by Amazon on behalf of the seller.  * Standard - Tax is paid to the seller and not remitted to the taxing authority by Amazon. | [optional] [default to null]
**FulfillmentNetworkSKU** | **string** | The Amazon fulfillment network SKU for the item. | [optional] [default to null]
**AdjustedQuantity** | **int32** | Adjusted quantity of removal shipmentItemAdjustment items. | [optional] [default to null]
**RevenueAdjustment** | [***Currency**](Currency.md) |  | [optional] [default to null]
**TaxAmountAdjustment** | [***Currency**](Currency.md) |  | [optional] [default to null]
**TaxWithheldAdjustment** | [***Currency**](Currency.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

