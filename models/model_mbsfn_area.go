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


type MbsfnArea struct {


	
	// simple type
 
	

	MbsfnAreaId *int32 `json:"mbsfnAreaId,omitempty" yaml:"mbsfnAreaId" bson:"mbsfnAreaId" mapstructure:"MbsfnAreaId"`
 
	

	CarrierFrequency *int32 `json:"carrierFrequency,omitempty" yaml:"carrierFrequency" bson:"carrierFrequency" mapstructure:"CarrierFrequency"`
}
