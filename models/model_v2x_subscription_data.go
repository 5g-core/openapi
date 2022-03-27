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


type V2xSubscriptionData struct {


	
	// simple type
 
	

	NrV2xServicesAuth *NrV2xAuth `json:"nrV2xServicesAuth,omitempty" yaml:"nrV2xServicesAuth" bson:"nrV2xServicesAuth" mapstructure:"NrV2xServicesAuth"`
 
	

	LteV2xServicesAuth *LteV2xAuth `json:"lteV2xServicesAuth,omitempty" yaml:"lteV2xServicesAuth" bson:"lteV2xServicesAuth" mapstructure:"LteV2xServicesAuth"`
 
	

	NrUePc5Ambr *string `json:"nrUePc5Ambr,omitempty" yaml:"nrUePc5Ambr" bson:"nrUePc5Ambr" mapstructure:"NrUePc5Ambr"`
 
	

	LtePc5Ambr *string `json:"ltePc5Ambr,omitempty" yaml:"ltePc5Ambr" bson:"ltePc5Ambr" mapstructure:"LtePc5Ambr"`
}