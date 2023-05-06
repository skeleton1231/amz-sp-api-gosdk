# PurchaseOrders

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PurchaseOrderNumber** | **string** | Purchase order numbers involved in this shipment, list all the PO that are involved as part of this shipment. | [optional] [default to null]
**PurchaseOrderDate** | [**time.Time**](time.Time.md) | Purchase order numbers involved in this shipment, list all the PO that are involved as part of this shipment. | [optional] [default to null]
**ShipWindow** | **string** | Date range in which shipment is expected for these purchase orders. | [optional] [default to null]
**Items** | [**[]PurchaseOrderItems**](PurchaseOrderItems.md) | A list of the items that are associated to the PO in this transport and their associated details. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

