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


type DnnConfiguration struct {


	
	// simple type
 
	

	PduSessionTypes PduSessionTypes `json:"pduSessionTypes" yaml:"pduSessionTypes" bson:"pduSessionTypes" mapstructure:"PduSessionTypes"`
 
	

	SscModes SscModes `json:"sscModes" yaml:"sscModes" bson:"sscModes" mapstructure:"SscModes"`
 
	

	IwkEpsInd *bool `json:"iwkEpsInd,omitempty" yaml:"iwkEpsInd" bson:"iwkEpsInd" mapstructure:"IwkEpsInd"`
 
	

	Var5gQosProfile *SubscribedDefaultQos `json:"5gQosProfile,omitempty" yaml:"5gQosProfile" bson:"5gQosProfile" mapstructure:"Var5gQosProfile"`
 
	

	SessionAmbr *Ambr1 `json:"sessionAmbr,omitempty" yaml:"sessionAmbr" bson:"sessionAmbr" mapstructure:"SessionAmbr"`
 
	

	Var3gppChargingCharacteristics *string `json:"3gppChargingCharacteristics,omitempty" yaml:"3gppChargingCharacteristics" bson:"3gppChargingCharacteristics" mapstructure:"Var3gppChargingCharacteristics"`
 
	

	StaticIpAddress *[]IpAddress `json:"staticIpAddress,omitempty" yaml:"staticIpAddress" bson:"staticIpAddress" mapstructure:"StaticIpAddress"`
 
	

	UpSecurity *UpSecurity `json:"upSecurity,omitempty" yaml:"upSecurity" bson:"upSecurity" mapstructure:"UpSecurity"`
 
	

	PduSessionContinuityInd *PduSessionContinuityInd `json:"pduSessionContinuityInd,omitempty" yaml:"pduSessionContinuityInd" bson:"pduSessionContinuityInd" mapstructure:"PduSessionContinuityInd"`

	// Identity of the NEF 
	

	NiddNefId *string `json:"niddNefId,omitempty" yaml:"niddNefId" bson:"niddNefId" mapstructure:"NiddNefId"`
 
	

	NiddInfo *NiddInformation `json:"niddInfo,omitempty" yaml:"niddInfo" bson:"niddInfo" mapstructure:"NiddInfo"`
 
	

	RedundantSessionAllowed *bool `json:"redundantSessionAllowed,omitempty" yaml:"redundantSessionAllowed" bson:"redundantSessionAllowed" mapstructure:"RedundantSessionAllowed"`
 
	

	AcsInfo *AcsInfo `json:"acsInfo,omitempty" yaml:"acsInfo" bson:"acsInfo" mapstructure:"AcsInfo"`
 
	

	Ipv4FrameRouteList *[]FrameRouteInfo `json:"ipv4FrameRouteList,omitempty" yaml:"ipv4FrameRouteList" bson:"ipv4FrameRouteList" mapstructure:"Ipv4FrameRouteList"`
 
	

	Ipv6FrameRouteList *[]FrameRouteInfo `json:"ipv6FrameRouteList,omitempty" yaml:"ipv6FrameRouteList" bson:"ipv6FrameRouteList" mapstructure:"Ipv6FrameRouteList"`
 
	

	AtsssAllowed *bool `json:"atsssAllowed,omitempty" yaml:"atsssAllowed" bson:"atsssAllowed" mapstructure:"AtsssAllowed"`
 
	

	SecondaryAuth *bool `json:"secondaryAuth,omitempty" yaml:"secondaryAuth" bson:"secondaryAuth" mapstructure:"SecondaryAuth"`
 
	

	DnAaaIpAddressAllocation *bool `json:"dnAaaIpAddressAllocation,omitempty" yaml:"dnAaaIpAddressAllocation" bson:"dnAaaIpAddressAllocation" mapstructure:"DnAaaIpAddressAllocation"`
 
	

	DnAaaAddress *IpAddress `json:"dnAaaAddress,omitempty" yaml:"dnAaaAddress" bson:"dnAaaAddress" mapstructure:"DnAaaAddress"`
 
	

	IptvAccCtrlInfo *string `json:"iptvAccCtrlInfo,omitempty" yaml:"iptvAccCtrlInfo" bson:"iptvAccCtrlInfo" mapstructure:"IptvAccCtrlInfo"`
}
