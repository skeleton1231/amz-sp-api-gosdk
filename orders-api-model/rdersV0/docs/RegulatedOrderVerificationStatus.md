# RegulatedOrderVerificationStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [***VerificationStatus**](VerificationStatus.md) |  | [default to null]
**RequiresMerchantAction** | **bool** | When true, the regulated information provided in the order requires a review by the merchant. | [default to null]
**ValidRejectionReasons** | [**[]RejectionReason**](RejectionReason.md) | A list of valid rejection reasons that may be used to reject the order&#x27;s regulated information. | [default to null]
**RejectionReason** | [***RejectionReason**](RejectionReason.md) |  | [optional] [default to null]
**ReviewDate** | **string** | The date the order was reviewed. In ISO 8601 date time format. | [optional] [default to null]
**ExternalReviewerId** | **string** | The identifier for the order&#x27;s regulated information reviewer. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

