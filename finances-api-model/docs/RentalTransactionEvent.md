# RentalTransactionEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmazonOrderId** | **string** | An Amazon-defined identifier for an order. | [optional] [default to null]
**RentalEventType** | **string** | The type of rental event.  Possible values:  * RentalCustomerPayment-Buyout - Transaction type that represents when the customer wants to buy out a rented item.  * RentalCustomerPayment-Extension - Transaction type that represents when the customer wants to extend the rental period.  * RentalCustomerRefund-Buyout - Transaction type that represents when the customer requests a refund for the buyout of the rented item.  * RentalCustomerRefund-Extension - Transaction type that represents when the customer requests a refund over the extension on the rented item.  * RentalHandlingFee - Transaction type that represents the fee that Amazon charges sellers who rent through Amazon.  * RentalChargeFailureReimbursement - Transaction type that represents when Amazon sends money to the seller to compensate for a failed charge.  * RentalLostItemReimbursement - Transaction type that represents when Amazon sends money to the seller to compensate for a lost item. | [optional] [default to null]
**ExtensionLength** | **int32** | The number of days that the buyer extended an already rented item. This value is only returned for RentalCustomerPayment-Extension and RentalCustomerRefund-Extension events. | [optional] [default to null]
**PostedDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**RentalChargeList** | [***[]ChargeComponent**](array.md) |  | [optional] [default to null]
**RentalFeeList** | [***[]FeeComponent**](array.md) |  | [optional] [default to null]
**MarketplaceName** | **string** | The name of the marketplace. | [optional] [default to null]
**RentalInitialValue** | [***Currency**](Currency.md) |  | [optional] [default to null]
**RentalReimbursement** | [***Currency**](Currency.md) |  | [optional] [default to null]
**RentalTaxWithheldList** | [***[]TaxWithheldComponent**](array.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

