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


type AmfSubscriptionInfo struct {


	
	// simple type
 
	

	AmfInstanceId string `json:"amfInstanceId" yaml:"amfInstanceId" bson:"amfInstanceId" mapstructure:"AmfInstanceId"`
 
	

	SubscriptionId string `json:"subscriptionId" yaml:"subscriptionId" bson:"subscriptionId" mapstructure:"SubscriptionId"`
 
	

	SubsChangeNotifyCorrelationId *string `json:"subsChangeNotifyCorrelationId,omitempty" yaml:"subsChangeNotifyCorrelationId" bson:"subsChangeNotifyCorrelationId" mapstructure:"SubsChangeNotifyCorrelationId"`
}
