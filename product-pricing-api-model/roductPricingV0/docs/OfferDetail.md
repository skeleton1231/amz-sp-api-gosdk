# OfferDetail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MyOffer** | **bool** | When true, this is the seller&#x27;s offer. | [optional] [default to null]
**OfferType** | [***OfferCustomerType**](OfferCustomerType.md) |  | [optional] [default to null]
**SubCondition** | **string** | The subcondition of the item. Subcondition values: New, Mint, Very Good, Good, Acceptable, Poor, Club, OEM, Warranty, Refurbished Warranty, Refurbished, Open Box, or Other. | [default to null]
**SellerId** | **string** | The seller identifier for the offer. | [optional] [default to null]
**ConditionNotes** | **string** | Information about the condition of the item. | [optional] [default to null]
**SellerFeedbackRating** | [***SellerFeedbackType**](SellerFeedbackType.md) |  | [optional] [default to null]
**ShippingTime** | [***DetailedShippingTimeType**](DetailedShippingTimeType.md) |  | [default to null]
**ListingPrice** | [***MoneyType**](MoneyType.md) |  | [default to null]
**QuantityDiscountPrices** | [**[]QuantityDiscountPriceType**](QuantityDiscountPriceType.md) |  | [optional] [default to null]
**Points** | [***Points**](Points.md) |  | [optional] [default to null]
**Shipping** | [***MoneyType**](MoneyType.md) |  | [default to null]
**ShipsFrom** | [***ShipsFromType**](ShipsFromType.md) |  | [optional] [default to null]
**IsFulfilledByAmazon** | **bool** | When true, the offer is fulfilled by Amazon. | [default to null]
**PrimeInformation** | [***PrimeInformationType**](PrimeInformationType.md) |  | [optional] [default to null]
**IsBuyBoxWinner** | **bool** | When true, the offer is currently in the Buy Box. There can be up to two Buy Box winners at any time per ASIN, one that is eligible for Prime and one that is not eligible for Prime. | [optional] [default to null]
**IsFeaturedMerchant** | **bool** | When true, the seller of the item is eligible to win the Buy Box. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

