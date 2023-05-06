# GetFeatureSkuResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceId** | **string** | The requested marketplace. | [default to null]
**FeatureName** | **string** | The name of the feature. | [default to null]
**IsEligible** | **bool** | When true, the seller SKU is eligible for the requested feature. | [default to null]
**IneligibleReasons** | **[]string** | A list of one or more reasons that the seller SKU is ineligibile for the feature.  Possible values: * MERCHANT_NOT_ENROLLED - The merchant isn&#x27;t enrolled for the feature. * SKU_NOT_ELIGIBLE - The SKU doesn&#x27;t reside in a warehouse that supports the feature. * INVALID_SKU - There is an issue with the SKU provided. | [optional] [default to null]
**SkuInfo** | [***FeatureSku**](FeatureSku.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

