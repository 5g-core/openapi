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


type AcsInfo1 struct {


	
	// simple type
 
	

	AcsUrl *string `json:"acsUrl,omitempty" yaml:"acsUrl" bson:"acsUrl" mapstructure:"AcsUrl"`
 
	

	AcsIpv4Addr *string `json:"acsIpv4Addr,omitempty" yaml:"acsIpv4Addr" bson:"acsIpv4Addr" mapstructure:"AcsIpv4Addr"`
 
	

	AcsIpv6Addr *Ipv6Addr `json:"acsIpv6Addr,omitempty" yaml:"acsIpv6Addr" bson:"acsIpv6Addr" mapstructure:"AcsIpv6Addr"`
}