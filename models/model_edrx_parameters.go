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


type EdrxParameters struct {


	
	// simple type
 
	

	RatType RatType `json:"ratType" yaml:"ratType" bson:"ratType" mapstructure:"RatType"`
 
	

	EdrxValue string `json:"edrxValue" yaml:"edrxValue" bson:"edrxValue" mapstructure:"EdrxValue"`
}
