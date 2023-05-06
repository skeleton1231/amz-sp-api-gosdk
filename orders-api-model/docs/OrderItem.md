# OrderItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ASIN** | **string** | The Amazon Standard Identification Number (ASIN) of the item. | [default to null]
**SellerSKU** | **string** | The seller stock keeping unit (SKU) of the item. | [optional] [default to null]
**OrderItemId** | **string** | An Amazon-defined order item identifier. | [default to null]
**Title** | **string** | The name of the item. | [optional] [default to null]
**QuantityOrdered** | **int32** | The number of items in the order.  | [default to null]
**QuantityShipped** | **int32** | The number of items shipped. | [optional] [default to null]
**ProductInfo** | [***ProductInfoDetail**](ProductInfoDetail.md) |  | [optional] [default to null]
**PointsGranted** | [***PointsGrantedDetail**](PointsGrantedDetail.md) |  | [optional] [default to null]
**ItemPrice** | [***Money**](Money.md) |  | [optional] [default to null]
**ShippingPrice** | [***Money**](Money.md) |  | [optional] [default to null]
**ItemTax** | [***Money**](Money.md) |  | [optional] [default to null]
**ShippingTax** | [***Money**](Money.md) |  | [optional] [default to null]
**ShippingDiscount** | [***Money**](Money.md) |  | [optional] [default to null]
**ShippingDiscountTax** | [***Money**](Money.md) |  | [optional] [default to null]
**PromotionDiscount** | [***Money**](Money.md) |  | [optional] [default to null]
**PromotionDiscountTax** | [***Money**](Money.md) |  | [optional] [default to null]
**PromotionIds** | [***[]string**](array.md) |  | [optional] [default to null]
**CODFee** | [***Money**](Money.md) |  | [optional] [default to null]
**CODFeeDiscount** | [***Money**](Money.md) |  | [optional] [default to null]
**IsGift** | **bool** | When true, the item is a gift. | [optional] [default to null]
**ConditionNote** | **string** | The condition of the item as described by the seller. | [optional] [default to null]
**ConditionId** | **string** | The condition of the item.  Possible values: New, Used, Collectible, Refurbished, Preorder, Club. | [optional] [default to null]
**ConditionSubtypeId** | **string** | The subcondition of the item.  Possible values: New, Mint, Very Good, Good, Acceptable, Poor, Club, OEM, Warranty, Refurbished Warranty, Refurbished, Open Box, Any, Other. | [optional] [default to null]
**ScheduledDeliveryStartDate** | **string** | The start date of the scheduled delivery window in the time zone of the order destination. In ISO 8601 date time format. | [optional] [default to null]
**ScheduledDeliveryEndDate** | **string** | The end date of the scheduled delivery window in the time zone of the order destination. In ISO 8601 date time format. | [optional] [default to null]
**PriceDesignation** | **string** | Indicates that the selling price is a special price that is available only for Amazon Business orders. For more information about the Amazon Business Seller Program, see the [Amazon Business website](https://www.amazon.com/b2b/info/amazon-business).   Possible values: BusinessPrice - A special price that is available only for Amazon Business orders. | [optional] [default to null]
**TaxCollection** | [***TaxCollection**](TaxCollection.md) |  | [optional] [default to null]
**SerialNumberRequired** | **bool** | When true, the product type for this item has a serial number.  Returned only for Amazon Easy Ship orders. | [optional] [default to null]
**IsTransparency** | **bool** | When true, transparency codes are required. | [optional] [default to null]
**IossNumber** | **string** | The IOSS number for the marketplace. Sellers shipping to the European Union (EU) from outside of the EU must provide this IOSS number to their carrier when Amazon has collected the VAT on the sale. | [optional] [default to null]
**StoreChainStoreId** | **string** | The store chain store identifier. Linked to a specific store in a store chain. | [optional] [default to null]
**DeemedResellerCategory** | **string** | The category of deemed reseller. This applies to selling partners that are not based in the EU and is used to help them meet the VAT Deemed Reseller tax laws in the EU and UK. | [optional] [default to null]
**BuyerInfo** | [***ItemBuyerInfo**](ItemBuyerInfo.md) |  | [optional] [default to null]
**BuyerRequestedCancel** | [***BuyerRequestedCancel**](BuyerRequestedCancel.md) |  | [optional] [default to null]
**ItemApprovalContext** | [***ItemApprovalContext**](ItemApprovalContext.md) |  | [optional] [default to null]
**SerialNumbers** | **[]string** | A list of serial numbers for electronic products that are shipped to customers. Returned for FBA orders only. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

