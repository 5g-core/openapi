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


type ChargingInformation struct {


	
	// simple type
 
	

	PrimaryChfAddress string `json:"primaryChfAddress" yaml:"primaryChfAddress" bson:"primaryChfAddress" mapstructure:"PrimaryChfAddress"`
 
	

	SecondaryChfAddress string `json:"secondaryChfAddress" yaml:"secondaryChfAddress" bson:"secondaryChfAddress" mapstructure:"SecondaryChfAddress"`
 
	

	PrimaryChfSetId *string `json:"primaryChfSetId,omitempty" yaml:"primaryChfSetId" bson:"primaryChfSetId" mapstructure:"PrimaryChfSetId"`
 
	

	PrimaryChfInstanceId *string `json:"primaryChfInstanceId,omitempty" yaml:"primaryChfInstanceId" bson:"primaryChfInstanceId" mapstructure:"PrimaryChfInstanceId"`
 
	

	SecondaryChfSetId *string `json:"secondaryChfSetId,omitempty" yaml:"secondaryChfSetId" bson:"secondaryChfSetId" mapstructure:"SecondaryChfSetId"`
 
	

	SecondaryChfInstanceId *string `json:"secondaryChfInstanceId,omitempty" yaml:"secondaryChfInstanceId" bson:"secondaryChfInstanceId" mapstructure:"SecondaryChfInstanceId"`
}
