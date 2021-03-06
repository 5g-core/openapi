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


// EthFlowDescription - Identifies an Ethernet flow
type EthFlowDescription struct {


	
	// simple type
 
	

	DestMacAddr *string `json:"destMacAddr,omitempty" yaml:"destMacAddr" bson:"destMacAddr" mapstructure:"DestMacAddr"`
 
	

	EthType string `json:"ethType" yaml:"ethType" bson:"ethType" mapstructure:"EthType"`

	// Defines a packet filter of an IP flow. 
	

	FDesc *string `json:"fDesc,omitempty" yaml:"fDesc" bson:"fDesc" mapstructure:"FDesc"`
 
	

	FDir *FlowDirection `json:"fDir,omitempty" yaml:"fDir" bson:"fDir" mapstructure:"FDir"`
 
	

	SourceMacAddr *string `json:"sourceMacAddr,omitempty" yaml:"sourceMacAddr" bson:"sourceMacAddr" mapstructure:"SourceMacAddr"`
 
	

	VlanTags *[]string `json:"vlanTags,omitempty" yaml:"vlanTags" bson:"vlanTags" mapstructure:"VlanTags"`
 
	

	SrcMacAddrEnd *string `json:"srcMacAddrEnd,omitempty" yaml:"srcMacAddrEnd" bson:"srcMacAddrEnd" mapstructure:"SrcMacAddrEnd"`
 
	

	DestMacAddrEnd *string `json:"destMacAddrEnd,omitempty" yaml:"destMacAddrEnd" bson:"destMacAddrEnd" mapstructure:"DestMacAddrEnd"`
}
