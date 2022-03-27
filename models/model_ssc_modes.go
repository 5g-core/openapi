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


type SscModes struct {


	
	// simple type
 
	

	DefaultSscMode SscMode `json:"defaultSscMode" yaml:"defaultSscMode" bson:"defaultSscMode" mapstructure:"DefaultSscMode"`
 
	

	AllowedSscModes *[]SscMode `json:"allowedSscModes,omitempty" yaml:"allowedSscModes" bson:"allowedSscModes" mapstructure:"AllowedSscModes"`
}
