# ImportDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodOfPayment** | **string** | If the recipient requests, contains the shipment method of payment. This is for import PO&#x27;s only. | [optional] [default to null]
**InternationalCommercialTerms** | **string** | Incoterms (International Commercial Terms) are used to divide transaction costs and responsibilities between buyer and seller and reflect state-of-the-art transportation practices. This is for import purchase orders only.  | [optional] [default to null]
**PortOfDelivery** | **string** | The port where goods on an import purchase order must be delivered by the vendor. This should only be specified when the internationalCommercialTerms is FOB. | [optional] [default to null]
**ImportContainers** | **string** | Types and numbers of container(s) for import purchase orders. Can be a comma-separated list if the shipment has multiple containers. HC signifies a high-capacity container. Free-text field, limited to 64 characters. The format will be a comma-delimited list containing values of the type: $NUMBER_OF_CONTAINERS_OF_THIS_TYPE-$CONTAINER_TYPE. The list of values for the container type is: 40&#x27;(40-foot container), 40&#x27;HC (40-foot high-capacity container), 45&#x27;, 45&#x27;HC, 30&#x27;, 30&#x27;HC, 20&#x27;, 20&#x27;HC. | [optional] [default to null]
**ShippingInstructions** | **string** | Special instructions regarding the shipment. This field is for import purchase orders. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

