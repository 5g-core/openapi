/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data). The API version is defined in 3GPP TS 29.504. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type TrafficProfileAnyOf string

// List of TrafficProfileAnyOf 
const (
	TRAFFICPROFILEANYOF_SINGLE_TRANS_UL TrafficProfileAnyOf = "SINGLE_TRANS_UL"
	TRAFFICPROFILEANYOF_SINGLE_TRANS_DL TrafficProfileAnyOf = "SINGLE_TRANS_DL"
	TRAFFICPROFILEANYOF_DUAL_TRANS_UL_FIRST TrafficProfileAnyOf = "DUAL_TRANS_UL_FIRST"
	TRAFFICPROFILEANYOF_DUAL_TRANS_DL_FIRST TrafficProfileAnyOf = "DUAL_TRANS_DL_FIRST"
	TRAFFICPROFILEANYOF_MULTI_TRANS TrafficProfileAnyOf = "MULTI_TRANS"
)
