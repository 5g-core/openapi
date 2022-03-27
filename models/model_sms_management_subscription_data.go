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


type SmsManagementSubscriptionData struct {


	
	// simple type
 
	

	SupportedFeatures *string `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures" mapstructure:"SupportedFeatures"`
 
	

	MtSmsSubscribed *bool `json:"mtSmsSubscribed,omitempty" yaml:"mtSmsSubscribed" bson:"mtSmsSubscribed" mapstructure:"MtSmsSubscribed"`
 
	

	MtSmsBarringAll *bool `json:"mtSmsBarringAll,omitempty" yaml:"mtSmsBarringAll" bson:"mtSmsBarringAll" mapstructure:"MtSmsBarringAll"`
 
	

	MtSmsBarringRoaming *bool `json:"mtSmsBarringRoaming,omitempty" yaml:"mtSmsBarringRoaming" bson:"mtSmsBarringRoaming" mapstructure:"MtSmsBarringRoaming"`
 
	

	MoSmsSubscribed *bool `json:"moSmsSubscribed,omitempty" yaml:"moSmsSubscribed" bson:"moSmsSubscribed" mapstructure:"MoSmsSubscribed"`
 
	

	MoSmsBarringAll *bool `json:"moSmsBarringAll,omitempty" yaml:"moSmsBarringAll" bson:"moSmsBarringAll" mapstructure:"MoSmsBarringAll"`
 
	

	MoSmsBarringRoaming *bool `json:"moSmsBarringRoaming,omitempty" yaml:"moSmsBarringRoaming" bson:"moSmsBarringRoaming" mapstructure:"MoSmsBarringRoaming"`
 
	

	SharedSmsMngDataIds *[]string `json:"sharedSmsMngDataIds,omitempty" yaml:"sharedSmsMngDataIds" bson:"sharedSmsMngDataIds" mapstructure:"SharedSmsMngDataIds"`
 
	

	TraceData *TraceData `json:"traceData,omitempty" yaml:"traceData" bson:"traceData" mapstructure:"TraceData"`
}
