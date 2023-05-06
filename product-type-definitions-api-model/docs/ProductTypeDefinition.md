# ProductTypeDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaSchema** | [***SchemaLink**](SchemaLink.md) |  | [optional] [default to null]
**Schema** | [***SchemaLink**](SchemaLink.md) |  | [default to null]
**Requirements** | **string** | Name of the requirements set represented in this product type definition. | [default to null]
**RequirementsEnforced** | **string** | Identifies if the required attributes for a requirements set are enforced by the product type definition schema. Non-enforced requirements enable structural validation of individual attributes without all of the required attributes being present (such as for partial updates). | [default to null]
**PropertyGroups** | [**map[string]PropertyGroup**](PropertyGroup.md) | Mapping of property group names to property groups. Property groups represent logical groupings of schema properties that can be used for display or informational purposes. | [default to null]
**Locale** | **string** | Locale of the display elements contained in the product type definition. | [default to null]
**MarketplaceIds** | **[]string** | Amazon marketplace identifiers for which the product type definition is applicable. | [default to null]
**ProductType** | **string** | The name of the Amazon product type that this product type definition applies to. | [default to null]
**ProductTypeVersion** | [***ProductTypeVersion**](ProductTypeVersion.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

