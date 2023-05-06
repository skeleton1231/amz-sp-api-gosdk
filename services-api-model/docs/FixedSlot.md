# FixedSlot

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDateTime** | [**time.Time**](time.Time.md) | Start date time of slot in ISO 8601 format with precision of seconds. | [optional] [default to null]
**ScheduledCapacity** | **int32** | Scheduled capacity corresponding to the slot. This capacity represents the originally allocated capacity as per resource schedule. | [optional] [default to null]
**AvailableCapacity** | **int32** | Available capacity corresponding to the slot. This capacity represents the capacity available for allocation to reservations. | [optional] [default to null]
**EncumberedCapacity** | **int32** | Encumbered capacity corresponding to the slot. This capacity represents the capacity allocated for Amazon Jobs/Appointments/Orders. | [optional] [default to null]
**ReservedCapacity** | **int32** | Reserved capacity corresponding to the slot. This capacity represents the capacity made unavailable due to events like Breaks/Leaves/Lunch. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

