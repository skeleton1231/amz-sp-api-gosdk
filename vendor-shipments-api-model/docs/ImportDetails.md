# ImportDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodOfPayment** | **string** | This is used for import purchase orders only. If the recipient requests, this field will contain the shipment method of payment. | [optional] [default to null]
**SealNumber** | **string** | The container&#x27;s seal number. | [optional] [default to null]
**Route** | [***Route**](Route.md) |  | [optional] [default to null]
**ImportContainers** | **string** | Types and numbers of container(s) for import purchase orders. Can be a comma-separated list if shipment has multiple containers. | [optional] [default to null]
**BillableWeight** | [***Weight**](Weight.md) |  | [optional] [default to null]
**EstimatedShipByDate** | [**time.Time**](time.Time.md) | Date on which the shipment is expected to be shipped. This value should not be in the past and not more than 60 days out in the future. | [optional] [default to null]
**HandlingInstructions** | **string** | Identification of the instructions on how specified item/carton/pallet should be handled. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

