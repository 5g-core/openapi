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


type NrV2xAuth struct {


	
	// simple type
 
	

	VehicleUeAuth *UeAuth `json:"vehicleUeAuth,omitempty" yaml:"vehicleUeAuth" bson:"vehicleUeAuth" mapstructure:"VehicleUeAuth"`
 
	

	PedestrianUeAuth *UeAuth `json:"pedestrianUeAuth,omitempty" yaml:"pedestrianUeAuth" bson:"pedestrianUeAuth" mapstructure:"PedestrianUeAuth"`
}
