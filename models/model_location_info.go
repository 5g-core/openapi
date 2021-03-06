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


type LocationInfo struct {


	
	// simple type
 
	

	Supi *string `json:"supi,omitempty" yaml:"supi" bson:"supi" mapstructure:"Supi"`
 
	

	Gpsi *string `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi" mapstructure:"Gpsi"`
 
	

	RegistrationLocationInfoList []RegistrationLocationInfo `json:"registrationLocationInfoList" yaml:"registrationLocationInfoList" bson:"registrationLocationInfoList" mapstructure:"RegistrationLocationInfoList"`
 
	

	SupportedFeatures *string `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures" mapstructure:"SupportedFeatures"`
}
