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


type SubscriptionDataSubscriptions struct {


	
	// simple type
 
	

	UeId *string `json:"ueId,omitempty" yaml:"ueId" bson:"ueId" mapstructure:"UeId"`
 
	

	CallbackReference string `json:"callbackReference" yaml:"callbackReference" bson:"callbackReference" mapstructure:"CallbackReference"`
 
	

	OriginalCallbackReference *string `json:"originalCallbackReference,omitempty" yaml:"originalCallbackReference" bson:"originalCallbackReference" mapstructure:"OriginalCallbackReference"`
 
	

	MonitoredResourceUris []string `json:"monitoredResourceUris" yaml:"monitoredResourceUris" bson:"monitoredResourceUris" mapstructure:"MonitoredResourceUris"`
 
	

	Expiry *time.Time `json:"expiry,omitempty" yaml:"expiry" bson:"expiry" mapstructure:"Expiry"`
 
	

	SdmSubscription *SdmSubscription `json:"sdmSubscription,omitempty" yaml:"sdmSubscription" bson:"sdmSubscription" mapstructure:"SdmSubscription"`
 
	

	SubscriptionId *string `json:"subscriptionId,omitempty" yaml:"subscriptionId" bson:"subscriptionId" mapstructure:"SubscriptionId"`
 
	

	SupportedFeatures *string `json:"supported-features,omitempty" yaml:"supported-features" bson:"supported-features" mapstructure:"SupportedFeatures"`
}
