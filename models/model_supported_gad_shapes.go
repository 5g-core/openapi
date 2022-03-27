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


type SupportedGadShapes string
// enum type
// List of SupportedGadShapes 
const (
	SUPPORTEDGADSHAPES_POINT SupportedGadShapes = "POINT"
	SUPPORTEDGADSHAPES_POINT_UNCERTAINTY_CIRCLE SupportedGadShapes = "POINT_UNCERTAINTY_CIRCLE"
	SUPPORTEDGADSHAPES_POINT_UNCERTAINTY_ELLIPSE SupportedGadShapes = "POINT_UNCERTAINTY_ELLIPSE"
	SUPPORTEDGADSHAPES_POLYGON SupportedGadShapes = "POLYGON"
	SUPPORTEDGADSHAPES_POINT_ALTITUDE SupportedGadShapes = "POINT_ALTITUDE"
	SUPPORTEDGADSHAPES_POINT_ALTITUDE_UNCERTAINTY SupportedGadShapes = "POINT_ALTITUDE_UNCERTAINTY"
	SUPPORTEDGADSHAPES_ELLIPSOID_ARC SupportedGadShapes = "ELLIPSOID_ARC"
)
