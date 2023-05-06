# OrderItemAcknowledgement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcknowledgementCode** | **string** | This indicates the acknowledgement code. | [default to null]
**AcknowledgedQuantity** | [***ItemQuantity**](ItemQuantity.md) |  | [default to null]
**ScheduledShipDate** | [**time.Time**](time.Time.md) | Estimated ship date per line item. Must be in ISO-8601 date/time format. | [optional] [default to null]
**ScheduledDeliveryDate** | [**time.Time**](time.Time.md) | Estimated delivery date per line item. Must be in ISO-8601 date/time format. | [optional] [default to null]
**RejectionReason** | **string** | Indicates the reason for rejection. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

