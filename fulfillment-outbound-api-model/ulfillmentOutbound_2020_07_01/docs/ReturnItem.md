# ReturnItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SellerReturnItemId** | **string** | An identifier assigned by the seller to the return item. | [default to null]
**SellerFulfillmentOrderItemId** | **string** | The identifier assigned to the item by the seller when the fulfillment order was created. | [default to null]
**AmazonShipmentId** | **string** | The identifier for the shipment that is associated with the return item. | [default to null]
**SellerReturnReasonCode** | **string** | The return reason code assigned to the return item by the seller. | [default to null]
**ReturnComment** | **string** | An optional comment about the return item. | [optional] [default to null]
**AmazonReturnReasonCode** | **string** | The return reason code that the Amazon fulfillment center assigned to the return item. | [optional] [default to null]
**Status** | [***FulfillmentReturnItemStatus**](FulfillmentReturnItemStatus.md) |  | [default to null]
**StatusChangedDate** | [***time.Time**](time.Time.md) |  | [default to null]
**ReturnAuthorizationId** | **string** | Identifies the return authorization used to return this item. See ReturnAuthorization. | [optional] [default to null]
**ReturnReceivedCondition** | [***ReturnItemDisposition**](ReturnItemDisposition.md) |  | [optional] [default to null]
**FulfillmentCenterId** | **string** | The identifier for the Amazon fulfillment center that processed the return item. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

