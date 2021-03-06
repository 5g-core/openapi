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


// UePolicySetPatch - Contains the UE policy set for a given subscriber.
type UePolicySetPatch struct {


	
	// simple type
 
	

	UePolicySections *map[string]UePolicySection `json:"uePolicySections,omitempty" yaml:"uePolicySections" bson:"uePolicySections" mapstructure:"UePolicySections"`
 
	

	Upsis *[]string `json:"upsis,omitempty" yaml:"upsis" bson:"upsis" mapstructure:"Upsis"`
 
	

	AndspInd *bool `json:"andspInd,omitempty" yaml:"andspInd" bson:"andspInd" mapstructure:"AndspInd"`
 
	

	Pei *string `json:"pei,omitempty" yaml:"pei" bson:"pei" mapstructure:"Pei"`
 
	

	OsIds *[]string `json:"osIds,omitempty" yaml:"osIds" bson:"osIds" mapstructure:"OsIds"`
}
