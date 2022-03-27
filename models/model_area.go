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


type Area struct {


	
	// simple type
 
	

	Tacs *[]string `json:"tacs,omitempty" yaml:"tacs" bson:"tacs" mapstructure:"Tacs"`
 
	

	AreaCode *string `json:"areaCode,omitempty" yaml:"areaCode" bson:"areaCode" mapstructure:"AreaCode"`
}
