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


type UeId struct {


	
	// simple type
 
	

	Supi string `json:"supi" yaml:"supi" bson:"supi" mapstructure:"Supi"`
 
	

	GpsiList *[]string `json:"gpsiList,omitempty" yaml:"gpsiList" bson:"gpsiList" mapstructure:"GpsiList"`
}