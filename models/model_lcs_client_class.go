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


type LcsClientClass string
// enum type
// List of LcsClientClass 
const (
	LCSCLIENTCLASS_BROADCAST_SERVICE LcsClientClass = "BROADCAST_SERVICE"
	LCSCLIENTCLASS_OM_IN_HPLMN LcsClientClass = "OM_IN_HPLMN"
	LCSCLIENTCLASS_OM_IN_VPLMN LcsClientClass = "OM_IN_VPLMN"
	LCSCLIENTCLASS_ANONYMOUS_LOCATION_SERVICE LcsClientClass = "ANONYMOUS_LOCATION_SERVICE"
	LCSCLIENTCLASS_SPECIFIC_SERVICE LcsClientClass = "SPECIFIC_SERVICE"
)