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


type CmInfo struct {


	
	// simple type
 
	

	CmState CmState `json:"cmState" yaml:"cmState" bson:"cmState" mapstructure:"CmState"`
 
	

	AccessType AccessType `json:"accessType" yaml:"accessType" bson:"accessType" mapstructure:"AccessType"`
}