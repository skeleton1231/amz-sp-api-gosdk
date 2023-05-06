# ItemOffersRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceId** | **string** |  | [default to null]
**ItemCondition** | [***ItemCondition**](ItemCondition.md) |  | [default to null]
**CustomerType** | [***CustomerType**](CustomerType.md) |  | [optional] [default to null]
**Uri** | **string** | The resource path of the operation you are calling in batch without any query parameters.  If you are calling &#x60;getItemOffersBatch&#x60;, supply the path of &#x60;getItemOffers&#x60;.  **Example:** &#x60;/products/pricing/v0/items/B000P6Q7MY/offers&#x60;  If you are calling &#x60;getListingOffersBatch&#x60;, supply the path of &#x60;getListingOffers&#x60;.  **Example:** &#x60;/products/pricing/v0/listings/B000P6Q7MY/offers&#x60; | [default to null]
**Method** | [***HttpMethod**](HttpMethod.md) |  | [default to null]
**Headers** | [***map[string]string**](map.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

