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

	"time"
)


type UtraLocation struct {


	
	// simple type
 
	

	Cgi *CellGlobalId `json:"cgi,omitempty" yaml:"cgi" bson:"cgi" mapstructure:"Cgi"`
 
	

	Sai *ServiceAreaId `json:"sai,omitempty" yaml:"sai" bson:"sai" mapstructure:"Sai"`
 
	

	Lai *LocationAreaId `json:"lai,omitempty" yaml:"lai" bson:"lai" mapstructure:"Lai"`
 
	

	Rai *RoutingAreaId `json:"rai,omitempty" yaml:"rai" bson:"rai" mapstructure:"Rai"`
 
	

	AgeOfLocationInformation *int32 `json:"ageOfLocationInformation,omitempty" yaml:"ageOfLocationInformation" bson:"ageOfLocationInformation" mapstructure:"AgeOfLocationInformation"`
 
	

	UeLocationTimestamp *time.Time `json:"ueLocationTimestamp,omitempty" yaml:"ueLocationTimestamp" bson:"ueLocationTimestamp" mapstructure:"UeLocationTimestamp"`
 
	

	GeographicalInformation *string `json:"geographicalInformation,omitempty" yaml:"geographicalInformation" bson:"geographicalInformation" mapstructure:"GeographicalInformation"`
 
	

	GeodeticInformation *string `json:"geodeticInformation,omitempty" yaml:"geodeticInformation" bson:"geodeticInformation" mapstructure:"GeodeticInformation"`
}