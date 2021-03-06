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


type LcsClientGroupExternal struct {


	
	// simple type
 
	

	LcsClientGroupId *string `json:"lcsClientGroupId,omitempty" yaml:"lcsClientGroupId" bson:"lcsClientGroupId" mapstructure:"LcsClientGroupId"`
 
	

	AllowedGeographicArea *[]GeographicArea `json:"allowedGeographicArea,omitempty" yaml:"allowedGeographicArea" bson:"allowedGeographicArea" mapstructure:"AllowedGeographicArea"`
 
	

	PrivacyCheckRelatedAction *PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty" yaml:"privacyCheckRelatedAction" bson:"privacyCheckRelatedAction" mapstructure:"PrivacyCheckRelatedAction"`
 
	

	ValidTimePeriod *ValidTimePeriod `json:"validTimePeriod,omitempty" yaml:"validTimePeriod" bson:"validTimePeriod" mapstructure:"ValidTimePeriod"`
}
