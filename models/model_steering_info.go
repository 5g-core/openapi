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


type SteeringInfo struct {


	
	// simple type
 
	

	PlmnId PlmnId `json:"plmnId" yaml:"plmnId" bson:"plmnId" mapstructure:"PlmnId"`
 
	

	AccessTechList *[]AccessTech `json:"accessTechList,omitempty" yaml:"accessTechList" bson:"accessTechList" mapstructure:"AccessTechList"`
}
