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


type ContextDataSets struct {


	
	// simple type
 
	

	Amf3Gpp *Amf3GppAccessRegistration `json:"amf3Gpp,omitempty" yaml:"amf3Gpp" bson:"amf3Gpp" mapstructure:"Amf3Gpp"`
 
	

	AmfNon3Gpp *AmfNon3GppAccessRegistration `json:"amfNon3Gpp,omitempty" yaml:"amfNon3Gpp" bson:"amfNon3Gpp" mapstructure:"AmfNon3Gpp"`
 
	

	SdmSubscriptions *[]SdmSubscription `json:"sdmSubscriptions,omitempty" yaml:"sdmSubscriptions" bson:"sdmSubscriptions" mapstructure:"SdmSubscriptions"`
 
	

	EeSubscriptions *[]EeSubscription `json:"eeSubscriptions,omitempty" yaml:"eeSubscriptions" bson:"eeSubscriptions" mapstructure:"EeSubscriptions"`
 
	

	Smsf3GppAccess *SmsfRegistration `json:"smsf3GppAccess,omitempty" yaml:"smsf3GppAccess" bson:"smsf3GppAccess" mapstructure:"Smsf3GppAccess"`
 
	

	SmsfNon3GppAccess *SmsfRegistration `json:"smsfNon3GppAccess,omitempty" yaml:"smsfNon3GppAccess" bson:"smsfNon3GppAccess" mapstructure:"SmsfNon3GppAccess"`
 
	

	SubscriptionDataSubscriptions *[]SubscriptionDataSubscriptions `json:"subscriptionDataSubscriptions,omitempty" yaml:"subscriptionDataSubscriptions" bson:"subscriptionDataSubscriptions" mapstructure:"SubscriptionDataSubscriptions"`
 
	

	SmfRegistrations *SmfRegList `json:"smfRegistrations,omitempty" yaml:"smfRegistrations" bson:"smfRegistrations" mapstructure:"SmfRegistrations"`
 
	

	IpSmGw *IpSmGwRegistration `json:"ipSmGw,omitempty" yaml:"ipSmGw" bson:"ipSmGw" mapstructure:"IpSmGw"`
}
