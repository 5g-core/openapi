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


type UserLocation struct {


	
	// simple type
 
	

	EutraLocation *EutraLocation `json:"eutraLocation,omitempty" yaml:"eutraLocation" bson:"eutraLocation" mapstructure:"EutraLocation"`
 
	

	NrLocation *NrLocation `json:"nrLocation,omitempty" yaml:"nrLocation" bson:"nrLocation" mapstructure:"NrLocation"`
 
	

	N3gaLocation *N3gaLocation `json:"n3gaLocation,omitempty" yaml:"n3gaLocation" bson:"n3gaLocation" mapstructure:"N3gaLocation"`
 
	

	UtraLocation *UtraLocation `json:"utraLocation,omitempty" yaml:"utraLocation" bson:"utraLocation" mapstructure:"UtraLocation"`
 
	

	GeraLocation *GeraLocation `json:"geraLocation,omitempty" yaml:"geraLocation" bson:"geraLocation" mapstructure:"GeraLocation"`
}
