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


type AccessRightStatus string
// enum type
// List of AccessRightStatus 
const (
	ACCESSRIGHTSTATUS_FULLY_ALLOWED AccessRightStatus = "FULLY_ALLOWED"
	ACCESSRIGHTSTATUS_PREVIEW_ALLOWED AccessRightStatus = "PREVIEW_ALLOWED"
	ACCESSRIGHTSTATUS_NO_ALLOWED AccessRightStatus = "NO_ALLOWED"
)
