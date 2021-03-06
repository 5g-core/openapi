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


type SorData struct {


	
	// simple type
 
	

	ProvisioningTime time.Time `json:"provisioningTime" yaml:"provisioningTime" bson:"provisioningTime" mapstructure:"ProvisioningTime"`
 
	

	UeUpdateStatus UeUpdateStatus `json:"ueUpdateStatus" yaml:"ueUpdateStatus" bson:"ueUpdateStatus" mapstructure:"UeUpdateStatus"`
 
	

	SorXmacIue *string `json:"sorXmacIue,omitempty" yaml:"sorXmacIue" bson:"sorXmacIue" mapstructure:"SorXmacIue"`
 
	

	SorMacIue *string `json:"sorMacIue,omitempty" yaml:"sorMacIue" bson:"sorMacIue" mapstructure:"SorMacIue"`
}
