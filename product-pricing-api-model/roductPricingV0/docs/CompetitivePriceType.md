# CompetitivePriceType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompetitivePriceId** | **string** | The pricing model for each price that is returned.  Possible values:  * 1 - New Buy Box Price. * 2 - Used Buy Box Price. | [default to null]
**Price** | [***PriceType**](PriceType.md) |  | [default to null]
**Condition** | **string** | Indicates the condition of the item whose pricing information is returned. Possible values are: New, Used, Collectible, Refurbished, or Club. | [optional] [default to null]
**Subcondition** | **string** | Indicates the subcondition of the item whose pricing information is returned. Possible values are: New, Mint, Very Good, Good, Acceptable, Poor, Club, OEM, Warranty, Refurbished Warranty, Refurbished, Open Box, or Other. | [optional] [default to null]
**OfferType** | [***OfferCustomerType**](OfferCustomerType.md) |  | [optional] [default to null]
**QuantityTier** | **int32** | Indicates at what quantity this price becomes active. | [optional] [default to null]
**QuantityDiscountType** | [***QuantityDiscountType**](QuantityDiscountType.md) |  | [optional] [default to null]
**SellerId** | **string** | The seller identifier for the offer. | [optional] [default to null]
**BelongsToRequester** | **bool** |  Indicates whether or not the pricing information is for an offer listing that belongs to the requester. The requester is the seller associated with the SellerId that was submitted with the request. Possible values are: true and false. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

