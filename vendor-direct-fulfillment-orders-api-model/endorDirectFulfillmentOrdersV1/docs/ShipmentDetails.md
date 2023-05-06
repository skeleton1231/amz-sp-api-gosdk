# ShipmentDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPriorityShipment** | **bool** | When true, this is a priority shipment. | [default to null]
**IsScheduledDeliveryShipment** | **bool** | When true, this order is part of a scheduled delivery program. | [optional] [default to null]
**IsPslipRequired** | **bool** | When true, a packing slip is required to be sent to the customer. | [default to null]
**IsGift** | **bool** | When true, the order contain a gift. Include the gift message and gift wrap information. | [optional] [default to null]
**ShipMethod** | **string** | Ship method to be used for shipping the order. Amazon defines ship method codes indicating the shipping carrier and shipment service level. To see the full list of ship methods in use, including both the code and the friendly name, search the &#x27;Help&#x27; section on Vendor Central for &#x27;ship methods&#x27;. | [default to null]
**ShipmentDates** | [***ShipmentDates**](ShipmentDates.md) |  | [default to null]
**MessageToCustomer** | **string** | Message to customer for order status. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

