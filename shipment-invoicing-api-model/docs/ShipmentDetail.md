# ShipmentDetail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WarehouseId** | **string** | The Amazon-defined identifier for the warehouse. | [optional] [default to null]
**AmazonOrderId** | **string** | The Amazon-defined identifier for the order. | [optional] [default to null]
**AmazonShipmentId** | **string** | The Amazon-defined identifier for the shipment. | [optional] [default to null]
**PurchaseDate** | [**time.Time**](time.Time.md) | The date and time when the order was created. | [optional] [default to null]
**ShippingAddress** | [***Address**](Address.md) |  | [optional] [default to null]
**PaymentMethodDetails** | [***[]string**](array.md) |  | [optional] [default to null]
**MarketplaceId** | **string** | The identifier for the marketplace where the order was placed. | [optional] [default to null]
**SellerId** | **string** | The seller identifier. | [optional] [default to null]
**BuyerName** | **string** | The name of the buyer. | [optional] [default to null]
**BuyerCounty** | **string** | The county of the buyer. | [optional] [default to null]
**BuyerTaxInfo** | [***BuyerTaxInfo**](BuyerTaxInfo.md) |  | [optional] [default to null]
**MarketplaceTaxInfo** | [***MarketplaceTaxInfo**](MarketplaceTaxInfo.md) |  | [optional] [default to null]
**SellerDisplayName** | **string** | The seller’s friendly name registered in the marketplace. | [optional] [default to null]
**ShipmentItems** | [***[]ShipmentItem**](array.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

