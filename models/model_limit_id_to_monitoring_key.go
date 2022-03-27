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


// LimitIdToMonitoringKey - Contains the limit identifier and the corresponding monitoring key for a given S-NSSAI and DNN.
type LimitIdToMonitoringKey struct {


	
	// simple type
 
	

	LimitId string `json:"limitId" yaml:"limitId" bson:"limitId" mapstructure:"LimitId"`
 
	

	Monkey *[]string `json:"monkey,omitempty" yaml:"monkey" bson:"monkey" mapstructure:"Monkey"`
}