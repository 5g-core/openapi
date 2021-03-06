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


type SorInfo struct {


	
	// simple type
 
	

	SteeringContainer *SteeringContainer `json:"steeringContainer,omitempty" yaml:"steeringContainer" bson:"steeringContainer" mapstructure:"SteeringContainer"`
 
	

	AckInd bool `json:"ackInd" yaml:"ackInd" bson:"ackInd" mapstructure:"AckInd"`
 
	

	SorMacIausf *string `json:"sorMacIausf,omitempty" yaml:"sorMacIausf" bson:"sorMacIausf" mapstructure:"SorMacIausf"`
 
	

	Countersor *string `json:"countersor,omitempty" yaml:"countersor" bson:"countersor" mapstructure:"Countersor"`
 
	

	ProvisioningTime time.Time `json:"provisioningTime" yaml:"provisioningTime" bson:"provisioningTime" mapstructure:"ProvisioningTime"`
}
