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


type AmfNon3GppAccessRegistration struct {


	
	// simple type
 
	

	AmfInstanceId string `json:"amfInstanceId" yaml:"amfInstanceId" bson:"amfInstanceId" mapstructure:"AmfInstanceId"`
 
	

	SupportedFeatures *string `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures" mapstructure:"SupportedFeatures"`
 
	

	PurgeFlag *bool `json:"purgeFlag,omitempty" yaml:"purgeFlag" bson:"purgeFlag" mapstructure:"PurgeFlag"`
 
	

	Pei *string `json:"pei,omitempty" yaml:"pei" bson:"pei" mapstructure:"Pei"`
 
	

	ImsVoPs ImsVoPs `json:"imsVoPs" yaml:"imsVoPs" bson:"imsVoPs" mapstructure:"ImsVoPs"`
 
	

	DeregCallbackUri string `json:"deregCallbackUri" yaml:"deregCallbackUri" bson:"deregCallbackUri" mapstructure:"DeregCallbackUri"`
 
	

	AmfServiceNameDereg *ServiceName `json:"amfServiceNameDereg,omitempty" yaml:"amfServiceNameDereg" bson:"amfServiceNameDereg" mapstructure:"AmfServiceNameDereg"`
 
	

	PcscfRestorationCallbackUri *string `json:"pcscfRestorationCallbackUri,omitempty" yaml:"pcscfRestorationCallbackUri" bson:"pcscfRestorationCallbackUri" mapstructure:"PcscfRestorationCallbackUri"`
 
	

	AmfServiceNamePcscfRest *ServiceName `json:"amfServiceNamePcscfRest,omitempty" yaml:"amfServiceNamePcscfRest" bson:"amfServiceNamePcscfRest" mapstructure:"AmfServiceNamePcscfRest"`
 
	

	Guami Guami `json:"guami" yaml:"guami" bson:"guami" mapstructure:"Guami"`
 
	

	BackupAmfInfo *[]BackupAmfInfo `json:"backupAmfInfo,omitempty" yaml:"backupAmfInfo" bson:"backupAmfInfo" mapstructure:"BackupAmfInfo"`
 
	

	RatType RatType `json:"ratType" yaml:"ratType" bson:"ratType" mapstructure:"RatType"`
 
	

	UrrpIndicator *bool `json:"urrpIndicator,omitempty" yaml:"urrpIndicator" bson:"urrpIndicator" mapstructure:"UrrpIndicator"`
 
	

	AmfEeSubscriptionId *string `json:"amfEeSubscriptionId,omitempty" yaml:"amfEeSubscriptionId" bson:"amfEeSubscriptionId" mapstructure:"AmfEeSubscriptionId"`
 
	

	RegistrationTime *time.Time `json:"registrationTime,omitempty" yaml:"registrationTime" bson:"registrationTime" mapstructure:"RegistrationTime"`
 
	

	VgmlcAddress *VgmlcAddress `json:"vgmlcAddress,omitempty" yaml:"vgmlcAddress" bson:"vgmlcAddress" mapstructure:"VgmlcAddress"`
 
	

	ContextInfo *ContextInfo `json:"contextInfo,omitempty" yaml:"contextInfo" bson:"contextInfo" mapstructure:"ContextInfo"`
 
	

	NoEeSubscriptionInd *bool `json:"noEeSubscriptionInd,omitempty" yaml:"noEeSubscriptionInd" bson:"noEeSubscriptionInd" mapstructure:"NoEeSubscriptionInd"`
 
	

	Supi *string `json:"supi,omitempty" yaml:"supi" bson:"supi" mapstructure:"Supi"`
}
