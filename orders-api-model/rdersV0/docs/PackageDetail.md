# PackageDetail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageReferenceId** | **string** |  | [default to null]
**CarrierCode** | **string** | Identifies the carrier that will deliver the package. This field is required for all marketplaces, see [reference](https://developer-docs.amazon.com/sp-api/changelog/carriercode-value-required-in-shipment-confirmations-for-br-mx-ca-sg-au-in-jp-marketplaces). | [default to null]
**CarrierName** | **string** | Carrier Name that will deliver the package. Required when carrierCode is \&quot;Others\&quot;  | [optional] [default to null]
**ShippingMethod** | **string** | Ship method to be used for shipping the order. | [optional] [default to null]
**TrackingNumber** | **string** | The tracking number used to obtain tracking and delivery information. | [default to null]
**ShipDate** | [**time.Time**](time.Time.md) | The shipping date for the package. Must be in ISO-8601 date/time format. | [default to null]
**ShipFromSupplySourceId** | **string** | The unique identifier of the supply source. | [optional] [default to null]
**OrderItems** | [***[]ConfirmShipmentOrderItem**](array.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

