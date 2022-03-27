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


type PointUncertaintyCircleAllOf struct {


	
	// simple type
 
	

	Point GeographicalCoordinates `json:"point" yaml:"point" bson:"point" mapstructure:"Point"`
 
	

	Uncertainty float32 `json:"uncertainty" yaml:"uncertainty" bson:"uncertainty" mapstructure:"Uncertainty"`
}