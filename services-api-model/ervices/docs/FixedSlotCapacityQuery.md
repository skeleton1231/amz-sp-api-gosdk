# FixedSlotCapacityQuery

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityTypes** | [**[]CapacityType**](CapacityType.md) | An array of capacity types which are being requested. Default value is &#x60;[SCHEDULED_CAPACITY]&#x60;. | [optional] [default to null]
**SlotDuration** | **float64** | Size in which slots are being requested. This value should be a multiple of 5 and fall in the range: 5 &lt;&#x3D; &#x60;slotDuration&#x60; &lt;&#x3D; 360. | [optional] [default to null]
**StartDateTime** | [**time.Time**](time.Time.md) | Start date time from which the capacity slots are being requested in ISO 8601 format. | [default to null]
**EndDateTime** | [**time.Time**](time.Time.md) | End date time up to which the capacity slots are being requested in ISO 8601 format. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

