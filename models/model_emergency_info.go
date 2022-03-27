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


type EmergencyInfo struct {


	
	// simple type
 
	

	PgwFqdn *string `json:"pgwFqdn,omitempty" yaml:"pgwFqdn" bson:"pgwFqdn" mapstructure:"PgwFqdn"`
 
	

	PgwIpAddress *IpAddress `json:"pgwIpAddress,omitempty" yaml:"pgwIpAddress" bson:"pgwIpAddress" mapstructure:"PgwIpAddress"`
 
	

	SmfInstanceId *string `json:"smfInstanceId,omitempty" yaml:"smfInstanceId" bson:"smfInstanceId" mapstructure:"SmfInstanceId"`
 
	

	EpdgInd *bool `json:"epdgInd,omitempty" yaml:"epdgInd" bson:"epdgInd" mapstructure:"EpdgInd"`
}