# FeePreview

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asin** | **string** | The Amazon Standard Identification Number (ASIN) value used to identify the item. | [optional] [default to null]
**Price** | [***MoneyType**](MoneyType.md) |  | [optional] [default to null]
**FeeBreakdown** | [**[]FeeLineItem**](FeeLineItem.md) | A list of the Small and Light fees for the item. | [optional] [default to null]
**TotalFees** | [***MoneyType**](MoneyType.md) |  | [optional] [default to null]
**Errors** | [**[]ModelError**](Error.md) | One or more unexpected errors occurred during the getSmallAndLightFeePreview operation. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

