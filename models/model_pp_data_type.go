/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (

)


type PpDataType string
// enum type
// List of PpDataType 
const (
	PPDATATYPE_COMMUNICATION_CHARACTERISTICS PpDataType = "COMMUNICATION_CHARACTERISTICS"
	PPDATATYPE_EXPECTED_UE_BEHAVIOUR PpDataType = "EXPECTED_UE_BEHAVIOUR"
	PPDATATYPE_EC_RESTRICTION PpDataType = "EC_RESTRICTION"
	PPDATATYPE_ACS_INFO PpDataType = "ACS_INFO"
	PPDATATYPE_TRACE PpDataType = "TRACE"
	PPDATATYPE_STN_SR PpDataType = "STN_SR"
	PPDATATYPE_LCS_PRIVACY PpDataType = "LCS_PRIVACY"
	PPDATATYPE_SOR_INFO PpDataType = "SOR_INFO"
)
