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


type PlmnOperatorClass struct {


	
	// simple type
 
	

	LcsClientClass LcsClientClass `json:"lcsClientClass" yaml:"lcsClientClass" bson:"lcsClientClass" mapstructure:"LcsClientClass"`
 
	

	LcsClientIds []string `json:"lcsClientIds" yaml:"lcsClientIds" bson:"lcsClientIds" mapstructure:"LcsClientIds"`
}
