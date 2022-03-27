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


// NetworkAreaInfo - Describes a network area information in which the NF service consumer requests the number of UEs.
type NetworkAreaInfo struct {


	
	// simple type

	// Contains a list of E-UTRA cell identities. 
	

	Ecgis *[]Ecgi `json:"ecgis,omitempty" yaml:"ecgis" bson:"ecgis" mapstructure:"Ecgis"`

	// Contains a list of NR cell identities. 
	

	Ncgis *[]Ncgi `json:"ncgis,omitempty" yaml:"ncgis" bson:"ncgis" mapstructure:"Ncgis"`

	// Contains a list of NG RAN nodes. 
	

	GRanNodeIds *[]GlobalRanNodeId `json:"gRanNodeIds,omitempty" yaml:"gRanNodeIds" bson:"gRanNodeIds" mapstructure:"GRanNodeIds"`

	// Contains a list of tracking area identities. 
	

	Tais *[]Tai `json:"tais,omitempty" yaml:"tais" bson:"tais" mapstructure:"Tais"`
}
