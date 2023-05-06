/*
 * Selling Partner API for Merchant Fulfillment
 *
 * The Selling Partner API for Merchant Fulfillment helps you build applications that let sellers purchase shipping for non-Prime and Prime orders using Amazon’s Buy Shipping Services.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// PredefinedPackageDimensions : An enumeration of predefined parcel tokens. If you specify a PredefinedPackageDimensions token, you are not obligated to use a branded package from a carrier. For example, if you specify the FedEx_Box_10kg token, you do not have to use that particular package from FedEx. You are only obligated to use a box that matches the dimensions specified by the token.  Note: Please note that carriers can have restrictions on the type of package allowed for certain ship methods. Check the carrier website for all details. Example: Flat rate pricing is available when materials are sent by USPS in a USPS-produced Flat Rate Envelope or Box.
type PredefinedPackageDimensions string

// List of PredefinedPackageDimensions
const (
	FED_EX_BOX_10KG_PredefinedPackageDimensions PredefinedPackageDimensions = "FedEx_Box_10kg"
	FED_EX_BOX_25KG_PredefinedPackageDimensions PredefinedPackageDimensions = "FedEx_Box_25kg"
	FED_EX_BOX_EXTRA_LARGE_1_PredefinedPackageDimensions PredefinedPackageDimensions = "FedEx_Box_Extra_Large_1"
	FED_EX_BOX_EXTRA_LARGE_2_PredefinedPackageDimensions PredefinedPackageDimensions = "FedEx_Box_Extra_Large_2"
	FED_EX_BOX_LARGE_1_PredefinedPackageDimensions PredefinedPackageDimensions = "FedEx_Box_Large_1"
	FED_EX_BOX_LARGE_2_PredefinedPackageDimensions PredefinedPackageDimensions = "FedEx_Box_Large_2"
	FED_EX_BOX_MEDIUM_1_PredefinedPackageDimensions PredefinedPackageDimensions = "FedEx_Box_Medium_1"
	FED_EX_BOX_MEDIUM_2_PredefinedPackageDimensions PredefinedPackageDimensions = "FedEx_Box_Medium_2"
	FED_EX_BOX_SMALL_1_PredefinedPackageDimensions PredefinedPackageDimensions = "FedEx_Box_Small_1"
	FED_EX_BOX_SMALL_2_PredefinedPackageDimensions PredefinedPackageDimensions = "FedEx_Box_Small_2"
	FED_EX_ENVELOPE_PredefinedPackageDimensions PredefinedPackageDimensions = "FedEx_Envelope"
	FED_EX_PADDED_PAK_PredefinedPackageDimensions PredefinedPackageDimensions = "FedEx_Padded_Pak"
	FED_EX_PAK_1_PredefinedPackageDimensions PredefinedPackageDimensions = "FedEx_Pak_1"
	FED_EX_PAK_2_PredefinedPackageDimensions PredefinedPackageDimensions = "FedEx_Pak_2"
	FED_EX_TUBE_PredefinedPackageDimensions PredefinedPackageDimensions = "FedEx_Tube"
	FED_EX_XL_PAK_PredefinedPackageDimensions PredefinedPackageDimensions = "FedEx_XL_Pak"
	UPS_BOX_10KG_PredefinedPackageDimensions PredefinedPackageDimensions = "UPS_Box_10kg"
	UPS_BOX_25KG_PredefinedPackageDimensions PredefinedPackageDimensions = "UPS_Box_25kg"
	UPS_EXPRESS_BOX_PredefinedPackageDimensions PredefinedPackageDimensions = "UPS_Express_Box"
	UPS_EXPRESS_BOX_LARGE_PredefinedPackageDimensions PredefinedPackageDimensions = "UPS_Express_Box_Large"
	UPS_EXPRESS_BOX_MEDIUM_PredefinedPackageDimensions PredefinedPackageDimensions = "UPS_Express_Box_Medium"
	UPS_EXPRESS_BOX_SMALL_PredefinedPackageDimensions PredefinedPackageDimensions = "UPS_Express_Box_Small"
	UPS_EXPRESS_ENVELOPE_PredefinedPackageDimensions PredefinedPackageDimensions = "UPS_Express_Envelope"
	UPS_EXPRESS_HARD_PAK_PredefinedPackageDimensions PredefinedPackageDimensions = "UPS_Express_Hard_Pak"
	UPS_EXPRESS_LEGAL_ENVELOPE_PredefinedPackageDimensions PredefinedPackageDimensions = "UPS_Express_Legal_Envelope"
	UPS_EXPRESS_PAK_PredefinedPackageDimensions PredefinedPackageDimensions = "UPS_Express_Pak"
	UPS_EXPRESS_TUBE_PredefinedPackageDimensions PredefinedPackageDimensions = "UPS_Express_Tube"
	UPS_LABORATORY_PAK_PredefinedPackageDimensions PredefinedPackageDimensions = "UPS_Laboratory_Pak"
	UPS_PAD_PAK_PredefinedPackageDimensions PredefinedPackageDimensions = "UPS_Pad_Pak"
	UPS_PALLET_PredefinedPackageDimensions PredefinedPackageDimensions = "UPS_Pallet"
	USPS_CARD_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_Card"
	USPS_FLAT_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_Flat"
	USPS_FLAT_RATE_CARDBOARD_ENVELOPE_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_FlatRateCardboardEnvelope"
	USPS_FLAT_RATE_ENVELOPE_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_FlatRateEnvelope"
	USPS_FLAT_RATE_GIFT_CARD_ENVELOPE_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_FlatRateGiftCardEnvelope"
	USPS_FLAT_RATE_LEGAL_ENVELOPE_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_FlatRateLegalEnvelope"
	USPS_FLAT_RATE_PADDED_ENVELOPE_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_FlatRatePaddedEnvelope"
	USPS_FLAT_RATE_WINDOW_ENVELOPE_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_FlatRateWindowEnvelope"
	USPS_LARGE_FLAT_RATE_BOARD_GAME_BOX_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_LargeFlatRateBoardGameBox"
	USPS_LARGE_FLAT_RATE_BOX_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_LargeFlatRateBox"
	USPS_LETTER_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_Letter"
	USPS_MEDIUM_FLAT_RATE_BOX1_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_MediumFlatRateBox1"
	USPS_MEDIUM_FLAT_RATE_BOX2_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_MediumFlatRateBox2"
	USPS_REGIONAL_RATE_BOX_A1_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_RegionalRateBoxA1"
	USPS_REGIONAL_RATE_BOX_A2_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_RegionalRateBoxA2"
	USPS_REGIONAL_RATE_BOX_B1_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_RegionalRateBoxB1"
	USPS_REGIONAL_RATE_BOX_B2_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_RegionalRateBoxB2"
	USPS_REGIONAL_RATE_BOX_C_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_RegionalRateBoxC"
	USPS_SMALL_FLAT_RATE_BOX_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_SmallFlatRateBox"
	USPS_SMALL_FLAT_RATE_ENVELOPE_PredefinedPackageDimensions PredefinedPackageDimensions = "USPS_SmallFlatRateEnvelope"
)
