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


type Lpi struct {


	
	// simple type
 
	

	LocationPrivacyInd LocationPrivacyInd `json:"locationPrivacyInd" yaml:"locationPrivacyInd" bson:"locationPrivacyInd" mapstructure:"LocationPrivacyInd"`
 
	

	ValidTimePeriod *ValidTimePeriod `json:"validTimePeriod,omitempty" yaml:"validTimePeriod" bson:"validTimePeriod" mapstructure:"ValidTimePeriod"`
}