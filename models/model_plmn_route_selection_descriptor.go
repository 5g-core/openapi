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


// PlmnRouteSelectionDescriptor - Contains the route selection descriptors (combinations of SNSSAI, DNNs, PDU session types, SSC modes and ATSSS information) allowed by subscription to the UE for a serving PLMN
type PlmnRouteSelectionDescriptor struct {


	
	// simple type
 
	

	ServingPlmn PlmnId `json:"servingPlmn" yaml:"servingPlmn" bson:"servingPlmn" mapstructure:"ServingPlmn"`
 
	

	SnssaiRouteSelDescs *[]SnssaiRouteSelectionDescriptor `json:"snssaiRouteSelDescs,omitempty" yaml:"snssaiRouteSelDescs" bson:"snssaiRouteSelDescs" mapstructure:"SnssaiRouteSelDescs"`
}
