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


type AuthorizationData struct {


	
	// simple type
 
	

	AuthorizationData []UserIdentifier `json:"authorizationData" yaml:"authorizationData" bson:"authorizationData" mapstructure:"AuthorizationData"`
 
	

	ValidityTime *time.Time `json:"validityTime,omitempty" yaml:"validityTime" bson:"validityTime" mapstructure:"ValidityTime"`
}