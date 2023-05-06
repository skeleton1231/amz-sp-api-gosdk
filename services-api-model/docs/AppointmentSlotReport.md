# AppointmentSlotReport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchedulingType** | **string** | Defines the type of slots. | [optional] [default to null]
**StartTime** | [**time.Time**](time.Time.md) | Start Time from which the appointment slots are generated in ISO 8601 format. | [optional] [default to null]
**EndTime** | [**time.Time**](time.Time.md) | End Time up to which the appointment slots are generated in ISO 8601 format. | [optional] [default to null]
**AppointmentSlots** | [**[]AppointmentSlot**](AppointmentSlot.md) | A list of time windows along with associated capacity in which the service can be performed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

