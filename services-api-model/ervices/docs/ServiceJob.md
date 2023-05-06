# ServiceJob

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | [**time.Time**](time.Time.md) | The date and time of the creation of the job in ISO 8601 format. | [optional] [default to null]
**ServiceJobId** | **string** |  | [optional] [default to null]
**ServiceJobStatus** | **string** | The status of the service job. | [optional] [default to null]
**ScopeOfWork** | [***ScopeOfWork**](ScopeOfWork.md) |  | [optional] [default to null]
**Seller** | [***Seller**](Seller.md) |  | [optional] [default to null]
**ServiceJobProvider** | [***ServiceJobProvider**](ServiceJobProvider.md) |  | [optional] [default to null]
**PreferredAppointmentTimes** | [**[]AppointmentTime**](AppointmentTime.md) | A list of appointment windows preferred by the buyer. Included only if the buyer selected appointment windows when creating the order. | [optional] [default to null]
**Appointments** | [**[]Appointment**](Appointment.md) | A list of appointments. | [optional] [default to null]
**ServiceOrderId** | **string** |  | [optional] [default to null]
**MarketplaceId** | **string** | The marketplace identifier. | [optional] [default to null]
**StoreId** | **string** | The Amazon-defined identifier for the region scope. | [optional] [default to null]
**Buyer** | [***Buyer**](Buyer.md) |  | [optional] [default to null]
**AssociatedItems** | [**[]AssociatedItem**](AssociatedItem.md) | A list of items associated with the service job. | [optional] [default to null]
**ServiceLocation** | [***ServiceLocation**](ServiceLocation.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

