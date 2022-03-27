/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data). The API version is defined in 3GPP TS 29.504. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type LocationAccuracyAnyOf string

// List of LocationAccuracyAnyOf 
const (
	LOCATIONACCURACYANYOF_CELL_LEVEL LocationAccuracyAnyOf = "CELL_LEVEL"
	LOCATIONACCURACYANYOF_TA_LEVEL LocationAccuracyAnyOf = "TA_LEVEL"
	LOCATIONACCURACYANYOF_N3_IWF_LEVEL LocationAccuracyAnyOf = "N3IWF_LEVEL"
	LOCATIONACCURACYANYOF_UE_IP LocationAccuracyAnyOf = "UE_IP"
	LOCATIONACCURACYANYOF_UE_PORT LocationAccuracyAnyOf = "UE_PORT"
)