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


type PointUncertaintyCircle struct {


	
	// simple type
	// parent
 
	
	// not map
	// not array
	//Item	GadShape
	Item *GadShape `json:"Item,omitempty" yaml:"item" bson:"item" mapstructure:"Item"`
 
	

	Point GeographicalCoordinates `json:"point" yaml:"point" bson:"point" mapstructure:"Point"`
 
	

	Uncertainty float32 `json:"uncertainty" yaml:"uncertainty" bson:"uncertainty" mapstructure:"Uncertainty"`
}
