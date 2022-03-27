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


type LocationReportingConfiguration struct {


	
	// simple type
 
	

	CurrentLocation bool `json:"currentLocation" yaml:"currentLocation" bson:"currentLocation" mapstructure:"CurrentLocation"`
 
	

	OneTime *bool `json:"oneTime,omitempty" yaml:"oneTime" bson:"oneTime" mapstructure:"OneTime"`
 
	

	Accuracy *LocationAccuracy `json:"accuracy,omitempty" yaml:"accuracy" bson:"accuracy" mapstructure:"Accuracy"`
 
	

	N3gppAccuracy *LocationAccuracy `json:"n3gppAccuracy,omitempty" yaml:"n3gppAccuracy" bson:"n3gppAccuracy" mapstructure:"N3gppAccuracy"`
}
