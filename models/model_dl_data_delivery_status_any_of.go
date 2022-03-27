/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data). The API version is defined in 3GPP TS 29.504. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type DlDataDeliveryStatusAnyOf string

// List of DlDataDeliveryStatusAnyOf 
const (
	DLDATADELIVERYSTATUSANYOF_BUFFERED DlDataDeliveryStatusAnyOf = "BUFFERED"
	DLDATADELIVERYSTATUSANYOF_TRANSMITTED DlDataDeliveryStatusAnyOf = "TRANSMITTED"
	DLDATADELIVERYSTATUSANYOF_DISCARDED DlDataDeliveryStatusAnyOf = "DISCARDED"
)