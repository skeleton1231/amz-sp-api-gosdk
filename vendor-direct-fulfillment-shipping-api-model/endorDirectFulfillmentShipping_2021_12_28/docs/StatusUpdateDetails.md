# StatusUpdateDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingNumber** | **string** | This is required to be provided for every package and should match with the trackingNumber sent for the shipment confirmation. | [default to null]
**StatusCode** | **string** | Indicates the shipment status code of the package that provides transportation information for Amazon tracking systems and ultimately for the final customer. | [default to null]
**ReasonCode** | **string** | Provides a reason code for the status of the package that will provide additional information about the transportation status. | [default to null]
**StatusDateTime** | [**time.Time**](time.Time.md) | The date and time when the shipment status was updated. This field is expected to be in ISO-8601 date/time format, with UTC time zone or UTC offset. For example, 2020-07-16T23:00:00Z or 2020-07-16T23:00:00+01:00. | [default to null]
**StatusLocationAddress** | [***Address**](Address.md) |  | [default to null]
**ShipmentSchedule** | [***ShipmentSchedule**](ShipmentSchedule.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

