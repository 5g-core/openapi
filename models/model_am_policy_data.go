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


// AmPolicyData - Contains the AM policy data for a given subscriber.
type AmPolicyData struct {


	
	// simple type
 
	

	PraInfos *map[string]PresenceInfo `json:"praInfos,omitempty" yaml:"praInfos" bson:"praInfos" mapstructure:"PraInfos"`
 
	

	SubscCats *[]string `json:"subscCats,omitempty" yaml:"subscCats" bson:"subscCats" mapstructure:"SubscCats"`
}
