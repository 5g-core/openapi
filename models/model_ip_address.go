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


type IpAddress struct {


	
	// simple type
 
	

	Ipv4Addr *string `json:"ipv4Addr,omitempty" yaml:"ipv4Addr" bson:"ipv4Addr" mapstructure:"Ipv4Addr"`
 
	

	Ipv6Addr *Ipv6Addr `json:"ipv6Addr,omitempty" yaml:"ipv6Addr" bson:"ipv6Addr" mapstructure:"Ipv6Addr"`
 
	

	Ipv6Prefix *Ipv6Prefix `json:"ipv6Prefix,omitempty" yaml:"ipv6Prefix" bson:"ipv6Prefix" mapstructure:"Ipv6Prefix"`
}
