# PackageTrackingDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageNumber** | **int32** | The package identifier. | [default to null]
**TrackingNumber** | **string** | The tracking number for the package. | [optional] [default to null]
**CustomerTrackingLink** | **string** | Link on swiship.com that allows customers to track the package. | [optional] [default to null]
**CarrierCode** | **string** | The name of the carrier. | [optional] [default to null]
**CarrierPhoneNumber** | **string** | The phone number of the carrier. | [optional] [default to null]
**CarrierURL** | **string** | The URL of the carrier&#x27;s website. | [optional] [default to null]
**ShipDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**EstimatedArrivalDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**ShipToAddress** | [***TrackingAddress**](TrackingAddress.md) |  | [optional] [default to null]
**CurrentStatus** | [***CurrentStatus**](CurrentStatus.md) |  | [optional] [default to null]
**CurrentStatusDescription** | **string** | Description corresponding to the CurrentStatus value. | [optional] [default to null]
**SignedForBy** | **string** | The name of the person who signed for the package. | [optional] [default to null]
**AdditionalLocationInfo** | [***AdditionalLocationInfo**](AdditionalLocationInfo.md) |  | [optional] [default to null]
**TrackingEvents** | [***[]TrackingEvent**](array.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

