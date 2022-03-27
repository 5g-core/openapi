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


type PfdDataForApp struct {


	
	// simple type
 
	

	ApplicationId string `json:"applicationId" yaml:"applicationId" bson:"applicationId" mapstructure:"ApplicationId"`
 
	

	Pfds []PfdContent `json:"pfds" yaml:"pfds" bson:"pfds" mapstructure:"Pfds"`
 
	

	CachingTime *time.Time `json:"cachingTime,omitempty" yaml:"cachingTime" bson:"cachingTime" mapstructure:"CachingTime"`
}
