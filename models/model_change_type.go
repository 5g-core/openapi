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


type ChangeType string
// enum type
// List of ChangeType 
const (
	CHANGETYPE_ADD ChangeType = "ADD"
	CHANGETYPE_MOVE ChangeType = "MOVE"
	CHANGETYPE_REMOVE ChangeType = "REMOVE"
	CHANGETYPE_REPLACE ChangeType = "REPLACE"
)
