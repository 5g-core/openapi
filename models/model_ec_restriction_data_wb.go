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


type EcRestrictionDataWb struct {


	
	// simple type
 
	

	EcModeARestricted *bool `json:"ecModeARestricted,omitempty" yaml:"ecModeARestricted" bson:"ecModeARestricted" mapstructure:"EcModeARestricted"`
 
	

	EcModeBRestricted *bool `json:"ecModeBRestricted,omitempty" yaml:"ecModeBRestricted" bson:"ecModeBRestricted" mapstructure:"EcModeBRestricted"`
}
