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


type PgwInfo struct {


	
	// simple type
 
	

	Dnn string `json:"dnn" yaml:"dnn" bson:"dnn" mapstructure:"Dnn"`
 
	

	PgwFqdn string `json:"pgwFqdn" yaml:"pgwFqdn" bson:"pgwFqdn" mapstructure:"PgwFqdn"`
 
	

	PlmnId *PlmnId `json:"plmnId,omitempty" yaml:"plmnId" bson:"plmnId" mapstructure:"PlmnId"`
 
	

	EpdgInd *bool `json:"epdgInd,omitempty" yaml:"epdgInd" bson:"epdgInd" mapstructure:"EpdgInd"`
}
