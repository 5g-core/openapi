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


// UsageMonDataScope - Contains a SNSSAI and DNN combinations to which the UsageMonData instance belongs to.
type UsageMonDataScope struct {


	
	// simple type
 
	

	Snssai Snssai `json:"snssai" yaml:"snssai" bson:"snssai" mapstructure:"Snssai"`
 
	

	Dnn *[]string `json:"dnn,omitempty" yaml:"dnn" bson:"dnn" mapstructure:"Dnn"`
}